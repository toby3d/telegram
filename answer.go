package telegram

type (
	// AnswerCallbackQueryParameters represents data for AnswerCallbackQuery method.
	AnswerCallbackQueryParameters struct {
		// Unique identifier for the query to be answered
		CallbackQueryID string `json:"callback_query_id"`

		// Text of the notification. If not specified, nothing will be shown to the
		// user, 0-200 characters
		Text string `json:"text,omitempty"`

		// URL that will be opened by the user's client. If you have created a Game
		// and accepted the conditions via @Botfather, specify the URL that opens
		// your game – note that this will only work if the query comes from a
		// callback_game button.
		//
		// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your
		// bot with a parameter.
		URL string `json:"url,omitempty"`

		// If true, an alert will be shown by the client instead of a notification at
		// the top of the chat screen. Defaults to false.
		ShowAlert bool `json:"show_alert,omitempty"`

		// The maximum amount of time in seconds that the result of the callback
		// query may be cached client-side. Telegram apps will support caching
		// starting in version 3.14. Defaults to 0.
		CacheTime int `json:"cache_time,omitempty"`
	}

	// AnswerPreCheckoutQueryParameters represents data for AnswerPreCheckoutQuery
	// method.
	AnswerPreCheckoutQueryParameters struct {
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

	// AnswerShippingQueryParameters represents data for AnswerShippingQuery method.
	AnswerShippingQueryParameters struct {
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

	// AnswerInlineQueryParameters represents data for AnswerInlineQuery method.
	AnswerInlineQueryParameters struct {
		// Unique identifier for the answered query
		InlineQueryID string `json:"inline_query_id"`

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

		// A JSON-serialized array of results for the inline query
		Results []interface{} `json:"results"`

		// The maximum amount of time in seconds that the result of the inline query
		// may be cached on the server. Defaults to 300.
		CacheTime int `json:"cache_time,omitempty"`

		// Pass True, if results may be cached on the server side only for the user
		// that sent the query. By default, results may be returned to any user who
		// sends the same query
		IsPersonal bool `json:"is_personal,omitempty"`
	}
)

// NewAnswerCallbackQuery creates AnswerCallbackQueryParameters only with
// required parameters.
func NewAnswerCallbackQuery(callbackQueryID string) *AnswerCallbackQueryParameters {
	return &AnswerCallbackQueryParameters{CallbackQueryID: callbackQueryID}
}

// NewAnswerPreCheckoutQuery creates AnswerPreCheckoutQueryParameters only with
// required parameters.
func NewAnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool) *AnswerPreCheckoutQueryParameters {
	return &AnswerPreCheckoutQueryParameters{
		PreCheckoutQueryID: preCheckoutQueryID,
		Ok:                 ok,
	}
}

// NewAnswerShippingQuery creates AnswerShippingQueryParameters only with
// required parameters.
func NewAnswerShippingQuery(shippingQueryID string, ok bool) *AnswerShippingQueryParameters {
	return &AnswerShippingQueryParameters{
		ShippingQueryID: shippingQueryID,
		Ok:              ok,
	}
}

// NewAnswerInlineQuery creates AnswerInlineQueryParameters only with required
// parameters.
func NewAnswerInlineQuery(inlineQueryID string, results ...interface{}) *AnswerInlineQueryParameters {
	return &AnswerInlineQueryParameters{
		InlineQueryID: inlineQueryID,
		Results:       results,
	}
}

// AnswerCallbackQuery send answers to callback queries sent from inline
// keyboards. The answer will be displayed to the user as a notification at the
// top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this
// option to work, you must first create a game for your bot via @Botfather and
// accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX
// that open your bot with a parameter.
func (bot *Bot) AnswerCallbackQuery(params *AnswerCallbackQueryParameters) (ok bool, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerCallbackQuery)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
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
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerPreCheckoutQuery)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
}

// AnswerShippingQuery reply to shipping queries.
//
// If you sent an invoice requesting a shipping address and the parameter
// is_flexible was specified, the Bot API will send an Update with a
// shipping_query field to the bot. On success, True is returned.
func (bot *Bot) AnswerShippingQuery(params *AnswerShippingQueryParameters) (ok bool, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerShippingQuery)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
}

// AnswerInlineQuery send answers to an inline query. On success, True is returned.
//
// No more than 50 results per query are allowed.
func (bot *Bot) AnswerInlineQuery(params *AnswerInlineQueryParameters) (ok bool, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerInlineQuery)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &ok)
	return
}
