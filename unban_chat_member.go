package telegram

import json "github.com/pquerna/ffjson/ffjson"

type UnbanChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	UserID int `json:"user_id"`
}

// UnbanChatMember unban a previously kicked user in a supergroup or channel. The
// user will not return to the group or channel automatically, but will be able
// to join via link, etc. The bot must be an administrator for this to work.
// Returns True on success.
func (bot *Bot) UnbanChatMember(chatID int64, userID int) (bool, error) {
	dst, err := json.Marshal(&UnbanChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodUnbanChatMember)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
