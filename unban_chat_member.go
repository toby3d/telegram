package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// UnbanChatMember unban a previously kicked user in a supergroup or channel. The
// user will not return to the group or channel automatically, but will be able
// to join via link, etc. The bot must be an administrator for this to work.
// Returns True on success.
func (bot *Bot) UnbanChatMember(chatID int64, user int) (bool, error) {
	var args http.Args
	args.Add("user_id", strconv.Itoa(user))
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "unbanChatMember", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
