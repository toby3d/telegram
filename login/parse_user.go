package login

import (
	"net/url"
	"strconv"
)

// User contains data about authenticated user.
type User struct {
	ID        int    `json:"id"`
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	LastName  string `json:"last_name,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	Username  string `json:"username,omitempty"`
}

// ParseUser create User structure from input url.Values.
func ParseUser(src url.Values) (u *User, err error) {
	u = new(User)

	var ad int
	ad, err = strconv.Atoi(src.Get(KeyAuthDate))
	if err != nil {
		return
	}

	u.ID, err = strconv.Atoi(src.Get(KeyID))
	if err != nil {
		return
	}

	u.AuthDate = int64(ad)
	u.FirstName = src.Get(KeyFirstName)
	u.Hash = src.Get(KeyHash)
	u.LastName = src.Get(KeyLastName)
	u.PhotoURL = src.Get(KeyPhotoURL)
	u.Username = src.Get(KeyUsername)

	return
}
