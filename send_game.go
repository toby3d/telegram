package telegram

import json "github.com/pquerna/ffjson/ffjson"

type SendGameParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Short name of the game, serves as the unique identifier for the game. Set
	// up your games via Botfather.
	GameShortName string `json:"game_short_name"`

	// Sends the message silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one ‘Play
	// game_title’ button will be shown. If not empty, the first button must
	// launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewGame(chatID int64, gameShortName string) *SendGameParameters {
	return &SendGameParameters{
		ChatID:        chatID,
		GameShortName: gameShortName,
	}
}

// SendGame send a game. On success, the sent Message is returned.
func (bot *Bot) SendGame(params *SendGameParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "sendGame")
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
