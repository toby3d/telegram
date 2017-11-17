package telegram

import (
	"log"
	"strings"
	"time"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type UpdatesChannel <-chan Update

func NewForceReply(selective bool) *ForceReply {
	return &ForceReply{
		ForceReply: true,
		Selective:  selective,
	}
}

func NewReplyKeyboardMarkup(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}

func NewReplyKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

func NewReplyKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

func NewReplyKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

func NewReplyKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

func NewInlineKeyboardButton(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

func NewInlineKeyboardButtonSwitch(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

func NewInlineKeyboardButtonSwitchSelf(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		SwitchInlineQueryCurrentChat: query,
	}
}

func NewInlineKeyboardButtonGame(text string) InlineKeyboardButton {
	var game CallbackGame
	return InlineKeyboardButton{
		Text:         text,
		CallbackGame: &game,
	}
}

func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}

func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

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

func NewInputTextMessageContent(messageText string) *InputTextMessageContent {
	return &InputTextMessageContent{
		MessageText: messageText,
	}
}

func NewInputLocationMessageContent(latitude, longitude float32) *InputLocationMessageContent {
	return &InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func NewInputVenueMessageContent(latitude, longitude float32, title, address string) *InputVenueMessageContent {
	return &InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

func NewInputContactMessageContent(phoneNumber, firstName string) *InputContactMessageContent {
	return &InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

func (bot *Bot) NewLongPollingChannel(params *GetUpdatesParameters) UpdatesChannel {
	if params == nil {
		params = &GetUpdatesParameters{
			Offset:  0,
			Limit:   100,
			Timeout: 60,
		}
	}

	channel := make(chan Update, params.Limit)

	go func() {
		for {
			updates, err := bot.GetUpdates(params)
			if err != nil {
				log.Println(err.Error())
				log.Println("failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.ID >= params.Offset {
					params.Offset = update.ID + 1
					channel <- update
				}
			}
		}
	}()

	return channel
}

func NewWebhookChannel(listen, serve string, params *GetUpdatesParameters) UpdatesChannel {
	if params == nil {
		params = &GetUpdatesParameters{
			Offset:  0,
			Limit:   100,
			Timeout: 60,
		}
	}

	channel := make(chan Update, params.Limit)

	go func() {
		if err := http.ListenAndServe(serve, func(ctx *http.RequestCtx) {
			if strings.HasPrefix(string(ctx.Path()), listen) {
				var update Update
				json.Unmarshal(ctx.Request.Body(), &update)
				channel <- update
			}
		}); err != nil {
			panic(err.Error())
		}
	}()

	return channel
}
