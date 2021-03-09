package telegram

import (
	"strconv"
	"strings"
)

type (
	// Sticker represents a sticker.
	Sticker struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Sticker width
		Width int `json:"width"`

		// Sticker height
		Height int `json:"height"`

		// true, if the sticker is animated
		IsAnimated bool `json:"is_animated"`

		// Sticker thumbnail in the .webp or .jpg format
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// Emoji associated with the sticker
		Emoji string `json:"emoji,omitempty"`

		// Name of the sticker set to which the sticker belongs
		SetName string `json:"set_name,omitempty"`

		// For mask stickers, the position where the mask should be placed
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// StickerSet represents a sticker set.
	StickerSet struct {
		// Sticker set name
		Name string `json:"name"`

		// Sticker set title
		Title string `json:"title"`

		// List of all set stickers
		Stickers []*Sticker `json:"stickers"`

		// True, if the sticker set contains masks
		ContainsMasks bool `json:"contains_masks"`

		// true, if the sticker set contains animated stickers
		IsAnimated bool `json:"is_animated"`

		// Sticker set thumbnail in the .WEBP or .TGS format
		Thumb *PhotoSize `json:"thumb,omitempty"`
	}

	// MaskPosition describes the position on faces where a mask should be placed by default.
	MaskPosition struct {
		// The part of the face relative to which the mask should be placed. One of "forehead", "eyes", "mouth", or "chin".
		Point string `json:"point"`

		// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
		XShift float64 `json:"x_shift"`

		// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
		YShift float64 `json:"y_shift"`

		// Mask scaling coefficient. For example, 2.0 means double size.
		Scale float64 `json:"scale"`
	}

	// SendStickerParameters represents data for SetSticker method.
	SendSticker struct {
		ChatID ChatID `json:"chat_id"`

		// Sticker to send
		Sticker *InputFile `json:"sticker"`

		// Sends the message silently. Users will receive a notification with no sound
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// GetStickerSetParameters represents data for GetStickerSet method.
	GetStickerSet struct {
		// Name of the sticker set
		Name string `json:"name"`
	}

	UploadStickerFile struct {
		// User identifier of sticker file owner
		UserID int64 `json:"user_id"`

		// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
		PNGSticker *InputFile `json:"png_sticker"`
	}

	CreateNewStickerSet struct {
		// User identifier of created sticker set owner
		UserID int64 `json:"user_id"`

		// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
		Name string `json:"name"`

		// Sticker set title, 1-64 characters
		Title string `json:"title"`

		// PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
		PNGSticker *InputFile `json:"png_sticker,omitempty"`

		// TGS animation with the sticker, uploaded using multipart/form-data.
		// See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
		TGSSticker *InputFile `json:"tgs_sticker,omitempty"`

		// One or more emoji corresponding to the sticker
		Emojis string `json:"emojis"`

		// Pass True, if a set of mask stickers should be created
		ContainsMasks bool `json:"contains_masks,omitempty"`

		// A JSON-serialized object for position where the mask should be placed on faces
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	AddStickerToSet struct {
		// User identifier of sticker set owner
		UserID int64 `json:"user_id"`

		// Sticker set name
		Name string `json:"name"`

		// PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
		PNGSticker *InputFile `json:"png_sticker"`

		// TGS animation with the sticker, uploaded using multipart/form-data.
		// See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
		TGSSticker *InputFile `json:"tgs_sticker,omitempty"`

		// One or more emoji corresponding to the sticker
		Emojis string `json:"emojis"`

		// A JSON-serialized object for position where the mask should be placed on faces
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	// SetStickerPositionInSetParameters represents data for SetStickerPositionInSet method.
	SetStickerPositionInSet struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`

		// New sticker position in the set, zero-based
		Position int `json:"position"`
	}

	// DeleteStickerFromSetParameters represents data for DeleteStickerFromSet method.
	DeleteStickerFromSet struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`
	}

	// SetStickerSetThumb represents data for SetStickerSetThumb method.
	SetStickerSetThumb struct {
		// Sticker set name
		Name string `json:"name"`

		// User identifier of the sticker set owner
		UserID int64 `json:"user_id"`

		// A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and height
		// exactly 100px, or a TGS animation with the thumbnail up to 32 kilobytes in size;
		// see https://core.telegram.org/animated_stickers#technical-requirements for animated sticker
		// technical requirements. Pass a file_id as a String to send a file that already exists on the
		// Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or
		// upload a new one using multipart/form-data. More info on Sending Files ». Animated sticker set
		// thumbnail can't be uploaded via HTTP URL.
		Thumb *InputFile `json:"thumb,omitempty"`
	}
)

func NewSticker(chatID ChatID, sticker *InputFile) SendSticker {
	return SendSticker{
		ChatID:  chatID,
		Sticker: sticker,
	}
}

// SendSticker send .webp stickers. On success, the sent Message is returned.
func (b Bot) SendSticker(p SendSticker) (*Message, error) {
	src, err := b.Do(MethodSendSticker, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (b Bot) GetStickerSet(name string) (*StickerSet, error) {
	src, err := b.Do(MethodGetStickerSet, GetStickerSet{Name: name})
	if err != nil {
		return nil, err
	}

	result := new(StickerSet)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// UploadStickerFile upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
func (b Bot) UploadStickerFile(uid int, sticker *InputFile) (*File, error) {
	params := make(map[string]string)
	params["user_id"] = strconv.Itoa(uid)

	var err error
	if params["png_sticker"], err = b.marshler.MarshalToString(sticker); err != nil {
		return nil, err
	}

	src, err := b.Upload(MethodUploadStickerFile, params, sticker)
	if err != nil {
		return nil, err
	}

	result := new(File)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewStickerSet(userID int64, name, title string, pngSticker *InputFile, emojis ...string) CreateNewStickerSet {
	return CreateNewStickerSet{
		UserID:     userID,
		Name:       name,
		Title:      title,
		PNGSticker: pngSticker,
		Emojis:     strings.Join(emojis, ""),
	}
}

// CreateNewStickerSet create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
func (b *Bot) CreateNewStickerSet(p CreateNewStickerSet) (ok bool, err error) {
	params := make(map[string]string)
	params["user_id"] = strconv.FormatInt(p.UserID, 10)
	params["name"] = p.Name
	params["title"] = p.Title
	params["emojis"] = p.Emojis
	params["contains_masks"] = strconv.FormatBool(p.ContainsMasks)

	if params["png_sticker"], err = b.marshler.MarshalToString(p.PNGSticker); err != nil {
		return
	}

	if params["mask_position"], err = b.marshler.MarshalToString(p.MaskPosition); err != nil {
		return
	}

	files := make([]*InputFile, 0)
	if p.PNGSticker.IsAttachment() {
		files = append(files, p.PNGSticker)
	}

	src, err := b.Upload(MethodCreateNewStickerSet, params, files...)
	if err != nil {
		return ok, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

// AddStickerToSet add a new sticker to a set created by the b. Returns True on success.
func (b *Bot) AddStickerToSet(p AddStickerToSet) (ok bool, err error) {
	params := make(map[string]string)
	params["user_id"] = strconv.FormatInt(p.UserID, 10)
	params["name"] = p.Name
	params["emojis"] = p.Emojis

	if params["png_sticker"], err = b.marshler.MarshalToString(p.PNGSticker); err != nil {
		return
	}

	if params["mask_position"], err = b.marshler.MarshalToString(p.MaskPosition); err != nil {
		return
	}

	files := make([]*InputFile, 0)
	if p.PNGSticker.IsAttachment() {
		files = append(files, p.PNGSticker)
	}

	src, err := b.Upload(MethodAddStickerToSet, params, files...)
	if err != nil {
		return ok, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

// SetStickerPositionInSet move a sticker in a set created by the bot to a specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (ok bool, err error) {
	src, err := b.marshler.Marshal(&SetStickerPositionInSet{
		Sticker:  sticker,
		Position: position,
	})
	if err != nil {
		return ok, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

// DeleteStickerFromSet delete a sticker from a set created by the b. Returns True on success.
func (b *Bot) DeleteStickerFromSet(sticker string) (ok bool, err error) {
	src, err := b.Do(MethodDeleteStickerFromSet, DeleteStickerFromSet{Sticker: sticker})
	if err != nil {
		return ok, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

// SetStickerSetThumb set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets
// only. Returns True on success.
func (b *Bot) SetStickerSetThumb(p SetStickerSetThumb) (ok bool, err error) {
	params := make(map[string]string)
	params["name"] = p.Name
	params["user_id"] = strconv.FormatInt(p.UserID, 10)

	if params["thumb"], err = b.marshler.MarshalToString(p.Thumb); err != nil {
		return
	}

	files := make([]*InputFile, 0)
	if p.Thumb.IsAttachment() {
		files = append(files, p.Thumb)
	}

	src, err := b.Upload(MethodSetStickerSetThumb, params, files...)
	if err != nil {
		return ok, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

// InSet checks that the current sticker in the stickers set.
func (s Sticker) InSet() bool { return s.SetName != "" }

func (s Sticker) HasThumb() bool { return s.Thumb != nil }

func (s Sticker) IsMask() bool { return s.MaskPosition != nil }

func (s Sticker) File() File {
	return File{
		FileID:       s.FileID,
		FileUniqueID: s.FileUniqueID,
		FileSize:     s.FileSize,
	}
}
