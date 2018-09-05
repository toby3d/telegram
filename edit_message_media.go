package telegram

import json "github.com/pquerna/ffjson/ffjson"

type EditMessageMediaParameters struct {
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
