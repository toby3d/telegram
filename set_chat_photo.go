package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type SetChatPhotoParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	ChatPhoto interface{} `json:"chat_photo"`
}

// SetChatPhoto set a new profile photo for the chat. Photos can't be changed for private chats. The
// bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the 'All Members Are
// Admins' setting is off in the target group.
func (bot *Bot) SetChatPhoto(chatID int64, chatPhoto interface{}) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.upload(chatPhoto, TypePhoto, "chat_photo", "setChatPhoto", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
