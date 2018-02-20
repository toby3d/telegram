package telegram

import json "github.com/pquerna/ffjson/ffjson"

type UnpinChatMessageParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// UnpinChatMessage unpin a message in a supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
func (bot *Bot) UnpinChatMessage(chatID int64) (bool, error) {
	dst, err := json.Marshal(&UnpinChatMessageParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "unpinChatMessage")
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
