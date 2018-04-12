package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetMe testing your bot's auth token. Requires no parameters. Returns basic
// information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	resp, err := bot.request(nil, MethodGetMe)
	if err != nil {
		return nil, err
	}

	var data User
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
