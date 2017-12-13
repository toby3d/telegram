package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetStickerPositionInSet move a sticker in a set created by the bot to a
// specific position. Returns True on success.
func (bot *Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("sticker", sticker)
	args.Add("position", strconv.Itoa(position))

	resp, err := bot.request(nil, "setStickerPositionInSet", args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
