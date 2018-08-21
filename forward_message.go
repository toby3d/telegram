package telegram

import json "github.com/pquerna/ffjson/ffjson"

// ForwardMessageParameters represents data for ForwardMessage method.
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

// NewForwardMessage creates ForwardMessageParameters only with reqired parameters.
func NewForwardMessage(from, to int64, messageID int) *ForwardMessageParameters {
	return &ForwardMessageParameters{
		FromChatID: from,
		ChatID:     to,
		MessageID:  messageID,
	}
}

// ForwardMessage forward messages of any kind. On success, the sent Message is returned.
func (bot *Bot) ForwardMessage(params *ForwardMessageParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodForwardMessage)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}
