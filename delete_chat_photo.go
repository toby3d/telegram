package telegram

import json "github.com/pquerna/ffjson/ffjson"

// DeleteChatPhotoParameters represents data for DeleteChatPhoto method.
type DeleteChatPhotoParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// DeleteChatPhoto delete a chat photo. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the
// 'All Members Are Admins' setting is off in the target group.
func (bot *Bot) DeleteChatPhoto(chatID int64) (bool, error) {
	dst, err := json.Marshal(&DeleteChatPhotoParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteChatPhoto)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
