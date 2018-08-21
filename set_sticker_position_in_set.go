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
func (bot *Bot) SetStickerPositionInSet(sticker string, position int) (ok bool, err error) {
	dst, err := json.Marshal(&SetStickerPositionInSetParameters{
		Sticker:  sticker,
		Position: position,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSetStickerPositionInSet)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
