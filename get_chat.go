package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetChatParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChat get up to date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel,
// etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chatID int64) (*Chat, error) {
	dst, err := json.Marshal(&GetChatParameters{ChatID: chatID})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "getChat")
	if err != nil {
		return nil, err
	}

	var data Chat
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
