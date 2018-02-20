package telegram

import json "github.com/pquerna/ffjson/ffjson"

type ForwardMessageParameters struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID int64 `json:"chat_id"`

	// Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatID int64 `json:"from_chat_id"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`
}

func NewForwardMessage(from, to int64, messageID int) *ForwardMessageParameters {
	return &ForwardMessageParameters{
		FromChatID: from,
		ChatID:     to,
		MessageID:  messageID,
	}
}

// ForwardMessage forward messages of any kind. On success, the sent Message is returned.
func (bot *Bot) ForwardMessage(params *ForwardMessageParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "forwardMessage")
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
