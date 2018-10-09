package telegram

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Language parse LanguageCode of current user and returns language.Tag.
func (u *User) Language() language.Tag {
	if u == nil {
		return language.Und
	}

	tag, err := language.Parse(u.LanguageCode)
	if err != nil {
		return language.Und
	}

	return tag
}

// NewPrinter create simple message.Printer with User.Language() by default.
func (u *User) NewPrinter() *message.Printer {
	return message.NewPrinter(u.Language())
}

// FullName returns the full name of user or FirstName if LastName is not
// available.
func (u *User) FullName() string {
	if u == nil {
		return ""
	}

	if u.HasLastName() {
		return u.FirstName + " " + u.LastName
	}

	return u.FirstName
}

// HaveLastName checks what the current user has a LastName.
func (u *User) HasLastName() bool {
	return u != nil && u.LastName != ""
}

// HaveUsername checks what the current user has a username.
func (u *User) HasUsername() bool {
	return u != nil && u.Username != ""
}
