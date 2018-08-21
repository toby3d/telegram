package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendMediaGroupParameters represents data for SendMediaGroup method.
type SendMediaGroupParameters struct {
	// Unique identifier for the target chat.
	ChatID int64 `json:"chat_id"`

	// A JSON-serialized array describing photos and videos to be sent, must
	// include 2–10 items
	Media []interface{} `json:"media"`

	// Sends the messages silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the messages are a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
}

// NewMediaGroup creates SendMediaGroupParameters only with required parameters.
func NewMediaGroup(chatID int64, media ...interface{}) *SendMediaGroupParameters {
	return &SendMediaGroupParameters{
		ChatID: chatID,
		Media:  media,
	}
}

// SendMediaGroup send a group of photos or videos as an album. On success, an array of the sent
// Messages is returned.
func (bot *Bot) SendMediaGroup(params *SendMediaGroupParameters) (album []Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendMediaGroup)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &album)
	return
}
