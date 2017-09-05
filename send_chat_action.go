package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const (
	ActionTyping          = "typing"
	ActionUploadPhoto     = "upload_photo"
	ActionRecordVideo     = "record_video"
	ActionUploadVideo     = "upload_video"
	ActionRecordAudio     = "record_audio"
	ActionUploadAudio     = "upload_audio"
	ActionUploadDocument  = "upload_document"
	ActionFindLocation    = "find_location"
	ActionRecordVideoNote = "record_video_note"
	ActionUploadVideoNote = "upload_video_note"
)

// SendChatAction tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (bot *Bot) SendChatAction(chat interface{}, action string) (bool, error) {
	var args http.Args
	args.Add("action", action) // Type of action to broadcast

	switch id := chat.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @channelusername)
		args.Add("chat_id", id)
	default:
		return false, errors.New(errorInt64OrString)
	}

	resp, err := bot.request("sendChatAction", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
