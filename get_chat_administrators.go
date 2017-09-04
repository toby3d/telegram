package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatAdministrators get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chat interface{}) (*[]ChatMember, error) {
	var args http.Args
	switch id := chat.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.get("getChatAdministrators", &args)
	if err != nil {
		return nil, err
	}

	var data []ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
