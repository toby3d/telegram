package telegram

// InSet checks that the current sticker in the stickers set.
//
// For uploaded WebP files this return false.
func (s *Sticker) InSet() bool {
	return s != nil && s.SetName != ""
}

// Set use bot for getting parent StickerSet if SetName is present.
//
// Return nil if current sticker has been uploaded by user as WebP file.
func (s *Sticker) Set(bot *Bot) *StickerSet {
	if !s.InSet() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(s.SetName)
	if err != nil {
		return nil
	}

	return set
}
