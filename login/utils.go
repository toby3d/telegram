package login

import "time"

// FullName return user first name only or full name if last name is present.
func (user *User) FullName() string {
	if user == nil {
		return ""
	}

	if user.HasLastName() {
		return user.FirstName + " " + user.LastName
	}

	return user.FirstName
}

// AuthTime convert AuthDate field into time.Time.
func (user *User) AuthTime() *time.Time {
	if user == nil {
		return nil
	}

	t := time.Unix(user.AuthDate, 0)
	return &t
}

// HaveLastName checks what the current user has a LastName.
func (u *User) HasLastName() bool {
	return u != nil && u.LastName != ""
}

// HaveUsername checks what the current user has a username.
func (u *User) HasUsername() bool {
	return u != nil && u.Username != ""
}
