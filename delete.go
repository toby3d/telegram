package telegram

type (
	// DeleteChatPhotoParameters represents data for DeleteChatPhoto method.
	DeleteChatPhotoParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// DeleteChatStickerSetParameters represents data for DeleteChatStickerSet method.
	DeleteChatStickerSetParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// DeleteMessageParameters represents data for DeleteMessage method.
	DeleteMessageParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Identifier of the message to delete
		MessageID int `json:"message_id"`
	}

	// DeleteStickerFromSetParameters represents data for DeleteStickerFromSet method.
	DeleteStickerFromSetParameters struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`
	}
)

// DeleteChatPhoto delete a chat photo. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the
// 'All Members Are Admins' setting is off in the target group.
func (bot *Bot) DeleteChatPhoto(chatID int64) (bool, error) {
	dst, err := parser.Marshal(&DeleteChatPhotoParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteChatPhoto)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// DeleteChatStickerSet delete a group sticker set from a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success.
func (bot *Bot) DeleteChatStickerSet(chatID int64) (bool, error) {
	dst, err := parser.Marshal(&DeleteChatStickerSetParameters{ChatID: chatID})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteChatStickerSet)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// DeleteWebhook remove webhook integration if you decide to switch back to
// getUpdates. Returns True on success. Requires no parameters.
func (bot *Bot) DeleteWebhook() (bool, error) {
	resp, err := bot.request(nil, MethodDeleteWebhook)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// DeleteMessage delete a message, including service messages, with the following
// limitations: A message can only be deleted if it was sent less than 48 hours
// ago; Bots can delete outgoing messages in groups and supergroups; Bots granted
// can_post_messages permissions can delete outgoing messages in channels; If the
// bot is an administrator of a group, it can delete any message there; If the
// bot has can_delete_messages permission in a supergroup or a channel, it can
// delete any message there. Returns True on success.
func (bot *Bot) DeleteMessage(chatID int64, messageID int) (bool, error) {
	dst, err := parser.Marshal(&DeleteMessageParameters{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteMessage)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// DeleteStickerFromSet delete a sticker from a set created by the bot. Returns
// True on success.
func (bot *Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	dst, err := parser.Marshal(&DeleteStickerFromSetParameters{Sticker: sticker})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteStickerFromSet)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
