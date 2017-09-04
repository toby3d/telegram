package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetStickerPositionInSet move a sticker in a set created by the bot to a specific position. Returns True on success.
func (bot *Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	var args http.Args
	args.Add("sticker", sticker)                 // File identifier of the sticker
	args.Add("position", strconv.Itoa(position)) // New sticker position in the set, zero-based

	resp, err := bot.post("setStickerPositionInSet", &args)
	if err != nil {
		return nil, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
