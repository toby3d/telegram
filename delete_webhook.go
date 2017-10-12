package telegram

import json "github.com/pquerna/ffjson/ffjson"

// DeleteWebhook remove webhook integration if you decide to switch back to
// getUpdates. Returns True on success. Requires no parameters.
func (bot *Bot) DeleteWebhook() (bool, error) {
	resp, err := bot.request(nil, "deleteWebhook", nil)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
