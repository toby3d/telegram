package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// UnbanChatMember unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
func (bot *Bot) UnbanChatMember(chat interface{}, user int) (bool, error) {
	var args http.Args
	args.Add("user_id", strconv.Itoa(user)) // Unique identifier of the target user

	switch id := chatID.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @username)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.post("unbanChatMember", &args)
	if err != nil {
		return nil, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
