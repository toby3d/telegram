package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetChatAdministratorsParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChatAdministrators get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID int64) ([]ChatMember, error) {
	dst, err := json.Marshal(&GetChatAdministratorsParameters{ChatID: chatID})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "getChatAdministrators")
	if err != nil {
		return nil, err
	}

	var data []ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
