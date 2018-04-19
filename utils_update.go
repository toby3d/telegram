package telegram

import (
	"bytes"
	"log"
	"time"

	"github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type UpdatesChannel <-chan Update

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

func (upd *Update) IsMessage() bool {
	return upd != nil && upd.Message != nil
}

func (upd *Update) IsEditedMessage() bool {
	return upd != nil && upd.EditedMessage != nil
}

func (upd *Update) IsChannelPost() bool {
	return upd != nil && upd.ChannelPost != nil
}

func (upd *Update) IsEditedChannelPost() bool {
	return upd != nil && upd.EditedChannelPost != nil
}

func (upd *Update) IsInlineQuery() bool {
	return upd != nil && upd.InlineQuery != nil
}

func (upd *Update) IsChosenInlineResult() bool {
	return upd != nil && upd.ChosenInlineResult != nil
}

func (upd *Update) IsCallbackQuery() bool {
	return upd != nil && upd.CallbackQuery != nil
}

func (upd *Update) IsShippingQuery() bool {
	return upd != nil && upd.ShippingQuery != nil
}

func (upd *Update) IsPreCheckoutQuery() bool {
	return upd != nil && upd.PreCheckoutQuery != nil
}
