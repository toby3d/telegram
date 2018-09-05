package telegram

// NewInlineQueryResultCachedAudio creates a new inline query result with cached
// audio.
func NewInlineQueryResultCachedAudio(resultID, fileID string) *InlineQueryResultCachedAudio {
	return &InlineQueryResultCachedAudio{
		Type:        TypeAudio,
		ID:          resultID,
		AudioFileID: fileID,
	}
}

// NewInlineQueryResultCachedDocument creates a new inline query result with
// cached document.
func NewInlineQueryResultCachedDocument(resultID, fileID, title string) *InlineQueryResultCachedDocument {
	return &InlineQueryResultCachedDocument{
		Type:           TypeDocument,
		ID:             resultID,
		Title:          title,
		DocumentFileID: fileID,
	}
}

// NewInlineQueryResultCachedGif creates a new inline query result with cached
// GIF.
func NewInlineQueryResultCachedGif(resultID, fileID string) *InlineQueryResultCachedGif {
	return &InlineQueryResultCachedGif{
		Type:      TypeGIF,
		ID:        resultID,
		GifFileID: fileID,
	}
}

// NewInlineQueryResultCachedMpeg4Gif creates a new inline query result with
// cached MPEG GIF.
func NewInlineQueryResultCachedMpeg4Gif(resultID, fileID string) *InlineQueryResultCachedMpeg4Gif {
	return &InlineQueryResultCachedMpeg4Gif{
		Type:        TypeMpeg4Gif,
		ID:          resultID,
		Mpeg4FileID: fileID,
	}
}

// NewInlineQueryResultCachedPhoto creates a new inline query result with cached
// photo.
func NewInlineQueryResultCachedPhoto(resultID, fileID string) *InlineQueryResultCachedPhoto {
	return &InlineQueryResultCachedPhoto{
		Type:        TypePhoto,
		ID:          resultID,
		PhotoFileID: fileID,
	}
}

// NewInlineQueryResultCachedSticker creates a new inline query result with
// cached sticker.
func NewInlineQueryResultCachedSticker(resultID, fileID string) *InlineQueryResultCachedSticker {
	return &InlineQueryResultCachedSticker{
		Type:          TypeSticker,
		ID:            resultID,
		StickerFileID: fileID,
	}
}

// NewInlineQueryResultCachedVideo creates a new inline query result with cached
// video.
func NewInlineQueryResultCachedVideo(resultID, fileID, title string) *InlineQueryResultCachedVideo {
	return &InlineQueryResultCachedVideo{
		Type:        TypeVideo,
		ID:          resultID,
		VideoFileID: fileID,
		Title:       title,
	}
}

// NewInlineQueryResultCachedVoice creates a new inline query result with cached
// voice.
func NewInlineQueryResultCachedVoice(resultID, fileID, title string) *InlineQueryResultCachedVoice {
	return &InlineQueryResultCachedVoice{
		Type:        TypeVoice,
		ID:          resultID,
		VoiceFileID: fileID,
		Title:       title,
	}
}

// NewInlineQueryResultArticle creates a new inline query result with article.
func NewInlineQueryResultArticle(resultID, title string, content interface{}) *InlineQueryResultArticle {
	return &InlineQueryResultArticle{
		Type:                TypeArticle,
		ID:                  resultID,
		Title:               title,
		InputMessageContent: content,
	}
}

// NewInlineQueryResultAudio creates a new inline query result with audio.
func NewInlineQueryResultAudio(resultID, audioURL, title string) *InlineQueryResultAudio {
	return &InlineQueryResultAudio{
		Type:     TypeAudio,
		ID:       resultID,
		AudioURL: audioURL,
		Title:    title,
	}
}

// NewInlineQueryResultContact creates a new inline query result with contact.
func NewInlineQueryResultContact(resultID, phoneNumber, firstName string) *InlineQueryResultContact {
	return &InlineQueryResultContact{
		Type:        TypeContact,
		ID:          resultID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewInlineQueryResultGame creates a new inline query result with game.
func NewInlineQueryResultGame(resultID, gameShortName string) *InlineQueryResultGame {
	return &InlineQueryResultGame{
		Type:          TypeGame,
		ID:            resultID,
		GameShortName: gameShortName,
	}
}

// NewInlineQueryResultDocument creates a new inline query result with document.
func NewInlineQueryResultDocument(resultID, title, documentURL, mimeType string) *InlineQueryResultDocument {
	return &InlineQueryResultDocument{
		Type:        TypeDocument,
		ID:          resultID,
		Title:       title,
		DocumentURL: documentURL,
		MimeType:    mimeType,
	}
}

// NewInlineQueryResultGif creates a new inline query result with GIF.
func NewInlineQueryResultGif(resultID, gifURL, thumbURL string) *InlineQueryResultGif {
	return &InlineQueryResultGif{
		Type:     TypeGIF,
		ID:       resultID,
		GifURL:   gifURL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultLocation creates a new inline query result with location.
func NewInlineQueryResultLocation(resultID, title string, latitude, longitude float32) *InlineQueryResultLocation {
	return &InlineQueryResultLocation{
		Type:      TypeLocation,
		ID:        resultID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
	}
}

// NewInlineQueryResultMpeg4Gif creates a new inline query result with MPEG GIF.
func NewInlineQueryResultMpeg4Gif(resultID, mpeg4URL, thumbURL string) *InlineQueryResultMpeg4Gif {
	return &InlineQueryResultMpeg4Gif{
		Type:     TypeMpeg4Gif,
		ID:       resultID,
		Mpeg4URL: mpeg4URL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultPhoto creates a new inline query result with photo.
func NewInlineQueryResultPhoto(resultID, photoURL, thumbURL string) *InlineQueryResultPhoto {
	return &InlineQueryResultPhoto{
		Type:     TypePhoto,
		ID:       resultID,
		PhotoURL: photoURL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultVenue creates a new inline query result with venue.
func NewInlineQueryResultVenue(resultID, title, address string, latitude, longitude float32) *InlineQueryResultVenue {
	return &InlineQueryResultVenue{
		Type:      TypeVenue,
		ID:        resultID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// NewInlineQueryResultVideo creates a new inline query result with video.
func NewInlineQueryResultVideo(resultID, videoURL, mimeType, thumbURL, title string) *InlineQueryResultVideo {
	return &InlineQueryResultVideo{
		Type:     TypeVideo,
		ID:       resultID,
		VideoURL: videoURL,
		MimeType: mimeType,
		ThumbURL: thumbURL,
		Title:    title,
	}
}

// NewInlineQueryResultVoice creates a new inline query result with voice.
func NewInlineQueryResultVoice(resultID, voiceURL, title string) *InlineQueryResultVoice {
	return &InlineQueryResultVoice{
		Type:     TypeVoice,
		ID:       resultID,
		VoiceURL: voiceURL,
		Title:    title,
	}
}
