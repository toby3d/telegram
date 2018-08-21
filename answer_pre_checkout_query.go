package telegram

import json "github.com/pquerna/ffjson/ffjson"

// AnswerPreCheckoutQueryParameters represents data for AnswerPreCheckoutQuery
// method.
type AnswerPreCheckoutQueryParameters struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Required if ok is False. Error message in human readable form that
	// explains the reason for failure to proceed with the checkout (e.g. "Sorry,
	// somebody just bought the last of our amazing black T-shirts while you were
	// busy filling out your payment details. Please choose a different color or
	// garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`

	// Specify True if everything is alright (goods are available, etc.) and the
	// bot is ready to proceed with the order. Use False if there are any
	// problems.
	Ok bool `json:"ok"`
}

// NewAnswerPreCheckoutQuery creates AnswerPreCheckoutQueryParameters only with
// required parameters.
func NewAnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool) *AnswerPreCheckoutQueryParameters {
	return &AnswerPreCheckoutQueryParameters{
		PreCheckoutQueryID: preCheckoutQueryID,
		Ok:                 ok,
	}
}

// AnswerPreCheckoutQuery respond to such pre-checkout queries.
//
// Once the user has confirmed their payment and shipping details, the Bot API
// sends the final confirmation in the form of an Update with the field
// pre_checkout_query. Use this method to respond to such pre-checkout queries.
// On success, True is returned.
//
// Note: The Bot API must receive an answer within 10 seconds after the
// pre-checkout query was sent.
func (bot *Bot) AnswerPreCheckoutQuery(params *AnswerShippingQueryParameters) (ok bool, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerPreCheckoutQuery)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
