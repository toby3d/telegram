package telegram

import json "github.com/pquerna/ffjson/ffjson"

type EditMessageCaptionParameters struct {
	// Required if inline_message_id is not specified. Unique identifier for the
	// target chat or username of the target channel (in the format
	// @channelusername)
	ChatID int64 `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of
	// the sent message
	MessageID int `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New caption of the message
	Caption string `json:"caption,omitempty"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
	// fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaption edit captions of messages sent by the bot or via the bot
// (for inline bots). On success, if edited message is sent by the bot, the
// edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageCaption(params *EditMessageCaptionParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodEditMessageCaption)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
