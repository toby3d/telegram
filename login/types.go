package login

import "errors"

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

var (
	// ErrUserNotDefined describes error of an unassigned structure of user.
	ErrUserNotDefined = errors.New("user is not defined")

	// ErrEmptyToken describes error of an empty access token of the bot.
	ErrEmptyToken = errors.New("empty bot access token")

	// ErrUnsupportedType describes error of unsupported input data type for
	// CheckAuthorization method.
	ErrUnsupportedType = errors.New("unsupported data type")
)
