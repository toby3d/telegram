package telegram

type (
	// LabeledPrice represents a portion of the price for goods or services.
	LabeledPrice struct {
		// Portion label
		Label string `json:"label"`

		//      Price of the product in the smallest units of the currency (integer,
		// not float/double). For example, for a price of US$ 1.45 pass amount =
		// 145. See the exp parameter in currencies.json, it shows the number of
		// digits past the decimal point for each currency (2 for the majority
		// of currencies).
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

		// Total price in the smallest units of the currency (integer, not
		// float/double). For example, for a price of US$ 1.45 pass amount = 145.
		// See the exp parameter in currencies.json, it shows the number of
		// digits past the decimal point for each currency (2 for the majority
		// of currencies).
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
		Prices []LabeledPrice `json:"prices"`
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

		// Total price in the smallest units of the currency (integer, not
		// float/double). For example, for a price of US$ 1.45 pass amount = 145.
		// See the exp parameter in currencies.json, it shows the number of
		// digits past the decimal point for each currency (2 for the majority
		// of currencies).
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

		// Total price in the smallest units of the currency (integer, not
		// float/double). For example, for a price of US$ 1.45 pass amount = 145.
		// See the exp parameter in currencies.json, it shows the number of
		// digits past the decimal point for each currency (2 for the majority of
		// currencies).
		TotalAmount int `json:"total_amount"`

		// Order info provided by the user
		OrderInfo *OrderInfo `json:"order_info,omitempty"`
	}
)
