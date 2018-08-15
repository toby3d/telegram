package telegram

import "fmt"

// IsPrivate checks that the current chat is a private chat with single user.
func (c *Chat) IsPrivate() bool {
	return c != nil && c.Type == ChatPrivate
}

// IsGroup checks that the current chat is a group.
func (c *Chat) IsGroup() bool {
	return c != nil && c.Type == ChatGroup
}

// IsSuperGroup checks that the current chat is a supergroup.
func (c *Chat) IsSuperGroup() bool {
	return c != nil && c.Type == ChatSuperGroup
}

// IsChannel checks that the current chat is a channel.
func (c *Chat) IsChannel() bool {
	return c != nil && c.Type == ChatChannel
}

// HasPinnedMessage checks that the current chat has a pinned message.
func (c *Chat) HasPinnedMessage() bool {
	return c != nil && c.PinnedMessage != nil
}

// HasStickerSet checks that the current chat has a sticker set.
func (c *Chat) HasStickerSet() bool {
	return c != nil && c.StickerSetName != ""
}

// StickerSet return StickerSet structure if StickerSetName is available.
func (c *Chat) StickerSet(bot *Bot) *StickerSet {
	if !c.HasStickerSet() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(c.StickerSetName)
	if err != nil {
		return nil
	}

	return set
}

// FullName returns the full name of chat or FirstName if LastName is not available.
func (c *Chat) FullName() string {
	if c == nil {
		return ""
	}

	if c.LastName != "" {
		return fmt.Sprintln(c.FirstName, c.LastName)
	}

	return c.FirstName
}
