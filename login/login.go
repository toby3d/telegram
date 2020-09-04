package login

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"

	http "github.com/valyala/fasthttp"
	"golang.org/x/text/language"
)

type (
	Config struct {
		// ClientSecret is the bot token.
		ClientSecret string

		// RedirectURL is the URL to redirect users going through the login flow.
		RedirectURL string

		// RequestWriteAccess request the permission for bot to send messages to the user.
		RequestWriteAccess bool
	}

	// User contains data about authenticated user.
	User struct {
		AuthDate  int64  `json:"auth_date"`
		FirstName string `json:"first_name"`
		Hash      string `json:"hash"`
		ID        int    `json:"id"`
		LastName  string `json:"last_name,omitempty"`
		PhotoURL  string `json:"photo_url,omitempty"`
		Username  string `json:"username,omitempty"`
	}
)

const Endpoint string = "https://oauth.telegram.org/auth"

// Key represents available and supported query arguments keys.
const (
	KeyAuthDate  string = "auth_date"
	KeyFirstName string = "first_name"
	KeyHash      string = "hash"
	KeyID        string = "id"
	KeyLastName  string = "last_name"
	KeyPhotoURL  string = "photo_url"
	KeyUsername  string = "username"
)

// ClientID returns bot ID from it's ClientSecret token.
func (c Config) ClientID() string {
	return strings.SplitN(c.ClientSecret, ":", 2)[0]
}

// AuthCodeURL returns a URL to Telegram login page that asks for permissions for the required scopes explicitly.
func (c *Config) AuthCodeURL(lang language.Tag) string {
	origin, _ := url.Parse(c.RedirectURL)

	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.Update(Endpoint)

	a := u.QueryArgs()
	a.Set("bot_id", c.ClientID())
	a.Set("origin", origin.Scheme+"://"+origin.Host)
	a.Add("embed", "0")

	if lang != language.Und {
		a.Set("lang", lang.String())
	}

	if c.RequestWriteAccess {
		a.Set("request_access", "write")
	}

	return u.String()
}

// Verify verify the authentication and the integrity of the data received by
// comparing the received hash parameter with the hexadecimal representation
// of the HMAC-SHA-256 signature of the data-check-string with the SHA256
// hash of the bot's token used as a secret key.
func (c *Config) Verify(u *User) bool {
	if u == nil || u.Hash == "" {
		return false
	}

	h, err := generateHash(c.ClientSecret, u)

	return err == nil && u.Hash == h
}

func generateHash(token string, u *User) (string, error) {
	a := http.AcquireArgs()
	defer http.ReleaseArgs(a)

	// WARN(toby3d): do not change order of this args, it must be alphabetical
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

	secretKey := sha256.Sum256([]byte(token))
	h := hmac.New(sha256.New, secretKey[0:])

	if _, err := h.Write(a.QueryString()); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
