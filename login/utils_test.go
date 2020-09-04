package login

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserFullName(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var u User
		assert.Empty(t, u.FullName())
	})
	t.Run("first name", func(t *testing.T) {
		u := User{
			FirstName: "Maxim",
		}
		assert.Equal(t, u.FirstName, u.FullName())
	})
	t.Run("first & last name", func(t *testing.T) {
		u := User{
			FirstName: "Maxim",
			LastName:  "Lebedev",
		}
		assert.Equal(t, u.FirstName+" "+u.LastName, u.FullName())
	})
}

func TestUserAuthTime(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var u User
		assert.True(t, u.AuthTime().IsZero())
	})
	t.Run("exists", func(t *testing.T) {
		u := User{AuthDate: time.Now().UTC().Unix()}
		assert.False(t, u.AuthTime().IsZero())
	})
}

func TestUserhasLastName(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		var u User
		assert.False(t, u.HasLastName())
	})
	t.Run("true", func(t *testing.T) {
		u := User{LastName: "Lebedev"}
		assert.True(t, u.HasLastName())
	})
}

func TestUserHasUsername(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		var u User
		assert.False(t, u.HasUsername())
	})
	t.Run("true", func(t *testing.T) {
		u := User{Username: "toby3d"}
		assert.True(t, u.HasUsername())
	})
}

func TestUserHasPhoto(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var u User
		assert.False(t, u.HasPhoto())
	})
	t.Run("exists", func(t *testing.T) {
		u := User{PhotoURL: "https://toby3d.me/avatar.jpg"}
		assert.True(t, u.HasPhoto())
	})
}
