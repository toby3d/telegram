package telegram

// UnbanChatMemberParameters represents data for UnbanChatMember method.
type UnbanChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	UserID int `json:"user_id"`
}

// UnbanChatMember unban a previously kicked user in a supergroup or channel. The
// user will not return to the group or channel automatically, but will be able
// to join via link, etc. The bot must be an administrator for this to work.
// Returns True on success.
func (bot *Bot) UnbanChatMember(chatID int64, userID int) (ok bool, err error) {
	params := UnbanChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	}
	dst, err := parser.Marshal(&params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodUnbanChatMember)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
}
