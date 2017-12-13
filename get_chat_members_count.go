package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatMembersCount get the number of members in a chat. Returns Int on
// success.
func (bot *Bot) GetChatMembersCount(chatID int64) (int, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "getChatMembersCount", args)
	if err != nil {
		return 0, err
	}

	var data int
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
