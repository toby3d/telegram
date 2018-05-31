package telegram

import (
	"fmt"
	"net/url"
	"strings"

	http "github.com/valyala/fasthttp"
)

// Bot represents a bot user with access token getted from @BotFather and
// fasthttp.Client for requests.
type Bot struct {
	AccessToken string
	Client      *http.Client
	*User
}

// SetClient allow set custom fasthttp.Client (for proxy traffic, for example).
func (bot *Bot) SetClient(newClient *http.Client) {
	bot.Client = newClient
}

// New creates a new default Bot structure based on the input access token.
func New(accessToken string) (*Bot, error) {
	var err error
	bot := new(Bot)
	bot.SetClient(defaultClient)
	bot.AccessToken = accessToken

	bot.User, err = bot.GetMe()
	return bot, err
}

// IsMessageFromMe checks that the input message is a message from the current
// bot.
func (bot *Bot) IsMessageFromMe(msg *Message) bool {
	return msg != nil && bot != nil && msg.From != nil && bot.User != nil && msg.From.ID == bot.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the
// current bot.
func (bot *Bot) IsForwardFromMe(msg *Message) bool {
	return msg.IsForward() && bot != nil && bot.User != nil && msg.ForwardFrom.ID == bot.ID
}

// IsReplyToMe checks that the input message is a reply to the current bot.
func (bot *Bot) IsReplyToMe(msg *Message) bool {
	return msg.Chat.IsPrivate() || (msg.IsReply() && bot.IsMessageFromMe(msg.ReplyToMessage))
}

// IsCommandToMe checks that the input message is a command for the current bot.
func (bot *Bot) IsCommandToMe(msg *Message) bool {
	if !msg.IsCommand() {
		return false
	}

	if msg.Chat.IsPrivate() {
		return true
	}

	parts := strings.Split(msg.RawCommand(), "@")
	if len(parts) <= 1 {
		return false
	}

	return strings.ToLower(parts[1]) == strings.ToLower(bot.User.Username)
}

// IsMessageMentionsMe checks that the input message mentions the current bot.
func (bot *Bot) IsMessageMentionsMe(msg *Message) bool {
	if msg == nil || bot == nil || bot.User == nil {
		return false
	}

	if bot.IsCommandToMe(msg) {
		return true
	}

	var entities []MessageEntity
	switch {
	case msg.HasMentions():
		entities = msg.Entities
	case msg.HasCaptionMentions():
		entities = msg.CaptionEntities
	}

	for _, entity := range entities {
		if entity.IsMention() {
			if bot.ID == entity.User.ID {
				return true
			}
		}
	}

	return false
}

// IsForwardMentionsMe checks that the input forwarded message mentions the
// current bot.
func (bot *Bot) IsForwardMentionsMe(msg *Message) bool {
	return msg.IsForward() && bot.IsMessageMentionsMe(msg)
}

// IsReplyMentionsMe checks that the input message mentions the current bot.
func (bot *Bot) IsReplyMentionsMe(msg *Message) bool {
	return msg.IsReply() && bot.IsMessageMentionsMe(msg.ReplyToMessage)
}

// IsMessageToMe checks that the input message is addressed to the current bot.
func (bot *Bot) IsMessageToMe(msg *Message) bool {
	if msg == nil || msg.Chat == nil {
		return false
	}

	if msg.Chat.IsPrivate() || bot.IsCommandToMe(msg) || bot.IsReplyToMe(msg) || bot.IsMessageMentionsMe(msg) {
		return true
	}

	return false
}

// NewFileURL creates a url.URL to file with path getted from GetFile method.
func (bot *Bot) NewFileURL(filePath string) *url.URL {
	if bot == nil || bot.AccessToken == "" || filePath == "" {
		return nil
	}

	result := defaultURI
	result.Path = fmt.Sprint("/file/bot", bot.AccessToken, "/", filePath)

	return result
}

// NewRedirectURL creates new url.URL for redirecting from one chat to another.
func (bot *Bot) NewRedirectURL(param string, group bool) *url.URL {
	if bot == nil || bot.User == nil || bot.User.Username == "" {
		return nil
	}

	link := &url.URL{
		Scheme: "https",
		Host:   "t.me",
		Path:   bot.User.Username,
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
