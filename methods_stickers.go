package telegram

import (
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type (
	// SendStickerParameters represents data for SetSticker method.
	SendStickerParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Sticker to send
		Sticker interface{} `json:"sticker"`

		// Sends the message silently. Users will receive a notification
		// with no sound
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an
		// inline keyboard, custom reply keyboard, instructions to remove
		// reply keyboard or to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	// GetStickerSetParameters represents data for GetStickerSet method.
	GetStickerSetParameters struct {
		// Name of the sticker set
		Name string `json:"name"`
	}

	UploadStickerFileParameters struct {
		// User identifier of sticker file owner
		UserID int `json:"user_id"`

		// Png image with the sticker, must be up to 512 kilobytes in size,
		// dimensions must not exceed 512px, and either width or height
		// must be exactly 512px.
		PNGSticker interface{} `json:"png_sticker"`
	}

	CreateNewStickerSetParameters struct {
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

	AddStickerToSetParameters struct {
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

	// SetStickerPositionInSetParameters represents data for SetStickerPositionInSet
	// method.
	SetStickerPositionInSetParameters struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`

		// New sticker position in the set, zero-based
		Position int `json:"position"`
	}

	// DeleteStickerFromSetParameters represents data for DeleteStickerFromSet method.
	DeleteStickerFromSetParameters struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`
	}
)

// SendSticker send .webp stickers. On success, the sent Message is returned.
func (b *Bot) SendSticker(params *SendStickerParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Set("chat_id", strconv.FormatInt(params.ChatID, 10))
	args.Set("disable_notification", strconv.FormatBool(params.DisableNotification))
	if params.ReplyToMessageID > 0 {
		args.SetUint("reply_to_message_id", params.ReplyToMessageID)
	}
	if params.ReplyMarkup != nil {
		rm, err := json.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}

		args.SetBytesV("reply_markup", rm)
	}

	resp, err := b.Upload(MethodSendSticker, TypeSticker, "sticker", params.Sticker, args)
	if err != nil {
		return nil, err
	}

	var m Message
	err = json.Unmarshal(*resp.Result, &m)
	return &m, err
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

// UploadStickerFile upload a .png file with a sticker for later use in
// createNewStickerSet and addStickerToSet methods (can be used multiple times).
// Returns the uploaded File on success.
func (b *Bot) UploadStickerFile(userID int, pngSticker interface{}) (*File, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.SetUint("user_id", userID)

	resp, err := b.Upload(MethodUploadStickerFile, TypeSticker, "sticker", pngSticker, args)
	if err != nil {
		return nil, err
	}

	var f File
	err = json.Unmarshal(*resp.Result, &f)
	return &f, err
}

// CreateNewStickerSet create new sticker set owned by a user. The bot will be
// able to edit the created sticker set. Returns True on success.
func (b *Bot) CreateNewStickerSet(params *CreateNewStickerSetParameters) (ok bool, err error) {
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
		mp, err := json.Marshal(params.MaskPosition)
		if err != nil {
			return false, err
		}

		args.SetBytesV("mask_position", mp)
	}

	resp, err := b.Upload(MethodCreateNewStickerSet, TypeSticker, "sticker", params.PNGSticker, args)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}

// AddStickerToSet add a new sticker to a set created by the bot. Returns True
// on success.
func (b *Bot) AddStickerToSet(params *AddStickerToSetParameters) (ok bool, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.SetUint("user_id", params.UserID)

	if !strings.HasSuffix(strings.ToLower(params.Name), strings.ToLower("_by_"+b.Username)) {
		params.Name = params.Name + "_by_" + b.Username
	}

	args.Set("emojis", params.Emojis)

	if params.MaskPosition != nil {
		mp, err := json.Marshal(params.MaskPosition)
		if err != nil {
			return false, err
		}

		args.SetBytesV("mask_position", mp)
	}

	resp, err := b.Upload(MethodAddStickerToSet, TypeSticker, "sticker", params.PNGSticker, args)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}

// SetStickerPositionInSet move a sticker in a set created by the bot to a
// specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (ok bool, err error) {
	dst, err := json.Marshal(&SetStickerPositionInSetParameters{
		Sticker:  sticker,
		Position: position,
	})
	if err != nil {
		return
	}

	resp, err := b.request(dst, MethodSetStickerPositionInSet)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
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
