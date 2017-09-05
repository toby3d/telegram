package telegram

import (
	"errors"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type AnswerInlineQueryParameters struct {
	// The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time"` // optional

	// Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	IsPersonal bool `json:"is_personal"` // optional

	// Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	NextOffset string `json:"next_offset"` // optional

	// If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPrivateMessageText string `json:"switch_pm_text"` // optional

	// Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	SwitchPrivateMessageParameter string `json:"switch_pm_parameter"` // optional
}

// AnswerInlineQuery send answers to an inline query. On success, True is returned.
//
// No more than 50 results per query are allowed.
func (bot *Bot) AnswerInlineQuery(id string, results []interface{}, params *AnswerInlineQueryParameters) (bool, error) {
	var args http.Args
	args.Add("inline_query_id", id) // Unique identifier for the query to be answered

	data, err := json.Marshal(results)
	if err != nil {
		return false, err
	}
	args.Add("results", string(data)) // A JSON-serialized array of results for the inline query

	args.Add("cache_time", strconv.Itoa(params.CacheTime))
	args.Add("is_personal", strconv.FormatBool(params.IsPersonal))
	args.Add("next_offset", params.NextOffset)
	args.Add("switch_pm_text", params.SwitchPrivateMessageText)
	args.Add("switch_pm_parameter", params.SwitchPrivateMessageParameter)

	resp, err := bot.request("answerInlineQuery", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
