package telegram

import json "github.com/pquerna/ffjson/ffjson"

type PinChatMessageParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification"`
}

// PinChatMessage pin a message in a supergroup or a channel. The bot must be an administrator in the
// chat for this to work and must have the 'can_pin_messages' admin right in the supergroup or
// 'can_edit_messages' admin right in the channel. Returns True on success.
func (bot *Bot) PinChatMessage(params *PinChatMessageParameters) (bool, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "pinChatMessage")
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
