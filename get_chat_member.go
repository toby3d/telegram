package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	UserID int `json:"user_id"`
}

// GetChatMember get information about a member of a chat. Returns a ChatMember
// object on success.
func (bot *Bot) GetChatMember(chatID int64, userID int) (*ChatMember, error) {
	dst, err := json.Marshal(&GetChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "getChatMember")
	if err != nil {
		return nil, err
	}

	var data ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
