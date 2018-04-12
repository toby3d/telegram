package telegram

import json "github.com/pquerna/ffjson/ffjson"

type DeleteChatStickerSetParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// DeleteChatStickerSet delete a group sticker set from a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success.
func (bot *Bot) DeleteChatStickerSet(chatID int64) (bool, error) {
	dst, err := json.Marshal(&DeleteChatStickerSetParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteChatStickerSet)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
