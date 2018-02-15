package test

import "testing"

func TestGetMe(t *testing.T) {
	var err error
	bot.Self, err = bot.GetMe()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if bot.Self == nil {
		t.Error("unexpected result: bot user is nil")
		t.FailNow()
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	photos, err := bot.GetUserProfilePhotos(chatID, 0, 100)
	if err != nil {
		t.Error(err.Error())
	}
	if photos == nil {
		t.Error("unexpected result: photos is nil")
	}
}

func TestGetFile(t *testing.T) {
	file, err := bot.GetFile(documentFileID)
	if err != nil {
		t.Error(err.Error())
	}
	if file == nil {
		t.Error("unexpected result: file is nil")
	}
}
