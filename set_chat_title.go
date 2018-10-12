package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SetChatTitleParameters represents data for SetChatTitle method.
type SetChatTitleParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// New chat title, 1-255 characters
	Title string `json:"title"`
}

// SetChatTitle change the title of a chat. Titles can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the
// 'All Members Are Admins' setting is off in the target group.
func (bot *Bot) SetChatTitle(chatID int64, title string) (ok bool, err error) {
	dst, err := json.Marshal(&SetChatTitleParameters{
		ChatID: chatID,
		Title:  title,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSetChatTitle)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
