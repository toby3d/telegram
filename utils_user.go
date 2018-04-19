package telegram

import (
	"fmt"

	"golang.org/x/text/language"
)

// Language parse LanguageCode of current user and returns language.Tag.
func (user *User) Language() *language.Tag {
	if user == nil {
		return nil
	}

	tag, err := language.Parse(user.LanguageCode)
	if err != nil {
		return nil
	}

	return &tag
}

// FullName returns the full name of user or FirstName if LastName is not
// available.
func (user *User) FullName() string {
	if user == nil {
		return ""
	}

	if user.LastName != "" {
		return fmt.Sprintln(user.FirstName, user.LastName)
	}

	return user.FirstName
}
