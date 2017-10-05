package telegram

func NewReplyKeyboard(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{Keyboard: keyboard, ResizeKeyboard: true}
}

func NewKeyboardButtonsRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

func NewKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text}
}

func NewKeyboardButtonContact(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestContact: true}
}

func NewKeyboardButtonLocation(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestLocation: true}
}

func NewInlineKeyboard(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

func NewInlineKeyboardButtonsRow(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

func NewInlineKeyboardButtonURL(text, url string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, URL: url}
}

func NewInlineKeyboardButtonData(text, data string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackData: data}
}

func NewInlineKeyboardButtonSwitch(text, query string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQuery: query}
}

func NewInlineKeyboardButtonSwitchSelf(text, query string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: query}
}

func NewInlineKeyboardButtonGame(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackGame: &CallbackGame{}}
}

func NewInlineKeyboardButtonPay(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, Pay: true}
}
