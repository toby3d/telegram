package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (*StickerSet, error) {
	var args http.Args
	args.Add("name", name) // Name of the sticker set

	resp, err := bot.request("getStickerSet", &args)
	if err != nil {
		return nil, err
	}

	var data StickerSet
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
