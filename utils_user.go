package telegram

import (
	"fmt"

	"golang.org/x/text/language"
)

// Language parse LanguageCode of current user and returns language.Tag.
func (u *User) Language() *language.Tag {
	if u == nil {
		return nil
	}

	tag, err := language.Parse(u.LanguageCode)
	if err != nil {
		return nil
	}

	return &tag
}

// FullName returns the full name of user or FirstName if LastName is not
// available.
func (u *User) FullName() string {
	if u == nil {
		return ""
	}

	if u.LastName != "" {
		return fmt.Sprintln(u.FirstName, u.LastName)
	}

	return u.FirstName
}
