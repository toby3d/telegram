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

func (iqra *InlineQueryResultArticle) ResultID() string {
	return iqra.ID
}

func (iqra *InlineQueryResultArticle) ResultType() string {
	return iqra.Type
}

func (iqra *InlineQueryResultArticle) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqra.ReplyMarkup
}

func (iqrp *InlineQueryResultPhoto) ResultID() string {
	return iqrp.ID
}

func (iqrp *InlineQueryResultPhoto) ResultType() string {
	return iqrp.Type
}

func (iqrp *InlineQueryResultPhoto) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrp.ReplyMarkup
}

func (iqrg *InlineQueryResultGif) ResultID() string {
	return iqrg.ID
}

func (iqrg *InlineQueryResultGif) ResultType() string {
	return iqrg.Type
}

func (iqrg *InlineQueryResultGif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrg.ReplyMarkup
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultID() string {
	return iqrm4g.ID
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultType() string {
	return iqrm4g.Type
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrm4g.ReplyMarkup
}

func (iqrv *InlineQueryResultVideo) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVideo) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVideo) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqra *InlineQueryResultAudio) ResultID() string {
	return iqra.ID
}

func (iqra *InlineQueryResultAudio) ResultType() string {
	return iqra.Type
}

func (iqra *InlineQueryResultAudio) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqra.ReplyMarkup
}

func (iqrv *InlineQueryResultVoice) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVoice) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVoice) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqrd *InlineQueryResultDocument) ResultID() string {
	return iqrd.ID
}

func (iqrd *InlineQueryResultDocument) ResultType() string {
	return iqrd.Type
}

func (iqrd *InlineQueryResultDocument) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrd.ReplyMarkup
}

func (iqrl *InlineQueryResultLocation) ResultID() string {
	return iqrl.ID
}

func (iqrl *InlineQueryResultLocation) ResultType() string {
	return iqrl.Type
}

func (iqrl *InlineQueryResultLocation) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrl.ReplyMarkup
}

func (iqrv *InlineQueryResultVenue) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVenue) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVenue) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqrc *InlineQueryResultContact) ResultID() string {
	return iqrc.ID
}

func (iqrc *InlineQueryResultContact) ResultType() string {
	return iqrc.Type
}

func (iqrc *InlineQueryResultContact) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrc.ReplyMarkup
}

func (iqrg *InlineQueryResultGame) ResultID() string {
	return iqrg.ID
}

func (iqrg *InlineQueryResultGame) ResultType() string {
	return iqrg.Type
}

func (iqrg *InlineQueryResultGame) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrg.ReplyMarkup
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultID() string {
	return iqrcp.ID
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultType() string {
	return iqrcp.Type
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcp.ReplyMarkup
}

func (iqrcg *InlineQueryResultCachedGif) ResultID() string {
	return iqrcg.ID
}

func (iqrcg *InlineQueryResultCachedGif) ResultType() string {
	return iqrcg.Type
}

func (iqrcg *InlineQueryResultCachedGif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcg.ReplyMarkup
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultID() string {
	return iqrcm4g.ID
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultType() string {
	return iqrcm4g.Type
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcm4g.ReplyMarkup
}

func (iqrcs *InlineQueryResultCachedSticker) ResultID() string {
	return iqrcs.ID
}

func (iqrcs *InlineQueryResultCachedSticker) ResultType() string {
	return iqrcs.Type
}

func (iqrcs *InlineQueryResultCachedSticker) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcs.ReplyMarkup
}

func (iqrcd *InlineQueryResultCachedDocument) ResultID() string {
	return iqrcd.ID
}

func (iqrcd *InlineQueryResultCachedDocument) ResultType() string {
	return iqrcd.Type
}

func (iqrcd *InlineQueryResultCachedDocument) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcd.ReplyMarkup
}

func (iqrcv *InlineQueryResultCachedVideo) ResultID() string {
	return iqrcv.ID
}

func (iqrcv *InlineQueryResultCachedVideo) ResultType() string {
	return iqrcv.Type
}

func (iqrcv *InlineQueryResultCachedVideo) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcv.ReplyMarkup
}

func (iqrcv *InlineQueryResultCachedVoice) ResultID() string {
	return iqrcv.ID
}

func (iqrcv *InlineQueryResultCachedVoice) ResultType() string {
	return iqrcv.Type
}

func (iqrcv *InlineQueryResultCachedVoice) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcv.ReplyMarkup
}

func (iqrca *InlineQueryResultCachedAudio) ResultID() string {
	return iqrca.ID
}

func (iqrca *InlineQueryResultCachedAudio) ResultType() string {
	return iqrca.Type
}

func (iqrca *InlineQueryResultCachedAudio) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrca.ReplyMarkup
}
