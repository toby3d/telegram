package telegram

import (
	"fmt"
	"net/url"
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
	return m != nil &&
		b != nil &&
		m.From != nil &&
		b.User != nil &&
		m.From.ID == b.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the
// current bot.
func (b *Bot) IsForwardFromMe(m *Message) bool {
	return m.IsForward() &&
		b != nil &&
		b.User != nil &&
		m.ForwardFrom.ID == b.ID
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
	if m == nil ||
		b == nil ||
		b.User == nil {
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
		if entity.IsMention() {
			if b.ID == entity.User.ID {
				return true
			}
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
func (b *Bot) NewFileURL(filePath string) *url.URL {
	if b == nil ||
		strings.EqualFold(b.AccessToken, "") ||
		strings.EqualFold(filePath, "") {
		return nil
	}

	result := defaultURI
	result.Path = path.Join("file", fmt.Sprint("bot", b.AccessToken), filePath)

	return result
}

// NewRedirectURL creates new url.URL for redirecting from one chat to another.
func (b *Bot) NewRedirectURL(param string, group bool) *url.URL {
	if b == nil ||
		b.User == nil ||
		b.User.Username == "" {
		return nil
	}

	link := &url.URL{
		Scheme: "https",
		Host:   "t.me",
		Path:   b.User.Username,
	}

	q := link.Query()
	key := "start"
	if group {
		key += "group"
	}
	q.Add(key, param)

	link.RawQuery = q.Encode()

	return link
}
