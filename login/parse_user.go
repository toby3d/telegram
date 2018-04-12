package login

import (
	"net/url"
	"strconv"
)

// User contains data about authenticated user.
type User struct {
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	ID        int    `json:"id"`
	LastName  string `json:"last_name,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	Username  string `json:"username,omitempty"`
}

// ParseUser create User structure from input url.Values.
func ParseUser(src url.Values) (*User, error) {
	authDate, err := strconv.Atoi(src.Get(KeyAuthDate))
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(src.Get(KeyID))
	if err != nil {
		return nil, err
	}

	return &User{
		AuthDate:  int64(authDate),
		FirstName: src.Get(KeyFirstName),
		Hash:      src.Get(KeyHash),
		ID:        id,
		LastName:  src.Get(KeyLastName),
		PhotoURL:  src.Get(KeyPhotoURL),
		Username:  src.Get(KeyUsername),
	}, nil
}
