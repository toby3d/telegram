package telegram

import (
	"log"
	"net/url"
	"time"

	router "github.com/buaazp/fasthttprouter"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type UpdatesChannel <-chan *Update

func NewAnswerCallback(id string) *AnswerCallbackQueryParameters {
	return &AnswerCallbackQueryParameters{
		CallbackQueryID: id}
}

func NewAnswerInline(id string, results ...InlineQueryResult) *AnswerInlineQueryParameters {
	return &AnswerInlineQueryParameters{
		InlineQueryID: id,
		Results:       results,
	}
}

func NewAnswerPreCheckout(id string, ok bool) *AnswerPreCheckoutQueryParameters {
	return &AnswerPreCheckoutQueryParameters{
		PreCheckoutQueryID: id,
		Ok:                 ok,
	}
}

func NewAnswerShipping(id string, ok bool) *AnswerShippingQueryParameters {
	return &AnswerShippingQueryParameters{
		ShippingQueryID: id,
		Ok:              ok,
	}
}

func NewMessage(chatID int64, text string) *SendMessageParameters {
	return &SendMessageParameters{
		ChatID: chatID,
		Text:   text,
	}
}

func NewMessageForward(from, to int64, messageID int) *ForwardMessageParameters {
	return &ForwardMessageParameters{
		FromChatID: from,
		ChatID:     to,
		MessageID:  messageID,
	}
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

func NewInlineKeyboard(rows ...[]InlineKeyboardButton) [][]InlineKeyboardButton {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return keyboard
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
	return InlineKeyboardButton{
		Text:         text,
		CallbackGame: &CallbackGame{},
	}
}

func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}

func NewWebhook(url string, file interface{}) *SetWebhookParameters {
	var input InputFile
	input = file
	return &SetWebhookParameters{
		URL:         url,
		Certificate: &input,
	}
}

func NewInlineKeyboardButtonPay(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, Pay: true}
}

func (bot *Bot) NewLongPollingChannel(params *GetUpdatesParameters) UpdatesChannel {
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

func (msg *Message) IsCommand() bool {
	if len(msg.Entities) <= 0 {
		return false
	}

	if msg.Entities[0].Type != EntityBotCommand &&
		msg.Entities[0].Offset != 0 {
		return false
	}

	return true
}

func (msg *Message) Command() string {
	if len(msg.Entities) <= 0 {
		return ""
	}

	if msg.Entities[0].Type != EntityBotCommand &&
		msg.Entities[0].Offset != 0 {
		return ""
	}

	return string([]rune(msg.Text)[:msg.Entities[0].Length])
}

func (chat *Chat) IsPrivate() bool {
	return chat.Type == ChatPrivate
}

func (chat *Chat) IsGroup() bool {
	return chat.Type == ChatGroup
}

func (chat *Chat) IsSuperGroup() bool {
	return chat.Type == ChatSuperGroup
}

func (chat *Chat) IsChannel() bool {
	return chat.Type == ChatChannel
}

func (entity *MessageEntity) ParseURL() (*url.URL, error) {
	if entity.Type != EntityTextLink {
		return nil, nil
	}

	return url.Parse(entity.URL)
}
