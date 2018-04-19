package telegram

import (
	"bytes"
	"log"
	"time"

	"github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// UpdatesChannel is a channel for reading updates of bot.
type UpdatesChannel <-chan Update

// NewLongPollingChannel creates channel for receive incoming updates using long
// polling.
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
				dlog.Ln(err.Error())
				dlog.Ln("Failed to get updates, retrying in 3 seconds...")
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

// NewWebhookChannel creates channel for receive incoming updates via an outgoing
// webhook.
func (bot *Bot) NewWebhookChannel(params *SetWebhookParameters, certFile, keyFile, set, listen, serve string) (updates UpdatesChannel) {
	if params == nil {
		params = &SetWebhookParameters{
			URL:            set,
			MaxConnections: 40,
		}
	}

	var err error
	channel := make(chan Update, 100)
	requiredPath := []byte(listen)
	dlog.Ln("requiredPath:", string(requiredPath))
	handleFunc := func(ctx *http.RequestCtx) {
		dlog.Ln("Request path:", string(ctx.Path()))
		if !bytes.HasPrefix(ctx.Path(), requiredPath) {
			dlog.Ln("Unsupported request path:", string(ctx.Path()))
			return
		}
		dlog.Ln("Catched supported request path:", string(ctx.Path()))

		var update Update
		if err = json.Unmarshal(ctx.Request.Body(), &update); err != nil {
			return
		}

		channel <- update
	}

	go func() {
		if certFile != "" && keyFile != "" {
			dlog.Ln("Creating TLS router...")
			err = http.ListenAndServeTLS(serve, certFile, keyFile, handleFunc)
		} else {
			dlog.Ln("Creating simple router...")
			err = http.ListenAndServe(serve, handleFunc)
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	if _, err = bot.SetWebhook(params); err != nil {
		log.Fatalln(err.Error())
	}

	return channel
}

// IsMessage checks that the current update is a message creation event.
func (upd *Update) IsMessage() bool {
	return upd != nil && upd.Message != nil
}

// IsEditedMessage checks that the current update is a editing message event.
func (upd *Update) IsEditedMessage() bool {
	return upd != nil && upd.EditedMessage != nil
}

// IsChannelPost checks that the current update is a post channel creation event.
func (upd *Update) IsChannelPost() bool {
	return upd != nil && upd.ChannelPost != nil
}

// IsEditedChannelPost checks that the current update is a editing post channel
// event.
func (upd *Update) IsEditedChannelPost() bool {
	return upd != nil && upd.EditedChannelPost != nil
}

// IsInlineQuery checks that the current update is a inline query update.
func (upd *Update) IsInlineQuery() bool {
	return upd != nil && upd.InlineQuery != nil
}

// IsChosenInlineResult checks that the current update is a chosen inline result
// update.
func (upd *Update) IsChosenInlineResult() bool {
	return upd != nil && upd.ChosenInlineResult != nil
}

// IsCallbackQuery checks that the current update is a callback query update.
func (upd *Update) IsCallbackQuery() bool {
	return upd != nil && upd.CallbackQuery != nil
}

// IsShippingQuery checks that the current update is a shipping query update.
func (upd *Update) IsShippingQuery() bool {
	return upd != nil && upd.ShippingQuery != nil
}

// IsPreCheckoutQuery checks that the current update is a pre checkout query
// update.
func (upd *Update) IsPreCheckoutQuery() bool {
	return upd != nil && upd.PreCheckoutQuery != nil
}

// Type return update type for current update.
func (upd *Update) Type() string {
	switch {
	case upd.IsCallbackQuery():
		return UpdateCallbackQuery
	case upd.IsChannelPost():
		return UpdateChannelPost
	case upd.IsChosenInlineResult():
		return UpdateChosenInlineResult
	case upd.IsEditedChannelPost():
		return UpdateEditedChannelPost
	case upd.IsEditedMessage():
		return UpdateEditedMessage
	case upd.IsInlineQuery():
		return UpdateInlineQuery
	case upd.IsMessage():
		return UpdateMessage
	case upd.IsPreCheckoutQuery():
		return UpdatePreCheckoutQuery
	case upd.IsShippingQuery():
		return UpdateShippingQuery
	default:
		return ""
	}
}
