package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// ExportChatInviteLink export an invite link to a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns exported invite link as String on success.
func (bot *Bot) ExportChatInviteLink(chatID int64) (string, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "exportChatInviteLink", &args)
	if err != nil {
		return "", err
	}

	var data string
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
