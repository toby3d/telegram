package test

import "testing"

func TestGetUpdates(t *testing.T) {
	updates, err := bot.GetUpdates(nil)
	if err != nil {
		t.Error(err.Error())
	}
	if len(updates) <= 0 {
		t.Error("unexpected result: no updates")
	}
}
