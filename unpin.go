package telegram

// UnpinChatMessageParameters represents data for UnpinChatMessage method.
type UnpinChatMessageParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// UnpinChatMessage unpin a message in a supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
func (bot *Bot) UnpinChatMessage(chatID int64) (ok bool, err error) {
	dst, err := parser.Marshal(&UnpinChatMessageParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodUnpinChatMessage)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
}
