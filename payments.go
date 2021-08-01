package telegram

type (
	// LabeledPrice represents a portion of the price for goods or services.
	LabeledPrice struct {
		// Portion label
		Label string `json:"label"`

		// Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		Amount int `json:"amount"`
	}

	// Invoice contains basic information about an invoice.
	Invoice struct {
		// Product name
		Title string `json:"title"`

		// Product description
		Description string `json:"description"`

		// Unique bot deep-linking parameter that can be used to generate this
		// invoice
		StartParameter string `json:"start_parameter"`

		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`
	}

	// ShippingAddress represents a shipping address.
	ShippingAddress struct {
		// ISO 3166-1 alpha-2 country code
		CountryCode string `json:"country_code"`

		// State, if applicable
		State string `json:"state"`

		// City
		City string `json:"city"`

		// First line for the address
		StreetLine1 string `json:"street_line1"`

		// Second line for the address
		StreetLine2 string `json:"street_line2"`

		// Address post code
		PostCode string `json:"post_code"`
	}

	// OrderInfo represents information about an order.
	OrderInfo struct {
		// User name
		Name string `json:"name,omitempty"`

		// User's phone number
		PhoneNumber string `json:"phone_number,omitempty"`

		// User email
		Email string `json:"email,omitempty"`

		// User shipping address
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	}

	// ShippingOption represents one shipping option.
	ShippingOption struct {
		// Shipping option identifier
		ID string `json:"id"`

		// Option title
		Title string `json:"title"`

		// List of price portions
		Prices []*LabeledPrice `json:"prices"`
	}

	// SuccessfulPayment contains basic information about a successful payment.
	SuccessfulPayment struct {
		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// Identifier of the shipping option chosen by the user
		ShippingOptionID string `json:"shipping_option_id,omitempty"`

		// Telegram payment identifier
		TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

		// Provider payment identifier
		ProviderPaymentChargeID string `json:"provider_payment_charge_id"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`

		// Order info provided by the user
		OrderInfo *OrderInfo `json:"order_info,omitempty"`
	}

	// ShippingQuery contains information about an incoming shipping query.
	ShippingQuery struct {
		// Unique query identifier
		ID string `json:"id"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// User who sent the query
		From *User `json:"from"`

		// User specified shipping address
		ShippingAddress *ShippingAddress `json:"shipping_address"`
	}

	// PreCheckoutQuery contains information about an incoming pre-checkout query.
	PreCheckoutQuery struct {
		// Unique query identifier
		ID string `json:"id"`

		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// Identifier of the shipping option chosen by the user
		ShippingOptionID string `json:"shipping_option_id,omitempty"`

		// User who sent the query
		From *User `json:"from"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`

		// Order info provided by the user
		OrderInfo *OrderInfo `json:"order_info,omitempty"`
	}

	// SendInvoiceParameters represents data for SendInvoice method.
	SendInvoice struct {
		// Unique identifier for the target private chat
		ChatID int64 `json:"chat_id"`

		// Product name, 1-32 characters
		Title string `json:"title"`

		// Product description, 1-255 characters
		Description string `json:"description"`

		// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
		Payload string `json:"payload"`

		// Payments provider token, obtained via Botfather
		ProviderToken string `json:"provider_token"`

		// Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
		StartParameter string `json:"start_parameter,omitempty"`

		// Three-letter ISO 4217 currency code, see more on currencies
		Currency string `json:"currency"`

		// JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
		ProviderData string `json:"provider_data,omitempty"`

		// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
		PhotoURL string `json:"photo_url,omitempty"`

		// Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
		Prices []*LabeledPrice `json:"prices"`

		// The maximum accepted amount for tips in the smallest units of the currency (integer, not
		// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp
		// parameter in currencies.json, it shows the number of digits past the decimal point for each currency
		// (2 for the majority of currencies). Defaults to 0
		MaxTipAmount int `json:"max_tip_amount,omitempty"`

		// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
		SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

		// Photo size
		PhotoSize int `json:"photo_size,omitempty"`

		// Photo width
		PhotoWidth int `json:"photo_width,omitempty"`

		// Photo height
		PhotoHeight int `json:"photo_height,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if you require the user's full name to complete the order
		NeedName bool `json:"need_name,omitempty"`

		// Pass True, if you require the user's phone number to complete the order
		NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

		// Pass True, if you require the user's email to complete the order
		NeedEmail bool `json:"need_email,omitempty"`

		// Pass True, if you require the user's shipping address to complete the order
		NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

		// Pass True, if the final price depends on the shipping method
		IsFlexible bool `json:"is_flexible,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// AnswerShippingQueryParameters represents data for AnswerShippingQuery method.
	AnswerShippingQuery struct {
		// Unique identifier for the query to be answered
		ShippingQueryID string `json:"shipping_query_id"`

		// Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
		ErrorMessage string `json:"error_message,omitempty"`

		// Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
		Ok bool `json:"ok"`

		// Required if ok is True. A JSON-serialized array of available shipping options.
		ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`
	}

	// AnswerPreCheckoutQueryParameters represents data for AnswerPreCheckoutQuery method.
	AnswerPreCheckoutQuery struct {
		// Unique identifier for the query to be answered
		PreCheckoutQueryID string `json:"pre_checkout_query_id"`

		// Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
		ErrorMessage string `json:"error_message,omitempty"`

		// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
		Ok bool `json:"ok"`
	}
)

func NewInvoice(chatID int64, title, description, payload, providerToken, startParameter, currency string,
	prices ...*LabeledPrice) SendInvoice {
	return SendInvoice{
		ChatID:         chatID,
		Title:          title,
		Description:    description,
		Payload:        payload,
		ProviderToken:  providerToken,
		StartParameter: startParameter,
		Currency:       currency,
		Prices:         prices,
	}
}

// SendInvoice send invoices. On success, the sent Message is returned.
func (b Bot) SendInvoice(p SendInvoice) (*Message, error) {
	src, err := b.Do(MethodSendInvoice, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewAnswerShipping(shippingQueryID string, ok bool) AnswerShippingQuery {
	return AnswerShippingQuery{
		ShippingQueryID: shippingQueryID,
		Ok:              ok,
	}
}

// AnswerShippingQuery reply to shipping queries.
//
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the b. On success, True is returned.
func (b Bot) AnswerShippingQuery(p AnswerShippingQuery) (ok bool, err error) {
	src, err := b.Do(MethodAnswerShippingQuery, p)
	if err != nil {
		return false, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}

func NewAnswerPreCheckout(preCheckoutQueryID string, ok bool) AnswerPreCheckoutQuery {
	return AnswerPreCheckoutQuery{
		PreCheckoutQueryID: preCheckoutQueryID,
		Ok:                 ok,
	}
}

// AnswerPreCheckoutQuery respond to such pre-checkout queries.
//
// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned.
//
// Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (b Bot) AnswerPreCheckoutQuery(p AnswerShippingQuery) (ok bool, err error) {
	src, err := b.Do(MethodAnswerPreCheckoutQuery, p)
	if err != nil {
		return false, err
	}

	if err = parseResponseError(b.marshler, src, &ok); err != nil {
		return
	}

	return
}
