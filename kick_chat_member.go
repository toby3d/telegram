package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// KickChatMember kick a user from a group, a supergroup or a channel. In the case of supergroups and
// channels, the user will not be able to return to the group on their own using invite links, etc.,
// unless unbanned first. The bot must be an administrator in the chat for this to work and must have
// the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the 'All Members Are
// Admins' setting is off in the target group. Otherwise members may only be removed by the group's
// creator or by the member that added them.
func (bot *Bot) KickChatMember(chatID int64, userID int, untilDate int64) (bool, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))
	args.Add("user_id", strconv.Itoa(userID))
	args.Add("until_date", strconv.FormatInt(untilDate, 10))

	resp, err := bot.request(nil, "kickChatMember", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
