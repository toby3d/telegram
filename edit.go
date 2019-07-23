//go:generate ffjson $GOFILE
package telegram

import json "github.com/pquerna/ffjson/ffjson"

type (
	// EditMessageLiveLocationParameters represents data for EditMessageLiveLocation
	// method.
	EditMessageLiveLocationParameters struct {
		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat or username of the target channel (in the format
		// @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// Latitude of new location
		Latitude float32 `json:"latitude"`

		// Longitude of new location
		Longitude float32 `json:"longitude"`

		// A JSON-serialized object for a new inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageTextParameters represents data for EditMessageText method.
	EditMessageTextParameters struct {
		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat or username of the target channel (in the format
		// @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// New text of the message
		Text string `json:"text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode,omitempty"`

		// Disables link previews for links in this message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

		// A JSON-serialized object for an inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageCaptionParameters represents data for EditMessageCaption method.
	EditMessageCaptionParameters struct {
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

	// EditMessageMediaParameters represents data for EditMessageMedia method.
	EditMessageMediaParameters struct {
		// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// A JSON-serialized object for a new media content of the message
		Media interface{} `json:"media"`

		// A JSON-serialized object for a new inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageReplyMarkupParameters represents data for EditMessageReplyMarkup method.
	EditMessageReplyMarkupParameters struct {
		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat or username of the target channel (in the format
		// @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}
)

// NewLiveLocation creates EditMessageLiveLocationParameters only with required
// parameters.
func NewLiveLocation(latitude, longitude float32) *EditMessageLiveLocationParameters {
	return &EditMessageLiveLocationParameters{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewMessageText creates EditMessageTextParameters only with required parameters.
func NewMessageText(text string) *EditMessageTextParameters {
	return &EditMessageTextParameters{
		Text: text,
	}
}

// EditMessageLiveLocation edit live location messages sent by the bot or via the
// bot (for inline bots). A location can be edited until its live_period expires
// or editing is explicitly disabled by a call to stopMessageLiveLocation. On
// success, if the edited message was sent by the bot, the edited Message is
// returned, otherwise True is returned.
func (bot *Bot) EditMessageLiveLocation(params *EditMessageLiveLocationParameters) (msg *Message, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageLiveLocation)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.UnmarshalFast(*resp.Result, msg)
	return
}

// EditMessageText edit text and game messages sent by the bot or via the bot
// (for inline bots). On success, if edited message is sent by the bot, the
// edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageText(params *EditMessageTextParameters) (msg *Message, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageText)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.UnmarshalFast(*resp.Result, msg)
	return
}

// EditMessageCaption edit captions of messages sent by the bot or via the bot
// (for inline bots). On success, if edited message is sent by the bot, the
// edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageCaption(params *EditMessageCaptionParameters) (msg *Message, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageCaption)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.UnmarshalFast(*resp.Result, msg)
	return
}

// EditMessageMedia edit audio, document, photo, or video messages. If a message
// is a part of a message album, then it can be edited only to a photo or a video.
// Otherwise, message type can be changed arbitrarily. When inline message is
// edited, new file can't be uploaded. Use previously uploaded file via its
// file_id or specify a URL. On success, if the edited message was sent by the
// bot, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageMedia(emmp *EditMessageMediaParameters) (msg *Message, err error) {
	var src []byte
	src, err = json.MarshalFast(emmp)
	if err != nil {
		return
	}

	resp, err := b.request(src, MethodEditMessageMedia)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.UnmarshalFast(*resp.Result, msg)
	return
}

// EditMessageReplyMarkup edit only the reply markup of messages sent by the bot
// or via the bot (for inline bots). On success, if edited message is sent by the
// bot, the edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageReplyMarkup(params *EditMessageReplyMarkupParameters) (msg *Message, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageReplyMarkup)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.UnmarshalFast(*resp.Result, msg)
	return
}
