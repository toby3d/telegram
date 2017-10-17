package telegram

import json "github.com/pquerna/ffjson/ffjson"

type AnswerInlineQueryParameters struct {
	// Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id"`

	// A JSON-serialized array of results for the inline query
	Results []InlineQueryResult `json:"results"`

	// The maximum amount of time in seconds that the result of the inline query
	// may be cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time,omitempty"`

	// Pass True, if results may be cached on the server side only for the user
	// that sent the query. By default, results may be returned to any user who
	// sends the same query
	IsPersonal bool `json:"is_personal,omitempty"`

	// Pass the offset that a client should send in the next query with the same
	// text to receive more results. Pass an empty string if there are no more
	// results or if you don‘t support pagination. Offset length can’t exceed 64
	// bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// If passed, clients will display a button with specified text that switches
	// the user to a private chat with the bot and sends the bot a start message
	// with the parameter switch_pm_parameter
	SwitchPrivateMessageText string `json:"switch_pm_text,omitempty"`

	// Deep-linking parameter for the /start message sent to the bot when user
	// presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and -
	// are allowed.
	SwitchPrivateMessageParameter string `json:"switch_pm_parameter,omitempty"`
}

func NewAnswerInlineQuery(inlineQueryID string, results ...InlineQueryResult) *AnswerInlineQueryParameters {
	return &AnswerInlineQueryParameters{
		InlineQueryID: inlineQueryID,
		Results:       results,
	}
}

// AnswerInlineQuery send answers to an inline query. On success, True is returned.
//
// No more than 50 results per query are allowed.
func (bot *Bot) AnswerInlineQuery(params *AnswerInlineQueryParameters) (bool, error) {
	dst, err := json.Marshal(*params)
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "answerInlineQuery", nil)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
