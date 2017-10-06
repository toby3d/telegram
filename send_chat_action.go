package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SendChatAction tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a
// noticeable amount of time to arrive.
func (bot *Bot) SendChatAction(chatID int64, action string) (bool, error) {
	var args http.Args
	args.Add("action", action) // Type of action to broadcast
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "sendChatAction", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
