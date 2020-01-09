package telegram

import (
	"bytes"
	"log"
	"net"
	"path"
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
	Updates     chan *Update

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

// IsMessageFromMe checks that the input message is a message from the current bot.
func (b *Bot) IsMessageFromMe(m *Message) bool {
	return b != nil && b.User != nil && m != nil && m.From != nil && m.From.ID == b.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the current bot.
func (b *Bot) IsForwardFromMe(m *Message) bool {
	return b != nil && b.User != nil && m.IsForward() && m.ForwardFrom.ID == b.ID
}

// IsReplyToMe checks that the input message is a reply to the current bot.
func (b *Bot) IsReplyToMe(m *Message) bool {
	return m.Chat.IsPrivate() || (m.IsReply() && b.IsMessageFromMe(m.ReplyToMessage))
}

// IsCommandToMe checks that the input message is a command for the current bot.
func (b *Bot) IsCommandToMe(m *Message) bool {
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
func (b *Bot) IsMessageMentionsMe(m *Message) bool {
	if b == nil || b.User == nil || m == nil {
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
func (b *Bot) IsForwardMentionsMe(m *Message) bool { return m.IsForward() && b.IsMessageMentionsMe(m) }

// IsReplyMentionsMe checks that the input message mentions the current b.
func (b *Bot) IsReplyMentionsMe(m *Message) bool {
	return m.IsReply() && b.IsMessageMentionsMe(m.ReplyToMessage)
}

// IsMessageToMe checks that the input message is addressed to the current b.
func (b *Bot) IsMessageToMe(m *Message) bool {
	if m == nil || m.Chat == nil {
		return false
	}

	if m.Chat.IsPrivate() || b.IsCommandToMe(m) || b.IsReplyToMe(m) || b.IsMessageMentionsMe(m) {
		return true
	}

	return false
}

// NewFileURL creates a url.URL to file with path getted from GetFile method.
func (b *Bot) NewFileURL(filePath string) *http.URI {
	if b == nil || b.AccessToken == "" ||
		filePath == "" {
		return nil
	}

	result := http.AcquireURI()
	result.SetScheme("https")
	result.SetHost("api.telegram.org")
	result.SetPath(path.Join("file", "bot"+b.AccessToken, filePath))

	return result
}

// NewRedirectURL creates new url.URL for redirecting from one chat to another.
func (b *Bot) NewRedirectURL(param string, group bool) *http.URI {
	if b == nil || b.User == nil || b.User.Username == "" {
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
func (b *Bot) NewLongPollingChannel(params *GetUpdates) chan *Update {
	if params == nil {
		params = &GetUpdates{
			Offset:  0,
			Limit:   100,
			Timeout: 60,
		}
	}

	b.Updates = make(chan *Update, params.Limit)

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

// NewWebhookChannel creates channel for receive incoming updates via an outgoing webhook.
//
// If cert argument is provided by two strings (["path/to/cert.file", "path/to/cert.key"]), then TLS server will be created by this filepaths.
func (b *Bot) NewWebhookChannel(u *http.URI, p SetWebhook, ln net.Listener, crt ...string) (chan *Update, func() error) {
	b.Updates = make(chan *Update, 100)
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

	var err error

	go func() {
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

	if _, err = b.SetWebhook(p); err != nil {
		log.Fatalln(err.Error())
	}

	return b.Updates, srv.Shutdown
}
