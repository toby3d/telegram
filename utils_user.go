package telegram

import (
	"fmt"

	"golang.org/x/text/language"
)

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

func (user *User) FullName() string {
	if user == nil {
		return ""
	}

	if user.LastName != "" {
		return fmt.Sprint(user.FirstName, " ", user.LastName)
	}

	return user.FirstName
}
