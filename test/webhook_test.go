package test

import (
	"testing"

	"gitlab.com/toby3d/telegram"
)

func TestSetWebhook(t *testing.T) {
	ok, err := bot.SetWebhook(
		telegram.NewWebhook("https://toby3d.ru/telegram", nil),
	)
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
