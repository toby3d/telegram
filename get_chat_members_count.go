package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatMembersCount get the number of members in a chat. Returns Int on success.
func (bot *Bot) GetChatMembersCount(chat interface{}) (int, error) {
	var args http.Args
	switch id := chat.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return 0, errors.New(errorInt64OrString)
	}

	resp, err := bot.get("getChatMembersCount", &args)
	if err != nil {
		return 0, err
	}

	var data int
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
