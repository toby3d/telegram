package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendVenueParameters represents data for SendVenue method.
type SendVenueParameters struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"`

	// Latitude of the venue
	Latitude float32 `json:"latitude"`

	// Longitude of the venue
	Longitude float32 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Foursquare type of the venue, if known. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or
	// "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`

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

// NewVenue creates SendVenueParameters only with required parameters.
func NewVenue(chatID int64, latitude, longitude float32, title, address string) *SendVenueParameters {
	return &SendVenueParameters{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// SendVenue send information about a venue. On success, the sent Message is returned.
func (bot *Bot) SendVenue(params *SendVenueParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendVenue)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}
