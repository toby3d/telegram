package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetChatMember get information about a member of a chat. Returns a ChatMember object on success.
func (bot *Bot) GetChatMember(chatID int64, user int) (*ChatMember, error) {
	var args http.Args
	args.Add("user_id", strconv.Itoa(user)) // Unique identifier of the target user
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.request(nil, "getChatMember", &args)
	if err != nil {
		return nil, err
	}

	var data ChatMember
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
