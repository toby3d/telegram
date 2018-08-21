package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendChatActionParameters represents data for SendChatAction method.
type SendChatActionParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	Action string `json:"action"`
}

// SendChatAction tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a
// noticeable amount of time to arrive.
func (bot *Bot) SendChatAction(chatID int64, action string) (ok bool, err error) {
	dst, err := json.Marshal(&SendChatActionParameters{
		ChatID: chatID,
		Action: action,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendChatAction)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
