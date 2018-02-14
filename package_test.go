package telegram

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	log "github.com/kirillDanshin/dlog"
)

const (
	accessToken = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"

	chatID           = 76918703
	superGroupID     = -1001120141283
	replyToMessageID = 35

	photoFileID     = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	documentFileID  = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	audioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	voiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	videoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	videoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	stickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

var (
	bot      = new(Bot)
	photoURL = url.URL{
		Scheme: "https",
		Host:   "simg3.gelbooru.com",
		Path:   "/images/46/24/46246c1b8c4fcc37050085a850c165c4.jpg",
	}

	messageID int
)

func TestMain(m *testing.M) {
	bot.AccessToken = accessToken
	os.Exit(m.Run())
}

func TestGetUpdates(t *testing.T) {
	updates, err := bot.GetUpdates(nil)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(updates)
	if len(updates) <= 0 {
		t.Error("unexpected result: no updates")
	}
}

func TestSetWebhook(t *testing.T) {
	ok, err := bot.SetWebhook(NewWebhook("https://toby3d.github.io/telegram", nil))
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestDeleteWebhook(t *testing.T) {
	ok, err := bot.DeleteWebhook()
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestGetWebhookInfo(t *testing.T) {
	info, err := bot.GetWebhookInfo()
	if err != nil {
		t.Error(err.Error())
	}
	if info == nil {
		t.Error("unexpected result: info is nil")
	}
}

func TestGetMe(t *testing.T) {
	var err error
	bot.Self, err = bot.GetMe()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	log.D(bot.Self)
	if bot.Self == nil {
		t.Error("unexpected result: bot user is nil")
		t.FailNow()
	}
}

func TestSendMessage(t *testing.T) {
	resp, err := bot.SendMessage(
		NewMessage(chatID, "Hello, World"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestForwardMessage(t *testing.T) {
	resp, err := bot.ForwardMessage(
		NewForwardMessage(chatID, superGroupID, replyToMessageID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendPhoto(t *testing.T) {
	resp, err := bot.SendPhoto(
		NewPhoto(chatID, photoFileID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	} else {
		messageID = resp.ID
	}
}

func TestSendDocument(t *testing.T) {
	resp, err := bot.SendDocument(
		NewDocument(chatID, documentFileID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendMediaGroup(t *testing.T) {
	resp, err := bot.SendMediaGroup(NewMediaGroup(
		chatID,
		NewInputMediaPhoto(photoFileID),
		NewInputMediaPhoto(photoURL.String()),
	))
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if len(resp) <= 0 {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendLocation(t *testing.T) {
	resp, err := bot.SendLocation(NewLocation(chatID, 36.724510, 139.268181))
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendContact(t *testing.T) {
	resp, err := bot.SendContact(NewContact(chatID, "+42410", "Telegram"))
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestSendChatAction(t *testing.T) {
	ok, err := bot.SendChatAction(chatID, ActionTyping)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	photos, err := bot.GetUserProfilePhotos(chatID, 0, 100)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(photos)
	if photos == nil {
		t.Error("unexpected result: photos is nil")
	}
}

func TestGetFile(t *testing.T) {
	file, err := bot.GetFile(documentFileID)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(file)
	if file == nil {
		t.Error("unexpected result: file is nil")
	}
}

func TestKickChatMember(t *testing.T) {
	ok, err := bot.KickChatMember(superGroupID, chatID, time.Now().Add(time.Second*30).Unix())
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestUnbanChatMember(t *testing.T) {
	ok, err := bot.UnbanChatMember(superGroupID, chatID)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestExportChatInviteLink(t *testing.T) {
	inviteLink, err := bot.ExportChatInviteLink(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	log.Ln("InviteLink:", inviteLink)
	if inviteLink == "" {
		t.Error("unexpected result: inviteLink is empty")
	}
}

func TestSetChatPhoto(t *testing.T) {
	ok, err := bot.SetChatPhoto(superGroupID, "./test/photo.png")
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestDeleteChatPhoto(t *testing.T) {
	ok, err := bot.DeleteChatPhoto(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestSetChatTitle(t *testing.T) {
	ok, err := bot.SetChatTitle(superGroupID, "Go Telegram Bot API")
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestSetChatDescription(t *testing.T) {
	ok, err := bot.SetChatDescription(
		superGroupID,
		fmt.Sprint("Go Telegram BotAPI testing chat (updated: ", time.Now().String(), ")"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestPinChatMessage(t *testing.T) {
	ok, err := bot.PinChatMessage(superGroupID, replyToMessageID, true)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestUnpinChatMessage(t *testing.T) {
	ok, err := bot.UnpinChatMessage(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestLeaveChat(t *testing.T) {
	ErrNotMember := "Forbidden: bot is not a member of the supergroup chat"
	ok, err := bot.LeaveChat(-1001037355946)
	if err != nil && err.Error() != ErrNotMember {
		t.Error(err.Error())
	}
	if !ok && err.Error() != ErrNotMember {
		t.Error("unexpected result: ok is not true")
	}
}

func TestGetChat(t *testing.T) {
	chat, err := bot.GetChat(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(chat)
	if chat == nil {
		t.Error("unexpected result: chat is nil")
	}
}

func TestGetChatAdministrators(t *testing.T) {
	admins, err := bot.GetChatAdministrators(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(admins)
	if len(admins) <= 0 {
		t.Error("unexpected result: admins not exist")
	}
}

func TestGetChatMembersCount(t *testing.T) {
	total, err := bot.GetChatMembersCount(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	log.Ln(total, "members")
	if total <= 0 {
		t.Error("unexpected result: members count is 0")
	}
}

func TestGetChatMember(t *testing.T) {
	member, err := bot.GetChatMember(superGroupID, chatID)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(member)
	if member == nil {
		t.Error("unexpected result: member is nil")
	}
}

func TestSetChatStickerSet(t *testing.T) {
	ErrCantSetStickerSet := "Bad Request: can't set supergroup sticker set"
	ok, err := bot.SetChatStickerSet(superGroupID, "HentaiDB")
	if err != nil &&
		err.Error() != ErrCantSetStickerSet {
		t.Error(err.Error())
	}
	if !ok &&
		err.Error() != ErrCantSetStickerSet {
		t.Error("unexpected result: ok is not true")
	}
}

func TestDeleteChatStickerSet(t *testing.T) {
	ErrCantSetStickerSet := "Bad Request: can't set supergroup sticker set"
	ok, err := bot.DeleteChatStickerSet(superGroupID)
	if err != nil &&
		err.Error() != ErrCantSetStickerSet {
		t.Error(err.Error())
	}
	if !ok &&
		err.Error() != ErrCantSetStickerSet {
		t.Error("unexpected result: ok is not true")
	}
}

func TestEditMessageText(t *testing.T) {
	text := NewMessageText(
		fmt.Sprint("Go Telegram BotAPI testing chat (updated: ", time.Now().String(), ")"),
	)
	text.ChatID = chatID
	text.MessageID = replyToMessageID
	resp, err := bot.EditMessageText(text)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestEditMessageCaption(t *testing.T) {
	var caption EditMessageCaptionParameters
	caption.Caption = fmt.Sprint("Go Telegram BotAPI testing chat (updated: ", time.Now().String(), ")")
	caption.ChatID = chatID
	caption.MessageID = messageID
	resp, err := bot.EditMessageCaption(&caption)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestEditMessageReplyMarkup(t *testing.T) {
	var markup EditMessageReplyMarkupParameters
	markup.ChatID = superGroupID
	markup.MessageID = replyToMessageID
	markup.ReplyMarkup = NewInlineKeyboardMarkup(
		NewInlineKeyboardRow(
			NewInlineKeyboardButton("hello", fmt.Sprint("time", time.Now().Unix())),
		),
	)
	resp, err := bot.EditMessageReplyMarkup(&markup)
	if err != nil {
		t.Error(err.Error())
	}
	log.D(resp)
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestDeleteMessage(t *testing.T) {
	ok, err := bot.DeleteMessage(chatID, messageID)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}
