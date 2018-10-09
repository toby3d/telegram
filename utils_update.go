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
func (b *Bot) NewLongPollingChannel(params *GetUpdatesParameters) UpdatesChannel {
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
			updates, err := b.GetUpdates(params)
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
func (b *Bot) NewWebhookChannel(setURL *http.URI, params *SetWebhookParameters, certFile, keyFile, serveAddr string) (updates UpdatesChannel) {
	if params == nil {
		params = &SetWebhookParameters{
			URL:            setURL.String(),
			MaxConnections: 40,
		}
	}

	var err error
	channel := make(chan Update, 100)
	handleFunc := func(ctx *http.RequestCtx) {
		dlog.Ln("Request path:", string(ctx.Path()))
		if !bytes.HasPrefix(ctx.Path(), setURL.Path()) {
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
			err = http.ListenAndServeTLS(serveAddr, certFile, keyFile, handleFunc)
		} else {
			dlog.Ln("Creating simple router...")
			err = http.ListenAndServe(serveAddr, handleFunc)
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	if _, err = b.SetWebhook(params); err != nil {
		log.Fatalln(err.Error())
	}

	return channel
}

// IsMessage checks that the current update is a message creation event.
func (u *Update) IsMessage() bool {
	return u != nil && u.Message != nil
}

// IsEditedMessage checks that the current update is a editing message event.
func (u *Update) IsEditedMessage() bool {
	return u != nil && u.EditedMessage != nil
}

// IsChannelPost checks that the current update is a post channel creation event.
func (u *Update) IsChannelPost() bool {
	return u != nil && u.ChannelPost != nil
}

// IsEditedChannelPost checks that the current update is a editing post channel
// event.
func (u *Update) IsEditedChannelPost() bool {
	return u != nil && u.EditedChannelPost != nil
}

// IsInlineQuery checks that the current update is a inline query update.
func (u *Update) IsInlineQuery() bool {
	return u != nil && u.InlineQuery != nil
}

// IsChosenInlineResult checks that the current update is a chosen inline result
// update.
func (u *Update) IsChosenInlineResult() bool {
	return u != nil && u.ChosenInlineResult != nil
}

// IsCallbackQuery checks that the current update is a callback query update.
func (u *Update) IsCallbackQuery() bool {
	return u != nil && u.CallbackQuery != nil
}

// IsShippingQuery checks that the current update is a shipping query update.
func (u *Update) IsShippingQuery() bool {
	return u != nil && u.ShippingQuery != nil
}

// IsPreCheckoutQuery checks that the current update is a pre checkout query
// update.
func (u *Update) IsPreCheckoutQuery() bool {
	return u != nil && u.PreCheckoutQuery != nil
}

// Type return update type for current update.
func (u *Update) Type() string {
	switch {
	case u.IsCallbackQuery():
		return UpdateCallbackQuery
	case u.IsChannelPost():
		return UpdateChannelPost
	case u.IsChosenInlineResult():
		return UpdateChosenInlineResult
	case u.IsEditedChannelPost():
		return UpdateEditedChannelPost
	case u.IsEditedMessage():
		return UpdateEditedMessage
	case u.IsInlineQuery():
		return UpdateInlineQuery
	case u.IsMessage():
		return UpdateMessage
	case u.IsPreCheckoutQuery():
		return UpdatePreCheckoutQuery
	case u.IsShippingQuery():
		return UpdateShippingQuery
	default:
		return ""
	}
}
