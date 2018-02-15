package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetChatMembersCountParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChatMembersCount get the number of members in a chat. Returns Int on
// success.
func (bot *Bot) GetChatMembersCount(chatID int64) (int, error) {
	dst, err := json.Marshal(&GetChatMembersCountParameters{ChatID: chatID})
	if err != nil {
		return 0, err
	}

	resp, err := bot.request(dst, "getChatMembersCount")
	if err != nil {
		return 0, err
	}

	var data int
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
