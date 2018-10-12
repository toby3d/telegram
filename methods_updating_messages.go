package telegram

import json "github.com/pquerna/ffjson/ffjson"

type (
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

	// DeleteMessageParameters represents data for DeleteMessage method.
	DeleteMessageParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Identifier of the message to delete
		MessageID int `json:"message_id"`
	}
)

// NewMessageText creates EditMessageTextParameters only with required parameters.
func NewMessageText(text string) *EditMessageTextParameters {
	return &EditMessageTextParameters{
		Text: text,
	}
}

// EditMessageText edit text and game messages sent by the bot or via the bot
// (for inline bots). On success, if edited message is sent by the bot, the
// edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageText(params *EditMessageTextParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageText)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}

// EditMessageCaption edit captions of messages sent by the bot or via the bot
// (for inline bots). On success, if edited message is sent by the bot, the
// edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageCaption(params *EditMessageCaptionParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageCaption)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
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
	src, err = json.Marshal(emmp)
	if err != nil {
		return
	}

	resp, err := b.request(src, MethodEditMessageMedia)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}

// EditMessageReplyMarkup edit only the reply markup of messages sent by the bot
// or via the bot (for inline bots). On success, if edited message is sent by the
// bot, the edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageReplyMarkup(params *EditMessageReplyMarkupParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageReplyMarkup)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}

// DeleteMessage delete a message, including service messages, with the following
// limitations: A message can only be deleted if it was sent less than 48 hours
// ago; Bots can delete outgoing messages in groups and supergroups; Bots granted
// can_post_messages permissions can delete outgoing messages in channels; If the
// bot is an administrator of a group, it can delete any message there; If the
// bot has can_delete_messages permission in a supergroup or a channel, it can
// delete any message there. Returns True on success.
func (bot *Bot) DeleteMessage(chatID int64, messageID int) (ok bool, err error) {
	dst, err := json.Marshal(&DeleteMessageParameters{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodDeleteMessage)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
