package test

import (
	"net/url"
	"testing"

	"gitlab.com/toby3d/telegram"
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
	photoURL = url.URL{
		Scheme: "https",
		Host:   "simg3.gelbooru.com",
		Path:   "/images/46/24/46246c1b8c4fcc37050085a850c165c4.jpg",
	}

	messageID int
)

func TestSendPhoto(t *testing.T) {
	resp, err := bot.SendPhoto(
		telegram.NewPhoto(chatID, photoFileID),
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
		telegram.NewDocument(chatID, documentFileID),
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
		telegram.NewMediaGroup(
			chatID,
			telegram.NewInputMediaPhoto(photoFileID),
			telegram.NewInputMediaPhoto(photoURL.String()),
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
		telegram.NewLocation(chatID, 36.724510, 139.268181),
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
		telegram.NewVenue(chatID, 36.724510, 139.268181, "Japan", "Japan"),
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
		telegram.NewContact(chatID, "+42410", "Telegram"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}
