package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SendDocumentParameters represents data for SendDocument method.
type SendDocumentParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or
	// upload a new one using multipart/form-data.
	Document InputFile `json:"document"`

	// Document caption (may also be used when resending documents by file_id), 0-200 characters
	Caption string `json:"caption,omitempty"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
	// fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// NewDocument creates SendDocumentParameters only with required parameters.
func NewDocument(chatID int64, document interface{}) *SendDocumentParameters {
	return &SendDocumentParameters{
		ChatID:   chatID,
		Document: document,
	}
}

// SendDocument send general files. On success, the sent Message is returned. Bots can currently send
// files of any type of up to 50 MB in size, this limit may be changed in the future.
func (bot *Bot) SendDocument(params *SendDocumentParameters) (*Message, error) {
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

	resp, err := bot.Upload(MethodSendDocument, "document", "", params.Document, args)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
