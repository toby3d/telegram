package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetChatParameters represents data for GetChat method.
type GetChatParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChat get up to date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel,
// etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chatID int64) (chat *Chat, err error) {
	dst, err := json.Marshal(&GetChatParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChat)
	if err != nil {
		return
	}

	chat = new(Chat)
	err = json.Unmarshal(*resp.Result, chat)
	return
}
