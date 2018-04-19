package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendContactParameters represents data for SendContact method.
type SendContactParameters struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Contact's last name
	LastName string `json:"last_name"`

	// Sends the message silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
	// price' button will be shown. If not empty, the first button must be a Pay
	// button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// NewContact creates SendContactParameters only with required parameters.
func NewContact(chatID int64, phoneNumber, firstName string) *SendContactParameters {
	return &SendContactParameters{
		ChatID:      chatID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// SendContact send phone contacts. On success, the sent Message is returned.
func (bot *Bot) SendContact(params *SendContactParameters) (*Message, error) {
	dst, err := json.Marshal(*params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodSendContact)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
