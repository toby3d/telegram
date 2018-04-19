package telegram

import "fmt"

func (chat *Chat) IsPrivate() bool {
	return chat != nil && chat.Type == ChatPrivate
}

func (chat *Chat) IsGroup() bool {
	return chat != nil && chat.Type == ChatGroup
}

func (chat *Chat) IsSuperGroup() bool {
	return chat != nil && chat.Type == ChatSuperGroup
}

func (chat *Chat) IsChannel() bool {
	return chat != nil && chat.Type == ChatChannel
}

func (chat *Chat) HasPinnedMessage() bool {
	return chat != nil && chat.PinnedMessage != nil
}

func (chat *Chat) HasStickerSet() bool {
	return chat != nil && chat.StickerSetName != ""
}

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

func (chat *Chat) FullName() string {
	if chat == nil {
		return ""
	}

	if chat.LastName != "" {
		return fmt.Sprintln(chat.FirstName, chat.LastName)
	}

	return chat.FirstName
}
