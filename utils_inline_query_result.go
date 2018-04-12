package telegram

func NewInlineQueryResultCachedAudio(resultID, fileID string) *InlineQueryResultCachedAudio {
	return &InlineQueryResultCachedAudio{
		Type:        TypeAudio,
		ID:          resultID,
		AudioFileID: fileID,
	}
}

func NewInlineQueryResultCachedDocument(resultID, fileID, title string) *InlineQueryResultCachedDocument {
	return &InlineQueryResultCachedDocument{
		Type:           TypeDocument,
		ID:             resultID,
		Title:          title,
		DocumentFileID: fileID,
	}
}

func NewInlineQueryResultCachedGif(resultID, fileID string) *InlineQueryResultCachedGif {
	return &InlineQueryResultCachedGif{
		Type:      TypeGIF,
		ID:        resultID,
		GifFileID: fileID,
	}
}

func NewInlineQueryResultCachedMpeg4Gif(resultID, fileID string) *InlineQueryResultCachedMpeg4Gif {
	return &InlineQueryResultCachedMpeg4Gif{
		Type:        TypeMpeg4Gif,
		ID:          resultID,
		Mpeg4FileID: fileID,
	}
}

func NewInlineQueryResultCachedPhoto(resultID, fileID string) *InlineQueryResultCachedPhoto {
	return &InlineQueryResultCachedPhoto{
		Type:        TypePhoto,
		ID:          resultID,
		PhotoFileID: fileID,
	}
}

func NewInlineQueryResultCachedSticker(resultID, fileID string) *InlineQueryResultCachedSticker {
	return &InlineQueryResultCachedSticker{
		Type:          TypeSticker,
		ID:            resultID,
		StickerFileID: fileID,
	}
}

func NewInlineQueryResultCachedVideo(resultID, fileID, title string) *InlineQueryResultCachedVideo {
	return &InlineQueryResultCachedVideo{
		Type:        TypeVideo,
		ID:          resultID,
		VideoFileID: fileID,
		Title:       title,
	}
}

func NewInlineQueryResultCachedVoice(resultID, fileID, title string) *InlineQueryResultCachedVoice {
	return &InlineQueryResultCachedVoice{
		Type:        TypeVoice,
		ID:          resultID,
		VoiceFileID: fileID,
		Title:       title,
	}
}

func NewInlineQueryResultArticle(resultID, title string, content *InputMessageContent) *InlineQueryResultArticle {
	return &InlineQueryResultArticle{
		Type:                TypeArticle,
		ID:                  resultID,
		Title:               title,
		InputMessageContent: content,
	}
}

func NewInlineQueryResultAudio(resultID, audioURL, title string) *InlineQueryResultAudio {
	return &InlineQueryResultAudio{
		Type:     TypeAudio,
		ID:       resultID,
		AudioURL: audioURL,
		Title:    title,
	}
}

func NewInlineQueryResultContact(resultID, phoneNumber, firstName string) *InlineQueryResultContact {
	return &InlineQueryResultContact{
		Type:        TypeContact,
		ID:          resultID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

func NewInlineQueryResultGame(resultID, gameShortName string) *InlineQueryResultGame {
	return &InlineQueryResultGame{
		Type:          TypeGame,
		ID:            resultID,
		GameShortName: gameShortName,
	}
}

func NewInlineQueryResultDocument(resultID, title, documentURL, mimeType string) *InlineQueryResultDocument {
	return &InlineQueryResultDocument{
		Type:        TypeDocument,
		ID:          resultID,
		Title:       title,
		DocumentURL: documentURL,
		MimeType:    mimeType,
	}
}

func NewInlineQueryResultGif(resultID, gifURL, thumbURL string) *InlineQueryResultGif {
	return &InlineQueryResultGif{
		Type:     TypeGIF,
		ID:       resultID,
		GifURL:   gifURL,
		ThumbURL: thumbURL,
	}
}

func NewInlineQueryResultLocation(resultID, title string, latitude, longitude float32) *InlineQueryResultLocation {
	return &InlineQueryResultLocation{
		Type:      TypeLocation,
		ID:        resultID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
	}
}

func NewInlineQueryResultMpeg4Gif(resultID, mpeg4URL, thumbURL string) *InlineQueryResultMpeg4Gif {
	return &InlineQueryResultMpeg4Gif{
		Type:     TypeMpeg4Gif,
		ID:       resultID,
		Mpeg4URL: mpeg4URL,
		ThumbURL: thumbURL,
	}
}

func NewInlineQueryResultPhoto(resultID, photoURL, thumbURL string) *InlineQueryResultPhoto {
	return &InlineQueryResultPhoto{
		Type:     TypePhoto,
		ID:       resultID,
		PhotoURL: photoURL,
		ThumbURL: thumbURL,
	}
}

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

func NewInlineQueryResultVoice(resultID, voiceURL, title string) *InlineQueryResultVoice {
	return &InlineQueryResultVoice{
		Type:     TypeVoice,
		ID:       resultID,
		VoiceURL: voiceURL,
		Title:    title,
	}
}
