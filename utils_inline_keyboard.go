package telegram

func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

func NewInlineKeyboardButton(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

func NewInlineKeyboardButtonSwitch(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

func NewInlineKeyboardButtonSwitchSelf(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		SwitchInlineQueryCurrentChat: query,
	}
}

func NewInlineKeyboardButtonGame(text string) InlineKeyboardButton {
	var game CallbackGame
	return InlineKeyboardButton{
		Text:         text,
		CallbackGame: &game,
	}
}

func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}
