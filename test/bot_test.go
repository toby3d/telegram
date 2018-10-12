package test

import (
	"os"
	"path"
	"testing"

	http "github.com/valyala/fasthttp"
	tg "gitlab.com/toby3d/telegram"
)

var bot = new(tg.Bot)

func TestMain(m *testing.M) {
	photoURL = http.AcquireURI()
	defer http.ReleaseURI(photoURL)

	photoURL.SetScheme("https")
	photoURL.SetHost("simg3.gelbooru.com")
	photoURL.SetPath(path.Join("images", "46", "24", "46246c1b8c4fcc37050085a850c165c4.jpg"))

	bot.AccessToken = os.Getenv("BOT_ACCESS_TOKEN")
	os.Exit(m.Run())
}
