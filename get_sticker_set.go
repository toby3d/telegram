package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetStickerSetParameters represents data for GetStickerSet method.
type GetStickerSetParameters struct {
	Name string `json:"name"`
}

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (set *StickerSet, err error) {
	dst, err := json.Marshal(&GetStickerSetParameters{Name: name})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetStickerSet)
	if err != nil {
		return
	}

	set = new(StickerSet)
	err = json.Unmarshal(*resp.Result, set)
	return
}
