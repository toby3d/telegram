package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// ExportChatInviteLink export an invite link to a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns exported invite link as String on success.
func (bot *Bot) ExportChatInviteLink(chat interface{}) (string, error) {
	var args http.Args
	switch id := chat.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @username)
		args.Add("chat_id", id)
	default:
		return "", errors.New(errorInt64OrString)
	}

	resp, err := bot.request(nil, "exportChatInviteLink", &args)
	if err != nil {
		return "", err
	}

	var data string
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
