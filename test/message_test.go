package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/toby3d/telegram"
)

const replyToMessageID = 35

func TestSendMessage(t *testing.T) {
	resp, err := bot.SendMessage(
		telegram.NewMessage(chatID, "Hello, World"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestForwardMessage(t *testing.T) {
	resp, err := bot.ForwardMessage(
		telegram.NewForwardMessage(chatID, superGroupID, replyToMessageID),
	)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: message is nil")
	}
}

func TestEditMessageText(t *testing.T) {
	text := telegram.NewMessageText(
		fmt.Sprint("Go Telegram BotAPI testing chat (", time.Now().Unix(), ")"),
	)
	text.ChatID = chatID
	text.MessageID = replyToMessageID
	resp, err := bot.EditMessageText(text)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestEditMessageCaption(t *testing.T) {
	var caption telegram.EditMessageCaptionParameters
	caption.Caption = fmt.Sprint("Go Telegram BotAPI testing chat (", time.Now().Unix(), ")")
	caption.ChatID = chatID
	caption.MessageID = messageID
	resp, err := bot.EditMessageCaption(&caption)
	if err != nil {
		t.Error(err.Error())
	}
	if resp == nil {
		t.Error("unexpected result: resp is nil")
	}
}

func TestEditMessageReplyMarkup(t *testing.T) {
	var markup telegram.EditMessageReplyMarkupParameters
	markup.ChatID = superGroupID
	markup.MessageID = replyToMessageID
	markup.ReplyMarkup = telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButton(
				"hello",
				fmt.Sprint(time.Now().Unix()),
			),
		),
	)
	resp, err := bot.EditMessageReplyMarkup(&markup)
	if err != nil {
		t.Error(err.Error())
	}
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
