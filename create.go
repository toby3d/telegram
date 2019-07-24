package telegram

import (
	"strconv"
	"strings"

	http "github.com/valyala/fasthttp"
)

type CreateNewStickerSetParameters struct {
	// User identifier of created sticker set owner
	UserID int `json:"user_id"`

	// Short name of sticker set, to be used in t.me/addstickers/ URLs
	// (e.g., animals). Can contain only english letters, digits and
	// underscores. Must begin with a letter, can't contain consecutive
	// underscores and must end in “_by_<bot username>”. <bot_username>
	// is case insensitive. 1-64 characters.
	Name string `json:"name"`

	// Sticker set title, 1-64 characters
	Title string `json:"title"`

	// Png image with the sticker, must be up to 512 kilobytes in size,
	// dimensions must not exceed 512px, and either width or height must
	// be exactly 512px. Pass a file_id as a String to send a file that
	// already exists on the Telegram servers, pass an HTTP URL as a
	// String for Telegram to get a file from the Internet, or upload
	// a new one using multipart/form-data.
	PNGSticker interface{} `json:"png_sticker"`

	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`

	// A JSON-serialized object for position where the mask should be
	// placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// CreateNewStickerSet create new sticker set owned by a user. The bot will be
// able to edit the created sticker set. Returns True on success.
func (b *Bot) CreateNewStickerSet(params *CreateNewStickerSetParameters) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.SetUint("user_id", params.UserID)

	if !strings.HasSuffix(strings.ToLower(params.Name), strings.ToLower("_by_"+b.Username)) {
		params.Name = params.Name + "_by_" + b.Username
	}

	args.Set("name", params.Name)
	args.Set("title", params.Title)
	args.Set("emojis", params.Emojis)
	args.Set("contains_masks", strconv.FormatBool(params.ContainsMasks))

	if params.MaskPosition != nil {
		mp, err := parser.Marshal(params.MaskPosition)
		if err != nil {
			return false, err
		}

		args.SetBytesV("mask_position", mp)
	}

	resp, err := b.Upload(MethodCreateNewStickerSet, TypeSticker, "sticker", params.PNGSticker, args)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
