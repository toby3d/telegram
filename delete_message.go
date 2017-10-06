package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// DeleteMessage delete a message, including service messages, with the following
// limitations: A message can only be deleted if it was sent less than 48 hours
// ago; Bots can delete outgoing messages in groups and supergroups; Bots granted
// can_post_messages permissions can delete outgoing messages in channels; If the
// bot is an administrator of a group, it can delete any message there; If the
// bot has can_delete_messages permission in a supergroup or a channel, it can
// delete any message there. Returns True on success.
func (bot *Bot) DeleteMessage(chatID int64, message int) (bool, error) {
	var args http.Args
	args.Add("message_id", strconv.Itoa(message))
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "deleteMessage", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
