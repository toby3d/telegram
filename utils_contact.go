package telegram

import (
	"fmt"
	"strings"
)

// FullName returns the full name of contact or FirstName if LastName is not
// available.
func (c *Contact) FullName() string {
	if c == nil {
		return ""
	}

	if c.HasLastName() {
		return fmt.Sprint(c.FirstName, " ", c.LastName)
	}

	return c.FirstName
}

// HaveLastName checks what the current contact has a LastName.
func (c *Contact) HasLastName() bool {
	return c != nil && !strings.EqualFold(c.LastName, "")
}

func (c *Contact) HasTelegram() bool {
	return c != nil && c.UserID != 0
}

func (c *Contact) HasVCard() bool {
	return c != nil && !strings.EqualFold(c.VCard, "")
}
