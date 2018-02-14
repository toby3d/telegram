package login

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CheckAuthorization verify the authentication and the integrity of the data
// received by comparing the received hash parameter with the hexadecimal
// representation of the HMAC-SHA-256 signature of the data-check-string with the
// SHA256 hash of the bot's token used as a secret key.
func (user *User) CheckAuthorization(botToken string) (bool, error) {
	dataCheckString := fmt.Sprint(
		"auth_date=", user.AuthDate.Unix(),
		"\n", "first_name=", user.FirstName,
		// Eliminate 'hash' to avoid recursion and incorrect data validation.
		"\n", "id=", user.ID,
	)

	// Add optional values if exist
	if user.LastName != "" {
		dataCheckString += fmt.Sprint("\n", "last_name=", user.LastName)
	}
	if user.PhotoURL != "" {
		dataCheckString += fmt.Sprint("\n", "photo_url=", user.PhotoURL)
	}
	if user.Username != "" {
		dataCheckString += fmt.Sprint("\n", "username=", user.Username)
	}

	secretKey := sha256.Sum256([]byte(botToken))
	hash := hmac.New(sha256.New, secretKey[0:])
	_, err := hash.Write([]byte(dataCheckString))
	return hex.EncodeToString(hash.Sum(nil)) == user.Hash, err
}
