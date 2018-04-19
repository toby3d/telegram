package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendMessageParameters represents data for SendMessage method.
type SendMessageParameters struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID int64 `json:"chat_id"`

	// Text of the message to be sent
	Text string `json:"text"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
	// fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`

	// Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// Sends the message silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline
	// keyboard, custom reply keyboard, instructions to remove reply keyboard or
	// to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// NewMessage creates SendMessageParameters only with required parameters.
func NewMessage(chatID int64, text string) *SendMessageParameters {
	return &SendMessageParameters{
		ChatID: chatID,
		Text:   text,
	}
}

// SendMessage send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(params *SendMessageParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodSendMessage)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
