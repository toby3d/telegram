package login

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	http "github.com/valyala/fasthttp"
)

type (
	Widget struct {
		accessToken string
	}

	// User contains data about authenticated user.
	User struct {
		ID        int    `json:"id"`
		AuthDate  int64  `json:"auth_date"`
		FirstName string `json:"first_name"`
		Hash      string `json:"hash"`
		LastName  string `json:"last_name,omitempty"`
		PhotoURL  string `json:"photo_url,omitempty"`
		Username  string `json:"username,omitempty"`
	}
)

// Key represents available and supported query arguments keys.
const (
	KeyAuthDate  = "auth_date"
	KeyFirstName = "first_name"
	KeyHash      = "hash"
	KeyID        = "id"
	KeyLastName  = "last_name"
	KeyPhotoURL  = "photo_url"
	KeyUsername  = "username"
)

func NewWidget(accessToken string) *Widget {
	return &Widget{accessToken: accessToken}
}

// CheckAuthorization verify the authentication and the integrity of the data
// received by comparing the received hash parameter with the hexadecimal
// representation of the HMAC-SHA-256 signature of the data-check-string with the
// SHA256 hash of the bot's token used as a secret key.
func (w *Widget) CheckAuthorization(u User) (bool, error) {
	hash, err := w.GenerateHash(u)
	return hash == u.Hash, err
}

func (w *Widget) GenerateHash(u User) (string, error) {
	a := http.AcquireArgs()
	defer http.ReleaseArgs(a)

	// WARN: do not change order of this args, it must be alphabetical
	a.SetUint(KeyAuthDate, int(u.AuthDate))
	a.Set(KeyFirstName, u.FirstName)
	a.SetUint(KeyID, u.ID)

	if u.LastName != "" {
		a.Set(KeyLastName, u.LastName)
	}

	if u.PhotoURL != "" {
		a.Set(KeyPhotoURL, u.PhotoURL)
	}

	if u.Username != "" {
		a.Set(KeyUsername, u.Username)
	}

	secretKey := sha256.Sum256([]byte(w.accessToken))
	h := hmac.New(sha256.New, secretKey[0:])

	if _, err := h.Write(a.QueryString()); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
