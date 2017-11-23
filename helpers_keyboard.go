package telegram

func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

func NewReplyKeyboardMarkup(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}

func NewReplyKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

func NewReplyKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

func NewReplyKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

func NewReplyKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}
