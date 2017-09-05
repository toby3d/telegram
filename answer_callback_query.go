package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type AnswerCallbackQueryParameters struct {
	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string `json:"text"` // optional

	// If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert"` // optional

	// URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url"` // optional

	// The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time"` // optional
}

// AnswerCallbackQuery send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @Botfather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
func (bot *Bot) AnswerCallbackQuery(id string, params *AnswerCallbackQueryParameters) (bool, error) {
	var args http.Args
	args.Add("callback_query_id", id) // Unique identifier for the query to be answered
	args.Add("text", params.Text)
	args.Add("show_alert", strconv.FormatBool(params.ShowAlert))
	args.Add("url", params.URL)
	args.Add("cache_time", strconv.Itoa(params.CacheTime))

	resp, err := bot.request("answerCallbackQuery", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
