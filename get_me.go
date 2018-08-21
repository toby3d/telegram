package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetMe testing your bot's auth token. Requires no parameters. Returns basic
// information about the bot in form of a User object.
func (bot *Bot) GetMe() (me *User, err error) {
	resp, err := bot.request(nil, MethodGetMe)
	if err != nil {
		return
	}

	me = new(User)
	err = json.Unmarshal(*resp.Result, me)
	return
}
