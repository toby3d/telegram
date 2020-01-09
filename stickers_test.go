package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStickerInSet(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := Sticker{SetName: "HotCherry"}
		assert.True(t, s.InSet())
	})
	t.Run("false", func(t *testing.T) {
		s := Sticker{}
		assert.False(t, s.InSet())
	})
}

func TestStickerHasThumb(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := Sticker{Thumb: &PhotoSize{FileID: "abc"}}
		assert.True(t, s.HasThumb())
	})
	t.Run("false", func(t *testing.T) {
		s := Sticker{}
		assert.False(t, s.HasThumb())
	})
}

func TestStickerIsMask(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := Sticker{MaskPosition: &MaskPosition{Point: PointEyes}}
		assert.True(t, s.IsMask())
	})
	t.Run("false", func(t *testing.T) {
		s := Sticker{}
		assert.False(t, s.IsMask())
	})
}

func TestStickerFile(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		s := Sticker{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42}
		assert.Equal(t, s.File(), File{
			FileID:       "abc",
			FileUniqueID: "really_abc",
			FileSize:     42,
		})
	})
	t.Run("empty", func(t *testing.T) {
		var s Sticker
		assert.Equal(t, s.File(), File{})
	})
}
