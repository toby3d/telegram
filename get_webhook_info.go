package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetWebhookInfo get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (bot *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	resp, err := bot.request("getWebhookInfo", nil)
	if err != nil {
		return nil, err
	}

	var data WebhookInfo
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
