package telegram

import "strings"

func (bot *Bot) IsMessageFromMe(msg *Message) bool {
	if msg == nil ||
		bot == nil {
		return false
	}

	if msg.From == nil ||
		bot.Self == nil {
		return false
	}

	return msg.From.ID == bot.Self.ID
}

func (bot *Bot) IsForwardFromMe(msg *Message) bool {
	if !msg.IsForward() {
		return false
	}

	if bot == nil {
		return false
	}

	if bot.Self == nil {
		return false
	}

	return msg.ForwardFrom.ID == bot.Self.ID
}

func (bot *Bot) IsReplyToMe(msg *Message) bool {
	if msg.Chat.IsPrivate() {
		return true
	}

	if !msg.IsReply() {
		return false
	}

	return bot.IsMessageFromMe(msg.ReplyToMessage)
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

	return strings.ToLower(parts[1]) == strings.ToLower(bot.Self.Username)
}

func (bot *Bot) IsMessageMentionsMe(msg *Message) bool {
	if msg == nil ||
		bot == nil {
		return false
	}

	if bot.Self == nil {
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
			if bot.Self.ID == entity.User.ID {
				return true
			}
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
	if msg == nil {
		return false
	}

	if msg.Chat == nil {
		return false
	}

	if msg.Chat.IsPrivate() ||
		bot.IsCommandToMe(msg) ||
		bot.IsReplyToMe(msg) ||
		bot.IsMessageMentionsMe(msg) {
		return true
	}

	return false
}
