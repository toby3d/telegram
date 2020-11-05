package login_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/toby3d/telegram/v5/login"
	"golang.org/x/text/language"
)

func TestClientID(t *testing.T) {
	c := login.Config{ClientSecret: "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"}
	assert.Equal(t, "123456", c.ClientID())
}

func TestAuthCodeURL(t *testing.T) {
	c := login.Config{
		ClientSecret:       "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
		RedirectURL:        "https://example.site/callback",
		RequestWriteAccess: true,
	}

	assert.Equal(t, "https://oauth.telegram.org/auth?bot_id=123456&origin=https%3A%2F%2Fexample.site"+
		"&embed=0&lang=ru&request_access=write", c.AuthCodeURL(language.Russian))
}

func TestVerify(t *testing.T) {
	c := login.Config{ClientSecret: "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"}
	assert.True(t, c.Verify(&login.User{
		ID:        123456,
		Username:  "toby3d",
		FirstName: "Maxim",
		LastName:  "Lebedev",
		PhotoURL:  "https://t.me/i/userpic/320/ABC-DEF1234ghIkl-zyx57W2v1u123ew11.jpg",
		AuthDate:  1410696795,
		Hash:      "d9b74e929cd4cfa7299031421db61949ecd49641c3b06e3a0361f593cf1fe064",
	}))
}
