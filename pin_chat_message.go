package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// PinChatMessage pin a message in a supergroup or a channel. The bot must be an administrator in the
// chat for this to work and must have the 'can_pin_messages' admin right in the supergroup or
// 'can_edit_messages' admin right in the channel. Returns True on success.
func (bot *Bot) PinChatMessage(chatID int64, messageID int, disableNotification bool) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	// Unique identifier for the target chat or username of the target channel
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	// Identifier of a message to pin
	args.Add("message_id", strconv.Itoa(messageID))

	// Pass True, if it is not necessary to send a notification to all chat members about the new
	// pinned message. Notifications are always disabled in channels.
	args.Add("disable_notification", strconv.FormatBool(disableNotification))

	resp, err := bot.request(nil, "pinChatMessage", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
