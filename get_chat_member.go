package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatMember get information about a member of a chat. Returns a ChatMember object on success.
func (bot *Bot) GetChatMember(chat interface{}, user int) (*ChatMember, error) {
	var args http.Args
	args.Add("user_id", strconv.Itoa(user)) // Unique identifier of the target user

	switch id := chatID.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.get("getChatMember", &args)
	if err != nil {
		return nil, err
	}

	var data ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
