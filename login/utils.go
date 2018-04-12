package login

import (
	"fmt"
	"time"
)

// FullName return user first name only or full name if last name is present.
func (user *User) FullName() string {
	if user == nil {
		return ""
	}

	if user.LastName != "" {
		return fmt.Sprintln(user.FirstName, user.LastName)
	}

	return user.FirstName
}

// AuthTime convert AuthDate field into time.Time.
func (user *User) AuthTime() time.Time {
	if user == nil {
		return time.Time{}
	}

	return time.Unix(user.AuthDate, 0)
}
