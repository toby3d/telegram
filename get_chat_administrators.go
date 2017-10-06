package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatAdministrators get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID int64) (*[]ChatMember, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "getChatAdministrators", &args)
	if err != nil {
		return nil, err
	}

	var data []ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
