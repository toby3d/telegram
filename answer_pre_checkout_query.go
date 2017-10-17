package telegram

import json "github.com/pquerna/ffjson/ffjson"

type AnswerPreCheckoutQueryParameters struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Specify True if everything is alright (goods are available, etc.) and the
	// bot is ready to proceed with the order. Use False if there are any
	// problems.
	Ok bool `json:"ok"`

	// Required if ok is False. Error message in human readable form that
	// explains the reason for failure to proceed with the checkout (e.g. "Sorry,
	// somebody just bought the last of our amazing black T-shirts while you were
	// busy filling out your payment details. Please choose a different color or
	// garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

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
func (bot *Bot) AnswerPreCheckoutQuery(params *AnswerShippingQueryParameters) (bool, error) {
	dst, err := json.Marshal(*params)
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, "answerPreCheckoutQuery", nil)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
