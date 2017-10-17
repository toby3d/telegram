package telegram

import json "github.com/pquerna/ffjson/ffjson"

type EditMessageLiveLocationParameters struct {
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

func NewLiveLocation(latitude, longitude float32) *EditMessageLiveLocationParameters {
	return &EditMessageLiveLocationParameters{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// EditMessageLiveLocation edit live location messages sent by the bot or via the
// bot (for inline bots). A location can be edited until its live_period expires
// or editing is explicitly disabled by a call to stopMessageLiveLocation. On
// success, if the edited message was sent by the bot, the edited Message is
// returned, otherwise True is returned.
func (bot *Bot) EditMessageLiveLocation(params *EditMessageLiveLocationParameters) (*Message, error) {
	dst, err := json.Marshal(*params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "editMessageLiveLocation", nil)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
