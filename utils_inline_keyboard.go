package telegram

// NewInlineKeyboardMarkup creates a new inline keyboard markup for message.
func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewInlineKeyboardRow creates a new inline keyboard row for buttons.
func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

// NewInlineKeyboardButton creates a new inline keyboard callback button.
func NewInlineKeyboardButton(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

// NewInlineKeyboardButtonURL creates a new inline keyboard button with URL.
func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

// NewInlineKeyboardButtonSwitch creates a new inline keyboard button to make
// specific inline query in other chat.
func NewInlineKeyboardButtonSwitch(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

// NewInlineKeyboardButtonSwitchSelf creates a new inline keyboard button to make
// specific inline query in same chat.
func NewInlineKeyboardButtonSwitchSelf(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		SwitchInlineQueryCurrentChat: query,
	}
}

// NewInlineKeyboardButtonGame creates a new inline keyboard button with game
// callback.
func NewInlineKeyboardButtonGame(text string) InlineKeyboardButton {
	var game CallbackGame
	return InlineKeyboardButton{
		Text:         text,
		CallbackGame: &game,
	}
}

// NewInlineKeyboardButtonPay creates a new inline keyboard button with pay
// callback.
func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}
