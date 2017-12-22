package telegram

/*
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type SendStickerParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Sticker to send. Pass a file_id as String to send a file that exists on
    // the Telegram servers (recommended), pass an HTTP URL as a String for
    // Telegram to get a .webp file from the Internet, or upload a new one using
    // multipart/form-data.
	Sticker *InputFile `json:"sticker"`

	// Sends the message silently. Users will receive a notification with no
    // sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline
    // keyboard, custom reply keyboard, instructions to remove reply keyboard or
    // to force a reply from the user.
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

func NewSticker(chatID int64, sticker interface{}) *SendStickerParameters {
	params := &SendStickerParameters{ChatID: chatID}

	if sticker != nil {
		var input InputFile = file
		params.Sticker = &input
	}

	return params
}

// SendSticker send .webp stickers. On success, the sent Message is returned.
func (bot *Bot) SendSticker(params *SendStickerParameters) (*Message, error) {
	args := http.AcquireArgs()
defer http.ReleaseArgs(args)
	args.Add("chat_id", params.ChatID)
	args.Add("disable_notification", strconv.FormatBool(params.DisableNotification))

	if params.ReplyToMessageID > 0 {
		args.Add("reply_to_message_id", strconv.Itoa(params.ReplyToMessageID))
	}

	var buffer bytes.Buffer
	multi := multipart.NewWriter(&buffer)

	sticker := *params.Sticker
	switch file := sticker.(type) {
	case string:
		f, err := os.Open(file)
		if err != nil {
			return false, err
		}
		defer f.Close()

		formFile, err := multi.CreateFormFile("sticker", f.Name())
		if err != nil {
			return false, err
		}
		if _, err = io.Copy(formFile, f); err != nil {
			return false, err
		}
		multi.Close()
    case []byte:
        file
        bytes.NewReader(file).
	default:
		return false, errors.New("use string only (for current version of go-telegram)")
	}

	resp, err := bot.upload(buffer.Bytes(), multi.Boundary(), "setWebhook", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
*/
