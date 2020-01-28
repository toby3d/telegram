package login

import "time"

// FullName return user first name only or full name if last name is present.
func (u User) FullName() string {
	if u.FirstName == "" {
		return ""
	}

	name := u.FirstName
	if u.HasLastName() {
		name += " " + u.LastName
	}

	return name
}

// AuthTime convert AuthDate field into time.Time.
func (u User) AuthTime() time.Time {
	if u.AuthDate == 0 {
		return time.Time{}
	}

	return time.Unix(u.AuthDate, 0)
}

// HasLastName checks what the current user has a LastName.
func (u User) HasLastName() bool { return u.LastName != "" }

// HasUsername checks what the current user has a username.
func (u User) HasUsername() bool { return u.Username != "" }

// HasPhoto checks what the current user has a photo.
func (u User) HasPhoto() bool { return u.PhotoURL != "" }
