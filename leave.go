//go:generate ffjson $GOFILE
package telegram

import json "github.com/pquerna/ffjson/ffjson"

// LeaveChatParameters represents data for LeaveChat method.
type LeaveChatParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// LeaveChat leave a group, supergroup or channel. Returns True on success.
func (bot *Bot) LeaveChat(chatID int64) (ok bool, err error) {
	dst, err := json.MarshalFast(&LeaveChatParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodLeaveChat)
	if err != nil {
		return
	}

	err = json.UnmarshalFast(*resp.Result, &ok)
	return
}
