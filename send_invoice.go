package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendInvoiceParameters represents data for SendInvoice method.
type SendInvoiceParameters struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"`

	// Product name, 1-32 characters
	Title string `json:"title"`

	// Product description, 1-255 characters
	Description string `json:"description"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to
	// the user, use for your internal processes.
	Payload string `json:"payload"`

	// Payments provider token, obtained via Botfather
	ProviderToken string `json:"provider_token"`

	// Unique deep-linking parameter that can be used to generate this invoice
	// when used as a start parameter
	StartParameter string `json:"start_parameter"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`

	// JSON-encoded data about the invoice, which will be shared with the payment
	// provider. A detailed description of required fields should be provided by
	// the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// URL of the product photo for the invoice. Can be a photo of the goods or a
	// marketing image for a service. People like it better when they see what
	// they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`

	// Price breakdown, a list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice `json:"prices"`

	// Photo size
	PhotoSize int `json:"photo_size,omitempty"`

	// Photo width
	PhotoWidth int `json:"photo_width,omitempty"`

	// Photo height
	PhotoHeight int `json:"photo_height,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Pass True, if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// Pass True, if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Pass True, if you require the user's email to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// Pass True, if you require the user's shipping address to complete the
	// order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`

	// Sends the message silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
	// price' button will be shown. If not empty, the first button must be a Pay
	// button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// NewInvoice creates SendInvoiceParameters only with required parameters.
func NewInvoice(chatID int64, title, description, payload, providerToken, startParameter, currency string, prices ...LabeledPrice) *SendInvoiceParameters {
	return &SendInvoiceParameters{
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
func (bot *Bot) SendInvoice(params *SendInvoiceParameters) (*Message, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodSendInvoice)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
