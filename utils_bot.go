package telegram

import (
	"path"
	"strings"

	http "github.com/valyala/fasthttp"
)

// Bot represents a bot user with access token getted from @BotFather and
// fasthttp.Client for requests.
type Bot struct {
	*User
	AccessToken string
	Client      *http.Client
}

// SetClient allow set custom fasthttp.Client (for proxy traffic, for example).
func (b *Bot) SetClient(newClient *http.Client) {
	if b == nil {
		b = new(Bot)
	}

	b.Client = newClient
}

// New creates a new default Bot structure based on the input access token.
func New(accessToken string) (*Bot, error) {
	var err error
	b := new(Bot)
	b.SetClient(defaultClient)
	b.AccessToken = accessToken

	b.User, err = b.GetMe()
	return b, err
}

// IsMessageFromMe checks that the input message is a message from the current
// bot.
func (b *Bot) IsMessageFromMe(m *Message) bool {
	return b != nil && b.User != nil &&
		m != nil && m.From != nil && m.From.ID == b.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the
// current bot.
func (b *Bot) IsForwardFromMe(m *Message) bool {
	return b != nil && b.User != nil &&
		m.IsForward() && m.ForwardFrom.ID == b.ID
}

// IsReplyToMe checks that the input message is a reply to the current bot.
func (b *Bot) IsReplyToMe(m *Message) bool {
	return m.Chat.IsPrivate() ||
		(m.IsReply() && b.IsMessageFromMe(m.ReplyToMessage))
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
	if b == nil || b.User == nil ||
		m == nil {
		return false
	}

	if b.IsCommandToMe(m) {
		return true
	}

	var entities []MessageEntity
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

// IsForwardMentionsMe checks that the input forwarded message mentions the
// current bot.
func (b *Bot) IsForwardMentionsMe(m *Message) bool {
	return m.IsForward() && b.IsMessageMentionsMe(m)
}

// IsReplyMentionsMe checks that the input message mentions the current bot.
func (b *Bot) IsReplyMentionsMe(m *Message) bool {
	return m.IsReply() && b.IsMessageMentionsMe(m.ReplyToMessage)
}

// IsMessageToMe checks that the input message is addressed to the current bot.
func (b *Bot) IsMessageToMe(m *Message) bool {
	if m == nil || m.Chat == nil {
		return false
	}

	if m.Chat.IsPrivate() ||
		b.IsCommandToMe(m) ||
		b.IsReplyToMe(m) ||
		b.IsMessageMentionsMe(m) {
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

func (b *Bot) DecryptPassportFile(pf *PassportFile, fc *FileCredentials) (data []byte, err error) {
	secret, err := decodeField(fc.Secret)
	if err != nil {
		return
	}

	hash, err := decodeField(fc.FileHash)
	if err != nil {
		return
	}

	key, iv := decryptSecretHash(secret, hash)
	file, err := b.GetFile(pf.FileID)
	if err != nil {
		return
	}

	_, data, err = b.Client.Get(nil, b.NewFileURL(file.FilePath).String())
	if err != nil {
		return
	}

	data, err = decryptData(key, iv, data)
	if err != nil {
		return
	}

	if !match(hash, data) {
		err = ErrNotEqual
		return
	}

	offset := int(data[0])
	data = data[offset:]

	return
}
