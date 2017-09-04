package telegram

// DeleteWebhook remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
func (bot *Bot) DeleteWebhook() (*Response, error) {
	return bot.get("deleteWebhook", nil)
}
