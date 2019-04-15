package telegram

import json "github.com/pquerna/ffjson/ffjson"

type SendPollConfig struct {
	// Unique identifier for the target chat. A native poll can't be sent to a private chat.
	ChatID int64 `json:"chat_id"`

	// Poll question, 1-255 characters
	Question string `json:"question"`

	// List of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func NewPoll(chatID int64, question string, options ...string) SendPollConfig {
	return SendPollConfig{
		ChatID:   chatID,
		Question: question,
		Options:  options,
	}
}

// SendPoll send a native poll. A native poll can't be sent to a private chat. On success, the sent Message is
// returned.
func (b *Bot) SendPoll(params SendPollConfig) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := b.request(dst, MethodSendPoll)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(*resp.Result, &msg)
	return &msg, err
}
