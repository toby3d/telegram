package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// LeaveChat leave a group, supergroup or channel. Returns True on success.
func (bot *Bot) LeaveChat(chat int64) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(chat, 10))

	resp, err := bot.request(nil, "leaveChat", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
