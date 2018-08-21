package telegram

import json "github.com/pquerna/ffjson/ffjson"

// AnswerShippingQueryParameters represents data for AnswerShippingQuery method.
type AnswerShippingQueryParameters struct {
	// Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id"`

	// Required if ok is False. Error message in human readable form that
	// explains why it is impossible to complete the order (e.g. "Sorry, delivery
	// to your desired address is unavailable'). Telegram will display this
	// message to the user.
	ErrorMessage string `json:"error_message,omitempty"`

	// Specify True if delivery to the specified address is possible and False
	// if there are any problems (for example, if delivery to the specified
	// address is not possible)
	Ok bool `json:"ok"`

	// Required if ok is True. A JSON-serialized array of available shipping
	// options.
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
}

// NewAnswerShippingQuery creates AnswerShippingQueryParameters only with
// required parameters.
func NewAnswerShippingQuery(shippingQueryID string, ok bool) *AnswerShippingQueryParameters {
	return &AnswerShippingQueryParameters{
		ShippingQueryID: shippingQueryID,
		Ok:              ok,
	}
}

// AnswerShippingQuery reply to shipping queries.
//
// If you sent an invoice requesting a shipping address and the parameter
// is_flexible was specified, the Bot API will send an Update with a
// shipping_query field to the bot. On success, True is returned.
func (bot *Bot) AnswerShippingQuery(params *AnswerShippingQueryParameters) (ok bool, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerShippingQuery)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
