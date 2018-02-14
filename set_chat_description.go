package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetChatDescription change the description of a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatDescription(chatID int64, description string) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("description", description)
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "setChatDescription", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
