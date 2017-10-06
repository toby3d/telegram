package telegram

import (
	"log"
	"time"
)

func NewAnswerCallback(id string) *AnswerCallbackQueryParameters {
	return &AnswerCallbackQueryParameters{CallbackQueryID: id}
}

func NewAnswerInline(id string, results ...InlineQueryResult) *AnswerInlineQueryParameters {
	return &AnswerInlineQueryParameters{InlineQueryID: id, Results: results}
}

func NewAnswerPreCheckout(id string, ok bool) *AnswerPreCheckoutQueryParameters {
	return &AnswerPreCheckoutQueryParameters{PreCheckoutQueryID: id, Ok: ok}
}

func NewAnswerShipping(id string, ok bool) *AnswerShippingQueryParameters {
	return &AnswerShippingQueryParameters{ShippingQueryID: id, Ok: ok}
}

func NewMessage(chatID int64, text string) *SendMessageParameters {
	return &SendMessageParameters{ChatID: chatID, Text: text}
}

func NewInvoice(chatID int64, title, description, payload, providerToken, startParameter, currency string, prices ...LabeledPrice) *SendInvoiceParameters {
	return &SendInvoiceParameters{
		ChatID:         chatID,
		Title:          title,
		Description:    description,
		Payload:        payload,
		ProviderToken:  providerToken,
		StartParameter: startParameter,
		Currency:       currency,
		Prices:         prices,
	}
}

func NewReplyKeyboard(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{Keyboard: keyboard, ResizeKeyboard: true}
}

func NewKeyboardButtonsRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

func NewKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text}
}

func NewKeyboardButtonContact(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestContact: true}
}

func NewKeyboardButtonLocation(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestLocation: true}
}

func NewInlineKeyboard(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

func NewInlineKeyboardButtonsRow(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

func NewInlineKeyboardButtonURL(text, url string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, URL: url}
}

func NewInlineKeyboardButtonData(text, data string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackData: data}
}

func NewInlineKeyboardButtonSwitch(text, query string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQuery: query}
}

func NewInlineKeyboardButtonSwitchSelf(text, query string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: query}
}

func NewInlineKeyboardButtonGame(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackGame: &CallbackGame{}}
}

func NewInlineKeyboardButtonPay(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, Pay: true}
}

func (bot *Bot) NewUpdatesChannel(params *GetUpdatesParameters) chan *Update {
	if params == nil {
		params = &GetUpdatesParameters{
			Limit:   100,
			Timeout: 60,
		}
	}

	channel := make(chan *Update, params.Limit)

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
					channel <- &update
				}
			}
		}
	}()

	return channel
}
