package telegram

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net"
	"path"
	"path/filepath"
	"strings"
	"time"

	json "github.com/json-iterator/go"
	"github.com/kirillDanshin/dlog"
	http "github.com/valyala/fasthttp"
)

// Bot represents a bot user with access token getted from @BotFather.
type Bot struct {
	*User
	AccessToken string
	Updates     UpdatesChannel

	client   *http.Client
	marshler json.API
}

// New creates a new default Bot structure based on the input access token.
func New(accessToken string) (b *Bot, err error) {
	b = new(Bot)
	b.marshler = json.ConfigFastest
	b.SetClient(&http.Client{})
	b.AccessToken = accessToken
	b.User, err = b.GetMe()

	return b, err
}

// SetClient allow set custom fasthttp.Client (for proxy traffic, for example).
func (b *Bot) SetClient(newClient *http.Client) {
	if b == nil {
		b = new(Bot)
	}

	b.client = newClient
}

func (b Bot) Do(method string, payload interface{}) ([]byte, error) {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegram.org")
	u.SetPath(path.Join("bot"+b.AccessToken, method))

	var buf bytes.Buffer
	if err := b.marshler.NewEncoder(&buf).Encode(payload); err != nil {
		return nil, err
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetUserAgent("toby3d/telegram")
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetHostBytes(u.Host())
	req.SetRequestURI(u.String())
	req.SetBody(buf.Bytes())

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := b.client.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (b Bot) Upload(method string, payload map[string]string, files ...*InputFile) ([]byte, error) {
	if len(files) == 0 {
		return b.Do(method, payload)
	}

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	for i := range files {
		_, fileName := filepath.Split(files[i].Attachment.Name())

		part, err := w.CreateFormFile(fileName, fileName)
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, files[i].Attachment); err != nil {
			return nil, err
		}
	}

	for key, val := range payload {
		if err := w.WriteField(key, val); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegram.org")
	u.SetPath(path.Join("bot"+b.AccessToken, method))

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURIBytes(u.RequestURI())
	req.Header.SetContentType(w.FormDataContentType())
	req.Header.SetMultipartFormBoundary(w.Boundary())

	if _, err := body.WriteTo(req.BodyWriter()); err != nil {
		return nil, err
	}

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := b.client.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

// IsMessageFromMe checks that the input message is a message from the current bot.
func (b Bot) IsMessageFromMe(m Message) bool {
	return b.User != nil && m.From != nil && m.From.ID == b.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the current bot.
func (b Bot) IsForwardFromMe(m Message) bool {
	return b.User != nil && m.IsForward() && m.ForwardFrom.ID == b.ID
}

// IsReplyToMe checks that the input message is a reply to the current bot.
func (b Bot) IsReplyToMe(m Message) bool {
	return m.Chat.IsPrivate() || (m.IsReply() && b.IsMessageFromMe(*m.ReplyToMessage))
}

// IsCommandToMe checks that the input message is a command for the current bot.
func (b Bot) IsCommandToMe(m Message) bool {
	if !m.IsCommand() {
		return false
	}

	if m.Chat.IsPrivate() {
		return true
	}

	parts := strings.Split(m.RawCommand(), "@")
	if len(parts) <= 1 {
		return false
	}

	return strings.EqualFold(parts[1], b.User.Username)
}

// IsMessageMentionsMe checks that the input message mentions the current bot.
func (b Bot) IsMessageMentionsMe(m Message) bool {
	if b.User == nil {
		return false
	}

	if b.IsCommandToMe(m) {
		return true
	}

	var entities []*MessageEntity

	switch {
	case m.HasMentions():
		entities = m.Entities
	case m.HasCaptionMentions():
		entities = m.CaptionEntities
	}

	for _, entity := range entities {
		if entity.IsMention() && entity.User.ID == b.ID {
			return true
		}
	}

	return false
}

// IsForwardMentionsMe checks that the input forwarded message mentions the current bot.
func (b Bot) IsForwardMentionsMe(m Message) bool { return m.IsForward() && b.IsMessageMentionsMe(m) }

// IsReplyMentionsMe checks that the input message mentions the current b.
func (b Bot) IsReplyMentionsMe(m Message) bool {
	return m.IsReply() && b.IsMessageMentionsMe(*m.ReplyToMessage)
}

// IsMessageToMe checks that the input message is addressed to the current b.
func (b Bot) IsMessageToMe(m Message) bool {
	return m.Chat != nil && (m.Chat.IsPrivate() || b.IsCommandToMe(m) || b.IsReplyToMe(m) ||
		b.IsMessageMentionsMe(m))
}

// NewFileURL creates a fasthttp.URI to file with path getted from GetFile method.
func (b Bot) NewFileURL(filePath string) *http.URI {
	if b.AccessToken == "" || filePath == "" {
		return nil
	}

	result := http.AcquireURI()
	result.SetScheme("https")
	result.SetHost("api.telegram.org")
	result.SetPath(path.Join("file", "bot"+b.AccessToken, filePath))

	return result
}

// NewRedirectURL creates new fasthttp.URI for redirecting from one chat to another.
func (b Bot) NewRedirectURL(param string, group bool) *http.URI {
	if b.User == nil || b.User.Username == "" {
		return nil
	}

	link := http.AcquireURI()
	link.SetScheme("https")
	link.SetHost("t.me")
	link.SetPath(b.User.Username)

	q := link.QueryArgs()
	key := "start"

	if group {
		key += "group"
	}

	q.Set(key, param)
	link.SetQueryStringBytes(q.QueryString())

	return link
}

// NewLongPollingChannel creates channel for receive incoming updates using long polling.
func (b *Bot) NewLongPollingChannel(params *GetUpdates) UpdatesChannel {
	if params == nil {
		params = &GetUpdates{
			Offset:  0,
			Limit:   100,
			Timeout: 60,
		}
	}

	b.Updates = make(UpdatesChannel, params.Limit)

	go func() {
		for {
			updates, err := b.GetUpdates(params)
			if err != nil {
				dlog.Ln(err.Error())
				dlog.Ln("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateID < params.Offset {
					continue
				}

				params.Offset = update.UpdateID + 1
				b.Updates <- update
			}
		}
	}()

	return b.Updates
}

// NewWebhookChannel creates channel for receive incoming updates via an outgoing webhook. Returns updates channel and
// shutdown func.
//
// If cert argument is provided by two strings (["path/to/cert.file", "path/to/cert.key"]), then TLS server will be
// created by this filepaths.
func (b *Bot) NewWebhookChannel(u *http.URI, p SetWebhook, ln net.Listener, crt ...string) (UpdatesChannel,
	func() error) {
	b.Updates = make(UpdatesChannel, 100) // NOTE(toby3d): channel size by default GetUpdates.Limit parameter
	handleFunc := func(ctx *http.RequestCtx) {
		dlog.Ln("Request path:", string(ctx.Path()))

		if !bytes.HasPrefix(ctx.Path(), u.Path()) {
			dlog.Ln("Unsupported request path:", string(ctx.Path()))

			return
		}

		dlog.Ln("Catched supported request path:", string(ctx.Path()))

		upd := new(Update)
		if err := b.marshler.Unmarshal(ctx.Request.Body(), upd); err != nil {
			return
		}

		b.Updates <- upd
	}

	srv := http.Server{
		Name:              b.Username,
		Concurrency:       p.MaxConnections,
		Handler:           handleFunc,
		ReduceMemoryUsage: true,
	}

	go func() {
		var err error

		switch {
		case len(crt) == 2:
			dlog.Ln("Creating TLS router...")
			err = srv.ServeTLS(ln, crt[0], crt[1])
		default:
			dlog.Ln("Creating simple router...")
			err = srv.Serve(ln)
		}

		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	if _, err := b.SetWebhook(p); err != nil {
		log.Fatalln(err.Error())
	}

	return b.Updates, srv.Shutdown
}
