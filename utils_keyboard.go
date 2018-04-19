package telegram

// NewReplyKeyboardRemove just hides keyboard.
func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewReplyKeyboardMarkup creates new keyboard markup of simple buttons.
func NewReplyKeyboardMarkup(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{Keyboard: keyboard}
}

// NewReplyKeyboardRow creates new keyboard row for buttons.
func NewReplyKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

// NewReplyKeyboardButton creates new button with custom text for sending it.
func NewReplyKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

// NewReplyKeyboardButtonContact creates new button with custom text for sending
// user contact.
func NewReplyKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// NewReplyKeyboardButtonLocation creates new button with custom text for sending
// user location.
func NewReplyKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}
