package telegram

import json "github.com/pquerna/ffjson/ffjson"

// DeleteMessageParameters represents data for DeleteMessage method.
type DeleteMessageParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	MessageID int `json:"message_id"`
}

// DeleteMessage delete a message, including service messages, with the following
// limitations: A message can only be deleted if it was sent less than 48 hours
// ago; Bots can delete outgoing messages in groups and supergroups; Bots granted
// can_post_messages permissions can delete outgoing messages in channels; If the
// bot is an administrator of a group, it can delete any message there; If the
// bot has can_delete_messages permission in a supergroup or a channel, it can
// delete any message there. Returns True on success.
func (bot *Bot) DeleteMessage(chatID int64, messageID int) (ok bool, err error) {
	dst, err := json.Marshal(&DeleteMessageParameters{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodDeleteMessage)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
