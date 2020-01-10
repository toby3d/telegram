package telegram

type (
	// EditMessageTextParameters represents data for EditMessageText method.
	EditMessageText struct {
		// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// New text of the message
		Text string `json:"text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode,omitempty"`

		// Disables link previews for links in this message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

		// A JSON-serialized object for an inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageCaptionParameters represents data for EditMessageCaption method.
	EditMessageCaption struct {
		// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// New caption of the message
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// A JSON-serialized object for an inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageMediaParameters represents data for EditMessageMedia method.
	EditMessageMedia struct {
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
	EditMessageReplyMarkup struct {
		// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	StopPoll struct {
		// Unique identifier for the target chat. A native poll can't be sent to a private chat.
		ChatID int64 `json:"chat_id"`

		// Identifier of the original message with the poll
		MessageID int `json:"message_id"`

		// A JSON-serialized object for a new message inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// DeleteMessageParameters represents data for DeleteMessage method.
	DeleteMessage struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Identifier of the message to delete
		MessageID int `json:"message_id"`
	}
)

// EditMessageText edit text and game messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (b Bot) EditMessageText(p *EditMessageText) (*Message, error) {
	src, err := b.Do(MethodEditMessageText, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(Message)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// EditMessageCaption edit captions of messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (b Bot) EditMessageCaption(p *EditMessageCaption) (*Message, error) {
	src, err := b.Do(MethodEditMessageCaption, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(Message)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// EditMessageMedia edit audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (b Bot) EditMessageMedia(p EditMessageMedia) (*Message, error) {
	src, err := b.Do(MethodEditMessageMedia, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(Message)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// EditMessageReplyMarkup edit only the reply markup of messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (b Bot) EditMessageReplyMarkup(p EditMessageReplyMarkup) (*Message, error) {
	src, err := b.Do(MethodEditMessageReplyMarkup, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(Message)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// StopPoll stop a poll which was sent by the bot. On success, the stopped Poll with the final results is returned.
func (b Bot) StopPoll(p StopPoll) (*Poll, error) {
	src, err := b.Do(MethodStopPoll, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(Poll)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteMessage delete a message, including service messages, with the following limitations:
//
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
//
// Returns True on success.
func (b Bot) DeleteMessage(cid int64, mid int) (bool, error) {
	src, err := b.Do(MethodDeleteMessage, DeleteMessage{ChatID: cid, MessageID: mid})
	if err != nil {
		return false, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return false, err
	}

	var result bool
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return false, err
	}

	return result, nil
}
