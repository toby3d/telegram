package telegram

import json "github.com/pquerna/ffjson/ffjson"

// DeleteWebhook remove webhook integration if you decide to switch back to
// getUpdates. Returns True on success. Requires no parameters.
func (bot *Bot) DeleteWebhook() (ok bool, err error) {
	resp, err := bot.request(nil, MethodDeleteWebhook)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
