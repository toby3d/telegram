package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInlineQueryHasQuery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		iq := InlineQuery{Query: "sample text"}
		assert.True(t, iq.HasQuery())
	})
	t.Run("false", func(t *testing.T) {
		iq := InlineQuery{}
		assert.False(t, iq.HasQuery())
	})
}

func TestInlineQueryHasOffset(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		iq := InlineQuery{Offset: "42"}
		assert.True(t, iq.HasOffset())
	})
	t.Run("false", func(t *testing.T) {
		iq := InlineQuery{}
		assert.False(t, iq.HasOffset())
	})
}

func TestInlineQueryHasLocation(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		iq := InlineQuery{Location: &Location{Latitude: 56.085180, Longitude: 60.735150}}
		assert.True(t, iq.HasLocation())
	})
	t.Run("false", func(t *testing.T) {
		iq := InlineQuery{}
		assert.False(t, iq.HasLocation())
	})
}

func TestChosenInlineResultInlineQueryHasLocation(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		cir := ChosenInlineResult{Location: &Location{Latitude: 56.085180, Longitude: 60.735150}}
		assert.True(t, cir.HasLocation())
	})
	t.Run("false", func(t *testing.T) {
		cir := ChosenInlineResult{}
		assert.False(t, cir.HasLocation())
	})
}
