package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetGameHighScoresParameters represents data for GetGameHighScores method.
type GetGameHighScoresParameters struct {
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
