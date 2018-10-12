package telegram

type (
	// Sticker represents a sticker.
	Sticker struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Emoji associated with the sticker
		Emoji string `json:"emoji,omitempty"`

		// Name of the sticker set to which the sticker belongs
		SetName string `json:"set_name,omitempty"`

		// Sticker width
		Width int `json:"width"`

		// Sticker height
		Height int `json:"height"`

		// File size
		FileSize int `json:"file_size,omitempty"`

		// Sticker thumbnail in the .webp or .jpg format
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// For mask stickers, the position where the mask should be placed
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	// StickerSet represents a sticker set.
	StickerSet struct {
		// Sticker set name
		Name string `json:"name"`

		// Sticker set title
		Title string `json:"title"`

		// True, if the sticker set contains masks
		ContainsMasks bool `json:"contains_masks"`

		// List of all set stickers
		Stickers []Sticker `json:"stickers"`
	}

	// MaskPosition describes the position on faces where a mask should be placed
	// by default.
	MaskPosition struct {
		// The part of the face relative to which the mask should be placed. One
		// of "forehead", "eyes", "mouth", or "chin".
		Point string `json:"point"`

		// Shift by X-axis measured in widths of the mask scaled to the face
		// size, from left to right. For example, choosing -1.0 will place mask
		// just to the left of the default mask position.
		XShift float32 `json:"x_shift"`

		// Shift by Y-axis measured in heights of the mask scaled to the face
		// size, from top to bottom. For example, 1.0 will place the mask just
		// below the default mask position.
		YShift float32 `json:"y_shift"`

		// Mask scaling coefficient. For example, 2.0 means double size.
		Scale float32 `json:"scale"`
	}
)
