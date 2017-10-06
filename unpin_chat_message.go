package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// UnpinChatMessage unpin a message in a supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
func (bot *Bot) UnpinChatMessage(chatID int64) (bool, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "unpinChatMessage", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
