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
func (bot *Bot) SetChatStickerSet(chatID int64, stickerSetName string) (bool, error) {
	dst, err := json.Marshal(&SetChatStickerSetParameters{
		ChatID:         chatID,
		StickerSetName: stickerSetName,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSetChatStickerSet)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
