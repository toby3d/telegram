package telegram

import json "github.com/pquerna/ffjson/ffjson"

type SetChatDescriptionParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	Description string `json:"description"`
}

// SetChatDescription change the description of a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatDescription(chatID int64, description string) (bool, error) {
	dst, err := json.Marshal(&SetChatDescriptionParameters{
		ChatID:      chatID,
		Description: description,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "setChatDescription")
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
