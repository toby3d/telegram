package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetWebhookInfo get current webhook status. Requires no parameters. On success,
// returns a WebhookInfo object. If the bot is using getUpdates, will return an
// object with the url field empty.
func (bot *Bot) GetWebhookInfo() (info *WebhookInfo, err error) {
	resp, err := bot.request(nil, MethodGetWebhookInfo)
	if err != nil {
		return
	}

	info = new(WebhookInfo)
	err = json.Unmarshal(*resp.Result, info)
	return
}
