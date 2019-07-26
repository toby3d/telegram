package telegram

type StopPollConfig struct {
	// Unique identifier for the target chat. A native poll can't be sent to a private chat.
	ChatID int64 `json:"chat_id"`

	// Identifier of the original message with the poll
	MessageID int `json:"message_id"`

	// A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (b *Bot) StopPoll(params StopPollConfig) (*Poll, error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := b.request(dst, MethodStopPoll)
	if err != nil {
		return nil, err
	}

	var poll Poll
	err = parser.Unmarshal(resp.Result, &poll)
	return &poll, err
}
