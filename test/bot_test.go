package test

import (
	"os"
	"testing"

	"gitlab.com/toby3d/telegram"
)

var bot = new(telegram.Bot)

func TestMain(m *testing.M) {
	bot.AccessToken = os.Getenv("BOT_ACCESS_TOKEN")
	os.Exit(m.Run())
}
