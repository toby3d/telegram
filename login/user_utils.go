package login

import (
	"time"

	http "github.com/valyala/fasthttp"
)

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

func (u *User) toArgs() *http.Args {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.SetUint(KeyAuthDate, int(u.AuthDate))
	args.Set(KeyFirstName, u.FirstName)
	args.SetUint(KeyID, u.ID)
	args.Set(KeyHash, u.Hash)

	// Add optional values if exist
	if u.LastName != "" {
		args.Set(KeyLastName, u.LastName)
	}
	if u.PhotoURL != "" {
		args.Set(KeyPhotoURL, u.PhotoURL)
	}
	if u.Username != "" {
		args.Set(KeyUsername, u.Username)
	}

	return args
}
