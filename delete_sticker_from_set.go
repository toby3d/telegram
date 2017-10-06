package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// DeleteStickerFromSet delete a sticker from a set created by the bot. Returns
// True on success.
func (bot *Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	var args http.Args
	args.Add("sticker", sticker)

	resp, err := bot.request(nil, "deleteStickerFromSet", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
