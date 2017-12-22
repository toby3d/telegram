package telegram

import (
	"fmt"
	"strings"
)

func (bot *Bot) IsMessageFromMe(msg *Message) bool {
	return msg.From.ID == bot.Self.ID
}

func (bot *Bot) IsForwardFromMe(msg *Message) bool {
	return msg.IsForward() &&
		msg.ForwardFrom.ID == bot.Self.ID
}

func (bot *Bot) IsReplyToMe(msg *Message) bool {
	if msg.Chat.IsPrivate() {
		return true
	}

	return msg.IsReply() &&
		bot.IsMessageFromMe(msg.ReplyToMessage)
}

func (bot *Bot) IsCommandToMe(msg *Message) bool {
	if msg.Chat.IsPrivate() {
		return msg.IsCommand()
	}

	return msg.IsCommand() &&
		strings.HasSuffix(
			strings.ToLower(msg.Command()),
			strings.ToLower(fmt.Sprint("@", bot.Self.Username)),
		)
}

func (bot *Bot) IsMessageMentionsMe(msg *Message) bool {
	if msg.Entities == nil ||
		len(msg.Entities) <= 0 {
		return false
	}

	for _, entity := range msg.Entities {
		if entity.Type != EntityMention ||
			entity.User == nil {
			continue
		}

		if entity.User.ID == bot.Self.ID {
			return true
		}
	}

	return false
}

func (bot *Bot) IsForwardMentionsMe(msg *Message) bool {
	return msg.IsForward() &&
		bot.IsMessageMentionsMe(msg)
}

func (bot *Bot) IsReplyMentionsMe(msg *Message) bool {
	return msg.IsReply() &&
		bot.IsMessageMentionsMe(msg.ReplyToMessage)
}

func (bot *Bot) IsMessageToMe(msg *Message) bool {
	switch {
	case msg.Chat.IsPrivate(),
		bot.IsCommandToMe(msg),
		bot.IsReplyToMe(msg),
		bot.IsMessageMentionsMe(msg):
		return true
	default:
		return false
	}
}
