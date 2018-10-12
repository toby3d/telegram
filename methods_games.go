package telegram

import json "github.com/pquerna/ffjson/ffjson"

type (
	// SendGameParameters represents data for SendGame method.
	SendGameParameters struct {
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

	// SetGameScoreParameters represents data for SetGameScore method.
	SetGameScoreParameters struct {
		// User identifier
		UserID int `json:"user_id"`

		// New score, must be non-negative
		Score int `json:"score"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Pass True, if the high score is allowed to decrease. This can be useful
		// when fixing mistakes or banning cheaters
		Force bool `json:"force,omitempty"`

		// Pass True, if the game message should not be automatically edited to
		// include the current scoreboard
		DisableEditMessage bool `json:"disable_edit_message,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}

	// GetGameHighScoresParameters represents data for GetGameHighScores method.
	GetGameHighScoresParameters struct {
		// Target user id
		UserID int `json:"user_id"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}
)

// NewGame creates SendGameParameters only with required parameters.
func NewGame(chatID int64, gameShortName string) *SendGameParameters {
	return &SendGameParameters{
		ChatID:        chatID,
		GameShortName: gameShortName,
	}
}

// SendGame send a game. On success, the sent Message is returned.
func (bot *Bot) SendGame(params *SendGameParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendGame)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}

// NewGameScore creates SetGameScoreParameters only with required parameters.
func NewGameScore(userID, score int) *SetGameScoreParameters {
	return &SetGameScoreParameters{
		UserID: userID,
		Score:  score,
	}
}

// SetGameScore set the score of the specified user in a game. On success, if the
// message was sent by the bot, returns the edited Message, otherwise returns
// True. Returns an error, if the new score is not greater than the user's
// current score in the chat and force is False.
func (bot *Bot) SetGameScore(params *SetGameScoreParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSetGameScore)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}

// NewGameHighScores creates GetGameHighScoresParameters only with required parameters.
func NewGameHighScores(userID int) *GetGameHighScoresParameters {
	return &GetGameHighScoresParameters{
		UserID: userID,
	}
}

// GetGameHighScores get data for high score tables. Will return the score of the
// specified user and several of his neighbors in a game. On success, returns an
// Array of GameHighScore objects.
func (bot *Bot) GetGameHighScores(params *GetGameHighScoresParameters) (scores []GameHighScore, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetGameHighScores)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &scores)
	return
}
