package test

import (
	"testing"

	http "github.com/valyala/fasthttp"
	tg "gitlab.com/toby3d/telegram"
)

const (
	photoFileID    = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	documentFileID = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	// audioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	// voiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	// videoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	// videoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	// stickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

var (
	photoURL  *http.URI
	messageID int
)

func TestSendPhoto(t *testing.T) {
	resp, err := bot.SendPhoto(
		tg.NewPhoto(chatID, photoFileID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	} else {
		messageID = resp.ID
	}
}

func TestSendDocument(t *testing.T) {
	resp, err := bot.SendDocument(
		tg.NewDocument(chatID, documentFileID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendMediaGroup(t *testing.T) {
	resp, err := bot.SendMediaGroup(
		tg.NewMediaGroup(
			chatID,
			tg.NewInputMediaPhoto(photoFileID),
			tg.NewInputMediaPhoto(photoURL.String()),
		),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if len(resp) <= 0 {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendLocation(t *testing.T) {
	resp, err := bot.SendLocation(
		tg.NewLocation(chatID, 36.724510, 139.268181),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestSendVenue(t *testing.T) {
	resp, err := bot.SendVenue(
		tg.NewVenue(chatID, 36.724510, 139.268181, "Japan", "Japan"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestSendContact(t *testing.T) {
	resp, err := bot.SendContact(
		tg.NewContact(chatID, "+42410", "Telegram"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}
