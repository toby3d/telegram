package test

import (
	"os"
	"testing"

	"github.com/toby3d/telegram"
)

const accessToken = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"

var bot = new(telegram.Bot)

func TestMain(m *testing.M) {
	bot.AccessToken = accessToken
	os.Exit(m.Run())
}
