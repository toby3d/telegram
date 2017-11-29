package telegram

import (
	"bytes"
	"log"
	"time"

	http "github.com/erikdubbelboer/fasthttp"
	"github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
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

func (bot *Bot) NewWebhookChannel(
	params *SetWebhookParameters,
	certFile, keyFile, set, listen, serve string,
) (updates UpdatesChannel) {
	if params == nil {
		params = &SetWebhookParameters{
			URL:            set,
			MaxConnections: 40,
		}
	}

	channel := make(chan Update, 100)
	go func() {
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
			if err := json.Unmarshal(ctx.Request.Body(), &update); err != nil {
				return
			}

			channel <- update
		}

		if certFile != "" && keyFile != "" {
			dlog.Ln("Creating TLS router...")
			log.Fatal(http.ListenAndServeTLS(serve, certFile, keyFile, handleFunc))
		} else {
			dlog.Ln("Creating simple router...")
			log.Fatal(http.ListenAndServe(serve, handleFunc))
		}
	}()

	if _, err := bot.SetWebhook(params); err != nil {
		log.Fatalln(err.Error())
	}

	return channel
}
