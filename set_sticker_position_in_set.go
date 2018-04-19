package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SetStickerPositionInSetParameters represents data for SetStickerPositionInSet
// method.
type SetStickerPositionInSetParameters struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

// SetStickerPositionInSet move a sticker in a set created by the bot to a
// specific position. Returns True on success.
func (bot *Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	dst, err := json.Marshal(&SetStickerPositionInSetParameters{
		Sticker:  sticker,
		Position: position,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSetStickerPositionInSet)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
