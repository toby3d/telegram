package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChat get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chatID int64) (*Chat, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "getChat", &args)
	if err != nil {
		return nil, err
	}

	var data Chat
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
