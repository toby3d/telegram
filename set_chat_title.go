package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetChatTitle change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
func (bot *Bot) SetChatTitle(chat interface{}, title string) (bool, error) {
	var args http.Args
	args.Add("title", title) // New chat title, 1-255 characters

	switch id := chatID.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.post("setChatTitle", &args)
	if err != nil {
		return nil, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
