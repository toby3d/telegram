package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// DeleteChatPhoto delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
func (bot *Bot) DeleteChatPhoto(chatID int64) (bool, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "deleteChatPhoto", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
