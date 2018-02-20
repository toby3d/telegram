package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetStickerSetParameters struct {
	Name string `json:"name"`
}

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (*StickerSet, error) {
	dst, err := json.Marshal(&GetStickerSetParameters{Name: name})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "getStickerSet")
	if err != nil {
		return nil, err
	}

	var data StickerSet
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
