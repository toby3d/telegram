package telegram

import "fmt"

// IsPrivate checks that the current chat is a private chat with single user.
func (chat *Chat) IsPrivate() bool {
	return chat != nil && chat.Type == ChatPrivate
}

// IsGroup checks that the current chat is a group.
func (chat *Chat) IsGroup() bool {
	return chat != nil && chat.Type == ChatGroup
}

// IsSuperGroup checks that the current chat is a supergroup.
func (chat *Chat) IsSuperGroup() bool {
	return chat != nil && chat.Type == ChatSuperGroup
}

// IsChannel checks that the current chat is a channel.
func (chat *Chat) IsChannel() bool {
	return chat != nil && chat.Type == ChatChannel
}

// HasPinnedMessage checks that the current chat has a pinned message.
func (chat *Chat) HasPinnedMessage() bool {
	return chat != nil && chat.PinnedMessage != nil
}

// HasStickerSet checks that the current chat has a sticker set.
func (chat *Chat) HasStickerSet() bool {
	return chat != nil && chat.StickerSetName != ""
}

// StickerSet return StickerSet structure if StickerSetName is available.
func (chat *Chat) StickerSet(bot *Bot) *StickerSet {
	if !chat.HasStickerSet() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(chat.StickerSetName)
	if err != nil {
		return nil
	}

	return set
}

// FullName returns the full name of chat or FirstName if LastName is not available.
func (chat *Chat) FullName() string {
	if chat == nil {
		return ""
	}

	if chat.LastName != "" {
		return fmt.Sprintln(chat.FirstName, chat.LastName)
	}

	return chat.FirstName
}
