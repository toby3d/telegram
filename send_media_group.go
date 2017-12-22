package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SendMediaGroup send a group of photos or videos as an album. On success, an array of the sent
// Messages is returned.
func (bot *Bot) SendMediaGroup(chatID int64, media []InputFile, replyToMessageID int, disableNotification bool) ([]*Message, error) {
	args := http.AcquireArgs()
	args.Add("chat_id", strconv.FormatInt(chatID, 10))
	args.Add("disable_notification", strconv.FormatBool(disableNotification))

	if replyToMessageID != 0 {
		args.Add("reply_to_message_id", strconv.Itoa(replyToMessageID))
	}

	dst, err := json.Marshal(media)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "sendMediaGroup", args)
	if err != nil {
		return nil, err
	}

	var data []*Message
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
