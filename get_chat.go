package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChat get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chat interface{}) (*Chat, error) {
	var args http.Args
	switch id := chat.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.get("getChat", &args)
	if err != nil {
		return nil, err
	}

	var data Chat
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
