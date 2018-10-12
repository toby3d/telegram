package telegram

import json "github.com/pquerna/ffjson/ffjson"

// PinChatMessageParameters represents data for PinChatMessage method.
type PinChatMessageParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Identifier of a message to pin
	MessageID int `json:"message_id"`

	// Pass true, if it is not necessary to send a notification to all chat
	// members about the new pinned message. Notifications are always
	// disabled in channels.
	DisableNotification bool `json:"disable_notification"`
}

// PinChatMessage pin a message in a supergroup or a channel. The bot must be an administrator in the
// chat for this to work and must have the 'can_pin_messages' admin right in the supergroup or
// 'can_edit_messages' admin right in the channel. Returns True on success.
func (bot *Bot) PinChatMessage(params *PinChatMessageParameters) (ok bool, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodPinChatMessage)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
