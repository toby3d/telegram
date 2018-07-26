package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SendAnimationParameters represents data for SendAnimation method.
type SendAnimationParameters struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID int64 `json:"chat_id"`

	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Animation InputFile `json:"animation"`

	// Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`

	// Animation width
	Width int `json:"width,omitempty"`

	// Animation height
	Height int `json:"height,omitempty"`

	// Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb InputFile `json:"thumb,omitempty"`

	// Animation caption (may also be used when resending animation by file_id), 0-200 characters
	Caption string `json:"caption,omitempty"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// NewAnimation creates SendAnimationParameters only with required parameters.
func NewAnimation(chatID int64, animation interface{}) *SendAnimationParameters {
	return &SendAnimationParameters{
		ChatID:    chatID,
		Animation: animation,
	}
}

// SendAnimation send animation files (GIF or H.264/MPEG-4 AVC video without
// sound). On success, the sent Message is returned. Bots can currently send
// animation files of up to 50 MB in size, this limit may be changed in the
// future.
func (bot *Bot) SendAnimation(params *SendAnimationParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(params.ChatID, 10))

	if params.Caption != "" {
		args.Add("caption", params.Caption)
	}

	if params.ReplyMarkup != nil {
		dst, err := json.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		args.Add("reply_markup", string(dst))
	}

	if params.ReplyToMessageID != 0 {
		args.Add("reply_to_message_id", strconv.Itoa(params.ReplyToMessageID))
	}

	args.Add("disable_notification", strconv.FormatBool(params.DisableNotification))

	resp, err := bot.Upload(MethodSendAnimation, "animation", "", params.Animation, args)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
