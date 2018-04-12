package telegram

import json "github.com/pquerna/ffjson/ffjson"

type DeleteStickerFromSetParameters struct {
	Sticker string `json:"sticker"`
}

// DeleteStickerFromSet delete a sticker from a set created by the bot. Returns
// True on success.
func (bot *Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	dst, err := json.Marshal(&DeleteStickerFromSetParameters{Sticker: sticker})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodDeleteStickerFromSet)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
