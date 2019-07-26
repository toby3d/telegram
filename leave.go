package telegram

// LeaveChatParameters represents data for LeaveChat method.
type LeaveChatParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// LeaveChat leave a group, supergroup or channel. Returns True on success.
func (bot *Bot) LeaveChat(chatID int64) (bool, error) {
	dst, err := parser.Marshal(&LeaveChatParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodLeaveChat)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
