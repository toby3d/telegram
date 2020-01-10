package telegram

type (
	// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
	Game struct {
		// Title of the game
		Title string `json:"title"`

		// Description of the game
		Description string `json:"description"`

		// Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
		Text string `json:"text,omitempty"`

		// Photo that will be displayed in the game message in chats.
		Photo []*PhotoSize `json:"photo"`

		// Special entities that appear in text, such as usernames, URLs, bot commands, etc.
		TextEntities []*MessageEntity `json:"text_entities,omitempty"`

		// Animation that will be displayed in the game message in chats. Upload via BotFather
		Animation *Animation `json:"animation,omitempty"`
	}

	// CallbackGame a placeholder, currently holds no information. Use BotFather to set up your game.
	CallbackGame struct{}

	// GameHighScore represents one row of the high scores table for a game.
	GameHighScore struct {
		// Position in high score table for the game
		Position int `json:"position"`

		// Score
		Score int `json:"score"`

		// User
		User *User `json:"user"`
	}

	// SendGameParameters represents data for SendGame method.
	SendGame struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
		GameShortName string `json:"game_short_name"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message.
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SetGameScoreParameters represents data for SetGameScore method.
	SetGameScore struct {
		// User identifier
		UserID int `json:"user_id"`

		// New score, must be non-negative
		Score int `json:"score"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
		Force bool `json:"force,omitempty"`

		// Pass True, if the game message should not be automatically edited to include the current scoreboard
		DisableEditMessage bool `json:"disable_edit_message,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}

	// GetGameHighScoresParameters represents data for GetGameHighScores method.
	GetGameHighScores struct {
		// Target user id
		UserID int `json:"user_id"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int `json:"message_id,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}
)

// SendGame send a game. On success, the sent Message is returned.
func (b Bot) SendGame(p SendGame) (*Message, error) {
	src, err := b.Do(MethodSendGame, p)
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

// SetGameScore set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (b Bot) SetGameScore(p SetGameScore) (*Message, error) {
	src, err := b.Do(MethodSetGameScore, p)
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

// GetGameHighScores get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
func (b Bot) GetGameHighScores(p GetGameHighScores) ([]*GameHighScore, error) {
	src, err := b.Do(MethodGetGameHighScores, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := make([]*GameHighScore, 0)
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}
