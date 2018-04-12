package telegram

import "fmt"

func (chat *Chat) IsPrivate() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatPrivate
}

func (chat *Chat) IsGroup() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatGroup
}

func (chat *Chat) IsSuperGroup() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatSuperGroup
}

func (chat *Chat) IsChannel() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatChannel
}

func (chat *Chat) HasPinnedMessage() bool {
	if chat == nil {
		return false
	}

	return chat.PinnedMessage != nil
}

func (chat *Chat) HasStickerSet() bool {
	if chat == nil {
		return false
	}

	return chat.StickerSetName != ""
}

func (chat *Chat) StickerSet(bot *Bot) *StickerSet {
	if !chat.HasStickerSet() {
		return nil
	}

	if bot == nil {
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
