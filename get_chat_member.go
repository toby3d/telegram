package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetChatMemberParameters represents data for GetChatMember method.
type GetChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	UserID int `json:"user_id"`
}

// GetChatMember get information about a member of a chat. Returns a ChatMember
// object on success.
func (bot *Bot) GetChatMember(chatID int64, userID int) (member *ChatMember, err error) {
	dst, err := json.Marshal(&GetChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatMember)
	if err != nil {
		return
	}

	member = new(ChatMember)
	err = json.Unmarshal(*resp.Result, member)
	return
}
