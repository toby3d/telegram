package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SetChatDescriptionParameters represents data for SetChatDescription method.
type SetChatDescriptionParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// New chat description, 0-255 characters
	Description string `json:"description"`
}

// SetChatDescription change the description of a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatDescription(chatID int64, description string) (ok bool, err error) {
	dst, err := json.Marshal(&SetChatDescriptionParameters{
		ChatID:      chatID,
		Description: description,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSetChatDescription)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
