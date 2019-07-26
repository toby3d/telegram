package telegram

import (
	"strings"

	http "github.com/valyala/fasthttp"
)

type AddStickerToSetParameters struct {
	// User identifier of sticker set owner
	UserID int `json:"user_id"`

	// Sticker set name
	Name string `json:"name"`

	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	PNGSticker interface{} `json:"png_sticker"`

	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// AddStickerToSet add a new sticker to a set created by the bot. Returns True
// on success.
func (b *Bot) AddStickerToSet(params *AddStickerToSetParameters) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.SetUint("user_id", params.UserID)

	if !strings.HasSuffix(strings.ToLower(params.Name), strings.ToLower("_by_"+b.Username)) {
		params.Name = params.Name + "_by_" + b.Username
	}

	args.Set("emojis", params.Emojis)

	if params.MaskPosition != nil {
		mp, err := parser.Marshal(params.MaskPosition)
		if err != nil {
			return false, err
		}

		args.SetBytesV("mask_position", mp)
	}

	resp, err := b.Upload(MethodAddStickerToSet, TypeSticker, "sticker", params.PNGSticker, args)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
