package telegram

import json "github.com/pquerna/ffjson/ffjson"

// DeleteStickerFromSetParameters represents data for DeleteStickerFromSet method.
type DeleteStickerFromSetParameters struct {
	Sticker string `json:"sticker"`
}

// DeleteStickerFromSet delete a sticker from a set created by the bot. Returns
// True on success.
func (bot *Bot) DeleteStickerFromSet(sticker string) (ok bool, err error) {
	dst, err := json.Marshal(&DeleteStickerFromSetParameters{Sticker: sticker})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodDeleteStickerFromSet)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
