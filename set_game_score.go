package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SetGameScoreParameters represents data for SetGameScore method.
type SetGameScoreParameters struct {
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
func (bot *Bot) SetGameScore(params *SetGameScoreParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodSetGameScore)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
