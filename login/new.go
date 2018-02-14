package login

import (
	"net/url"
	"strconv"
	"time"
)

// User contains data about authenticated user.
type User struct {
	AuthDate  time.Time `json:"auth_date"`
	FirstName string    `json:"first_name"`
	Hash      string    `json:"hash"`
	ID        int       `json:"id"`
	LastName  string    `json:"last_name,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty"`
	Username  string    `json:"username,omitempty"`
}

// New create User structure from input url.Values.
func New(src url.Values) (*User, error) {
	authDate, err := strconv.Atoi(src.Get("auth_date"))
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(src.Get("id"))
	if err != nil {
		return nil, err
	}

	return &User{
		AuthDate:  time.Unix(int64(authDate), 0),
		FirstName: src.Get("first_name"),
		Hash:      src.Get("hash"),
		ID:        id,
		LastName:  src.Get("last_name"),
		PhotoURL:  src.Get("photo_url"),
		Username:  src.Get("username"),
	}, nil
}
