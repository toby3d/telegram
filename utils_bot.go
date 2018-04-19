package telegram

import (
	"fmt"
	"net/url"
	"strings"
)

type Bot struct {
	AccessToken string
	*User
}

func New(accessToken string) (*Bot, error) {
	var err error
	bot := new(Bot)
	bot.AccessToken = accessToken

	bot.User, err = bot.GetMe()
	return bot, err
}

func (bot *Bot) IsMessageFromMe(msg *Message) bool {
	return msg != nil && bot != nil && msg.From != nil && bot.User != nil && msg.From.ID == bot.ID
}

func (bot *Bot) IsForwardFromMe(msg *Message) bool {
	return msg.IsForward() && bot != nil && bot.User != nil && msg.ForwardFrom.ID == bot.ID
}

func (bot *Bot) IsReplyToMe(msg *Message) bool {
	return msg.Chat.IsPrivate() || (msg.IsReply() && bot.IsMessageFromMe(msg.ReplyToMessage))
}

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

func (bot *Bot) IsForwardMentionsMe(msg *Message) bool {
	return msg.IsForward() && bot.IsMessageMentionsMe(msg)
}

func (bot *Bot) IsReplyMentionsMe(msg *Message) bool {
	return msg.IsReply() && bot.IsMessageMentionsMe(msg.ReplyToMessage)
}

func (bot *Bot) IsMessageToMe(msg *Message) bool {
	if msg == nil || msg.Chat == nil {
		return false
	}

	if msg.Chat.IsPrivate() || bot.IsCommandToMe(msg) || bot.IsReplyToMe(msg) || bot.IsMessageMentionsMe(msg) {
		return true
	}

	return false
}

func (bot *Bot) NewFileURL(filePath string) *url.URL {
	if bot == nil || bot.AccessToken == "" || filePath == "" {
		return nil
	}

	return &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("/file/bot", bot.AccessToken, "/", filePath),
	}
}

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
