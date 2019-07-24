package login

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, NewWidget("hackme"))
}

func TestGenerateHash(t *testing.T) {
	w := NewWidget("hackme")
	hash, err := w.GenerateHash(User{
		ID:        123,
		Username:  "toby3d",
		FirstName: "Maxim",
		LastName:  "Lebedev",
		AuthDate:  time.Now().UTC().Unix(),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCheckAuthorization(t *testing.T) {
	w := NewWidget("hackme")
	u := User{
		ID:        123,
		Username:  "toby3d",
		FirstName: "Maxim",
		LastName:  "Lebedev",
		PhotoURL:  "https://toby3d.me/avatar.jpg",
		AuthDate:  time.Now().UTC().Unix(),
	}
	t.Run("invalid", func(t *testing.T) {
		u.Hash = "wtf"
		ok, err := w.CheckAuthorization(u)
		assert.NoError(t, err)
		assert.False(t, ok)
	})
	t.Run("valid", func(t *testing.T) {
		var err error
		u.Hash, err = w.GenerateHash(u)
		assert.NoError(t, err)
		assert.NotEmpty(t, u.Hash)

		ok, err := w.CheckAuthorization(u)
		assert.NoError(t, err)
		assert.True(t, ok)
	})
}
