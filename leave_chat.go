package telegram

import json "github.com/pquerna/ffjson/ffjson"

type LeaveChatParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// LeaveChat leave a group, supergroup or channel. Returns True on success.
func (bot *Bot) LeaveChat(chatID int64) (bool, error) {
	dst, err := json.Marshal(&LeaveChatParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "leaveChat")
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
