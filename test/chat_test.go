package test

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/toby3d/telegram"
)

const (
	chatID       = 76918703
	superGroupID = -1001120141283
)

func TestSendChatAction(t *testing.T) {
	ok, err := bot.SendChatAction(chatID, telegram.ActionTyping)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

/*
func TestKickChatMember(t *testing.T) {
	ok, err := bot.KickChatMember(&telegram.KickChatMemberParameters{
		ChatID:    superGroupID,
		UserID:    chatID,
		UntilDate: time.Now().Add(time.Second * 30).Unix(),
	})
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
*/

func TestExportChatInviteLink(t *testing.T) {
	inviteLink, err := bot.ExportChatInviteLink(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("InviteLink:", inviteLink)
	if inviteLink == "" {
		t.Error("unexpected result: inviteLink is empty")
	}
}

func TestSetChatPhoto(t *testing.T) {
	ok, err := bot.SetChatPhoto(superGroupID, "./photo.png")
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
		fmt.Sprint("Go Telegram BotAPI testing chat (", time.Now().Unix(), ")"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("unexpected result: ok is not true")
	}
}

func TestPinChatMessage(t *testing.T) {
	ok, err := bot.PinChatMessage(&telegram.PinChatMessageParameters{
		ChatID:              superGroupID,
		MessageID:           replyToMessageID,
		DisableNotification: true,
	})
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

/*
func TestLeaveChat(t *testing.T) {
	ErrNotMember := "Forbidden: bot is not a member of the supergroup chat"
	ok, err := bot.LeaveChat(-1001037355946)
	if err != nil &&
		err.Error() != ErrNotMember {
		t.Error(err.Error())
	}
	if !ok &&
		err.Error() != ErrNotMember {
		t.Error("unexpected result: ok is not true")
	}
}
*/

func TestGetChat(t *testing.T) {
	chat, err := bot.GetChat(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	if chat == nil {
		t.Error("unexpected result: chat is nil")
	}
}

func TestGetChatAdministrators(t *testing.T) {
	admins, err := bot.GetChatAdministrators(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	if len(admins) <= 0 {
		t.Error("unexpected result: admins not exist")
	}
}

func TestGetChatMembersCount(t *testing.T) {
	total, err := bot.GetChatMembersCount(superGroupID)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(total, "members")
	if total <= 0 {
		t.Error("unexpected result: members count is 0")
	}
}

func TestGetChatMember(t *testing.T) {
	member, err := bot.GetChatMember(superGroupID, chatID)
	if err != nil {
		t.Error(err.Error())
	}
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
