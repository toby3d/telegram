package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetChatMembersCountParameters represents data for GetChatMembersCount method.
type GetChatMembersCountParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChatMembersCount get the number of members in a chat. Returns Int on
// success.
func (bot *Bot) GetChatMembersCount(chatID int64) (count int, err error) {
	dst, err := json.Marshal(&GetChatMembersCountParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatMembersCount)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &count)
	return
}
