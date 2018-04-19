package telegram

// InSet checks that the current sticker in the stickers set.
//
// For uploaded WebP files this return false.
func (sticker *Sticker) InSet() bool {
	return sticker != nil && sticker.SetName != ""
}

// Set use bot for getting parent StickerSet if SetName is present.
//
// Return nil if current sticker has been uploaded by user as WebP file.
func (sticker *Sticker) Set(bot *Bot) *StickerSet {
	if !sticker.InSet() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(sticker.SetName)
	if err != nil {
		return nil
	}

	return set
}
