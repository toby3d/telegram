package telegram

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	http "github.com/valyala/fasthttp"
)

const (
	testPhrase = "Hello, world!"
	testLink   = "https://toby3d.me/"
	testID     = 123456789
	// testUsername = "toby3d"
	// testChannel  = "toby3dRu"
	testCode = "<spoiler>hide me</spoiler>"
)

var (
	testPhotoSize = PhotoSize{
		FileID:   "cba",
		FileSize: 123,
		Width:    128,
		Height:   128,
	}
	testAnimation = Animation{
		FileID:   "abc",
		FileName: "animation",
		FileSize: 321,
		Thumb:    &testPhotoSize,
	}
	testAudio = Audio{
		Performer: "One-Aloner",
		Title:     "(Not) your reality (Doki Doki Literature Club cover)",
		Thumb:     &testPhotoSize,
	}
)

func TestMarkdown(t *testing.T) {
	t.Run("bold", func(t *testing.T) {
		assert.Equal(t, "*"+testPhrase+"*", NewMarkdownBold(testPhrase))
	})
	t.Run("italic", func(t *testing.T) {
		assert.Equal(t, "_"+testPhrase+"_", NewMarkdownItalic(testPhrase))
	})
	t.Run("link", func(t *testing.T) {
		link := http.AcquireURI()
		defer http.ReleaseURI(link)
		link.Update(testLink)
		assert.Equal(t, "["+testPhrase+"]("+testLink+")", NewMarkdownURL(testPhrase, link))
	})
	t.Run("mention", func(t *testing.T) {
		assert.Equal(t, "["+testPhrase+"](tg://user/?id="+strconv.Itoa(testID)+")", NewMarkdownMention(testPhrase, testID))
	})
	t.Run("code", func(t *testing.T) {
		assert.Equal(t, "`"+testCode+"`", NewMarkdownCode(testCode))
	})
	t.Run("code block", func(t *testing.T) {
		assert.Equal(t, "```"+testCode+"```", NewMarkdownCodeBlock(testCode))
	})
}

func TestHTML(t *testing.T) {
	t.Run("bold", func(t *testing.T) {
		assert.Equal(t, "<b>"+testPhrase+"</b>", NewHTMLBold(testPhrase))
	})
	t.Run("italic", func(t *testing.T) {
		assert.Equal(t, "<i>"+testPhrase+"</i>", NewHTMLItalic(testPhrase))
	})
	t.Run("link", func(t *testing.T) {
		link := http.AcquireURI()
		defer http.ReleaseURI(link)
		link.Update(testLink)
		assert.Equal(t, `<a href="`+link.String()+`">`+testPhrase+`</a>`, NewHTMLURL(testPhrase, link))
	})
	t.Run("mention", func(t *testing.T) {
		assert.Equal(t, `<a href="tg://user/?id=`+strconv.Itoa(testID)+`">`+testPhrase+`</a>`, NewHTMLMention(testPhrase, testID))
	})
	t.Run("code", func(t *testing.T) {
		assert.Equal(t, "<code>"+testCode+"</code>", NewHTMLCode(testCode))
	})
	t.Run("code block", func(t *testing.T) {
		assert.Equal(t, "<pre>"+testCode+"</pre>", NewHTMLCodeBlock(testCode))
	})
}

func TestAnimation(t *testing.T) {
	t.Run("has thumb", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			a := new(Animation)
			assert.False(t, a.HasThumb())
		})
		t.Run("true", func(t *testing.T) {
			assert.True(t, testAnimation.HasThumb())
		})
	})
	t.Run("file", func(t *testing.T) {
		assert.NotNil(t, testAnimation.File())
	})
}

func TestAudio(t *testing.T) {
	t.Run("has performer", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			a := new(Audio)
			assert.False(t, a.HasPerformer())
		})
		t.Run("true", func(t *testing.T) {
			assert.True(t, testAudio.HasPerformer())
		})
	})
	t.Run("has title", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			a := new(Audio)
			assert.False(t, a.HasTitle())
		})
		t.Run("true", func(t *testing.T) {
			assert.True(t, testAudio.HasTitle())
		})
	})
	t.Run("has thumb", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			a := new(Audio)
			assert.False(t, a.HasThumb())
		})
		t.Run("true", func(t *testing.T) {
			assert.True(t, testAudio.HasThumb())
		})
	})
	t.Run("file", func(t *testing.T) {
		assert.NotNil(t, testAudio.File())
	})
	t.Run("full name", func(t *testing.T) {
		for _, tc := range []struct {
			message   string
			audio     *Audio
			separator string
			expResult string
		}{{
			message:   "empty",
			expResult: DefaultAudioTitle,
		}, {
			message:   "separator only",
			separator: DefaultAudioSeparator,
			expResult: DefaultAudioTitle,
		}, {
			message: "title only",
			audio: &Audio{
				Title: testAudio.Title,
			},
			expResult: testAudio.Title,
		}, {
			message: "performer only",
			audio: &Audio{
				Performer: testAudio.Performer,
			},
			expResult: testAudio.Performer + DefaultAudioSeparator + DefaultAudioTitle,
		}, {
			message: "title & performer",
			audio: &Audio{
				Performer: testAudio.Performer,
				Title:     testAudio.Title,
			},
			expResult: testAudio.Performer + DefaultAudioSeparator + testAudio.Title,
		}, {
			message:   "title & separator",
			separator: " | ",
			audio: &Audio{
				Title: testAudio.Title,
			},
			expResult: testAudio.Title,
		}, {
			message:   "performer & separator",
			separator: " | ",
			audio: &Audio{
				Performer: testAudio.Performer,
			},
			expResult: testAudio.Performer + " | " + DefaultAudioTitle,
		}, {
			message:   "performer, title & separator",
			separator: " | ",
			audio: &Audio{
				Performer: testAudio.Performer,
				Title:     testAudio.Title,
			},
			expResult: testAudio.Performer + " | " + testAudio.Title,
		}} {
			tc := tc
			t.Run(tc.message, func(t *testing.T) {
				assert.Equal(t, tc.expResult, tc.audio.FullName(tc.separator))
			})
		}
	})
}
