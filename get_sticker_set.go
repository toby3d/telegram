package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (*StickerSet, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("name", name)

	resp, err := bot.request(nil, "getStickerSet", args)
	if err != nil {
		return nil, err
	}

	var data StickerSet
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
