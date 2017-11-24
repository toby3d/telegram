package telegram

import (
	"strings"
	"time"
	// "net/url"

	log "github.com/kirillDanshin/dlog"
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
				log.Ln(err.Error())
				log.Ln("Failed to get updates, retrying in 3 seconds...")
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

	if _, err := bot.SetWebhook(params); err != nil {
		panic(err.Error())
	}

	channel := make(chan Update, params.MaxConnections)
	go func() {
		var err error
		if certFile != "" && keyFile != "" {
			log.Ln("Creating TLS router...")
			err = http.ListenAndServeTLS(
				serve,
				certFile,
				keyFile,
				func(ctx *http.RequestCtx) {
					if !strings.HasPrefix(string(ctx.Path()), listen) {
						log.Ln("Unsupported request path:", string(ctx.Path()))
						return
					}

					log.Ln("Catched supported request path:", string(ctx.Path()))
					var update Update
					err = json.Unmarshal(ctx.Request.Body(), &update)
					if err != nil {
						log.Ln(err.Error())
						return
					}

					channel <- update
				},
			)
		} else {
			log.Ln("Creating simple router...")
			err = http.ListenAndServe(
				serve,
				func(ctx *http.RequestCtx) {
					if !strings.HasPrefix(string(ctx.Path()), listen) {
						log.Ln("Unsupported request path:", string(ctx.Path()))
						return
					}

					log.Ln("Catched supported request path:", string(ctx.Path()))
					var update Update
					err = json.Unmarshal(ctx.Request.Body(), &update)
					if err != nil {
						log.Ln(err.Error())
						return
					}

					channel <- update
				},
			)
		}
		if err != nil {
			panic(err.Error())
		}
	}()

	return channel
}
