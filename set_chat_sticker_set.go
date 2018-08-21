package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SetChatStickerSetParameters represents data for SetChatStickerSet method.
type SetChatStickerSetParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet set a new group sticker set for a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success.
func (bot *Bot) SetChatStickerSet(chatID int64, stickerSetName string) (ok bool, err error) {
	dst, err := json.Marshal(&SetChatStickerSetParameters{
		ChatID:         chatID,
		StickerSetName: stickerSetName,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSetChatStickerSet)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
