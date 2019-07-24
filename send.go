package telegram

import (
	"strconv"

	http "github.com/valyala/fasthttp"
)

type (
	// SendAnimationParameters represents data for SendAnimation method.
	SendAnimationParameters struct {
		// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
		ChatID int64 `json:"chat_id"`

		// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
		Animation InputFile `json:"animation"`

		// Duration of sent animation in seconds
		Duration int `json:"duration,omitempty"`

		// Animation width
		Width int `json:"width,omitempty"`

		// Animation height
		Height int `json:"height,omitempty"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
		Thumb InputFile `json:"thumb,omitempty"`

		// Animation caption (may also be used when resending animation by file_id), 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	// SendChatActionParameters represents data for SendChatAction method.
	SendChatActionParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Type of action to broadcast
		Action string `json:"action"`
	}

	// SendContactParameters represents data for SendContact method.
	SendContactParameters struct {
		// Unique identifier for the target private chat
		ChatID int64 `json:"chat_id"`

		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name"`

		// Additional data about the contact in the form of a vCard, 0-2048 bytes
		VCard string `json:"vcard,omitempty"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
		// price' button will be shown. If not empty, the first button must be a Pay
		// button.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SendDocumentParameters represents data for SendDocument method.
	SendDocumentParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// File to send. Pass a file_id as String to send a file that exists on the Telegram servers
		// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or
		// upload a new one using multipart/form-data.
		Document InputFile `json:"document"`

		// Document caption (may also be used when resending documents by file_id), 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
		// keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	// SendInvoiceParameters represents data for SendInvoice method.
	SendInvoiceParameters struct {
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

	// SendLocationParameters represents data for SendLocation method.
	SendLocationParameters struct {
		// Unique identifier for the target private chat
		ChatID int64 `json:"chat_id"`

		// Latitude of the location
		Latitude float32 `json:"latitude"`

		// Longitude of the location
		Longitude float32 `json:"longitude"`

		// Period in seconds for which the location will be updated (see Live
		// Locations), should be between 60 and 86400.
		LivePeriod int `json:"live_period,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
		// price' button will be shown. If not empty, the first button must be a Pay
		// button.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SendMediaGroupParameters represents data for SendMediaGroup method.
	SendMediaGroupParameters struct {
		// Unique identifier for the target chat.
		ChatID int64 `json:"chat_id"`

		// A JSON-serialized array describing photos and videos to be sent, must
		// include 2–10 items
		Media []interface{} `json:"media"`

		// Sends the messages silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the messages are a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	}

	// SendMessageParameters represents data for SendMessage method.
	SendMessageParameters struct {
		// Unique identifier for the target chat or username of the target channel
		// (in the format @channelusername)
		ChatID int64 `json:"chat_id"`

		// Text of the message to be sent
		Text string `json:"text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode,omitempty"`

		// Disables link previews for links in this message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline
		// keyboard, custom reply keyboard, instructions to remove reply keyboard or
		// to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	// SendPhotoParameters represents data for SendPhoto method.
	SendPhotoParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Photo to send. Pass a file_id as String to send a photo that exists on the
		// Telegram servers (recommended), pass an HTTP URL as a String for Telegram
		// to get a photo from the Internet, or upload a new photo using
		// multipart/form-data.
		Photo InputFile `json:"photo"`

		// Photo caption (may also be used when resending photos by file_id), 0-200
		// characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Disables link previews for links in this message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline
		// keyboard, custom reply keyboard, instructions to remove reply keyboard or
		// to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	SendPollConfig struct {
		// Unique identifier for the target chat. A native poll can't be sent to a private chat.
		ChatID int64 `json:"chat_id"`

		// Poll question, 1-255 characters
		Question string `json:"question"`

		// List of answer options, 2-10 strings 1-100 characters each
		Options []string `json:"options"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
		// instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}

	// SendVenueParameters represents data for SendVenue method.
	SendVenueParameters struct {
		// Unique identifier for the target private chat
		ChatID int64 `json:"chat_id"`

		// Latitude of the venue
		Latitude float32 `json:"latitude"`

		// Longitude of the venue
		Longitude float32 `json:"longitude"`

		// Name of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue
		FoursquareID string `json:"foursquare_id,omitempty"`

		// Foursquare type of the venue, if known. (For example,
		// "arts_entertainment/default", "arts_entertainment/aquarium" or
		// "food/icecream".)
		FoursquareType string `json:"foursquare_type,omitempty"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
		// price' button will be shown. If not empty, the first button must be a Pay
		// button.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SendGameParameters represents data for SendGame method.
	SendGameParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Short name of the game, serves as the unique identifier for the game. Set
		// up your games via Botfather.
		GameShortName string `json:"game_short_name"`

		// Sends the message silently. Users will receive a notification with no
		// sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one ‘Play
		// game_title’ button will be shown. If not empty, the first button must
		// launch the game.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SendStickerParameters represents data for SetSticker method.
	SendStickerParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Sticker to send
		Sticker interface{} `json:"sticker"`

		// Sends the message silently. Users will receive a notification
		// with no sound
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

		// Additional interface options. A JSON-serialized object for an
		// inline keyboard, custom reply keyboard, instructions to remove
		// reply keyboard or to force a reply from the user.
		ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	}
)

// NewAnimation creates SendAnimationParameters only with required parameters.
func NewAnimation(chatID int64, animation interface{}) *SendAnimationParameters {
	return &SendAnimationParameters{
		ChatID:    chatID,
		Animation: animation,
	}
}

// NewContact creates SendContactParameters only with required parameters.
func NewContact(chatID int64, phoneNumber, firstName string) *SendContactParameters {
	return &SendContactParameters{
		ChatID:      chatID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewDocument creates SendDocumentParameters only with required parameters.
func NewDocument(chatID int64, document interface{}) *SendDocumentParameters {
	return &SendDocumentParameters{
		ChatID:   chatID,
		Document: document,
	}
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

// NewLocation creates SendLocationParameters only with required parameters.
func NewLocation(chatID int64, latitude, longitude float32) *SendLocationParameters {
	return &SendLocationParameters{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewMediaGroup creates SendMediaGroupParameters only with required parameters.
func NewMediaGroup(chatID int64, media ...interface{}) *SendMediaGroupParameters {
	return &SendMediaGroupParameters{
		ChatID: chatID,
		Media:  media,
	}
}

// NewMessage creates SendMessageParameters only with required parameters.
func NewMessage(chatID int64, text string) *SendMessageParameters {
	return &SendMessageParameters{
		ChatID: chatID,
		Text:   text,
	}
}

// NewPhoto creates SendPhotoParameters only with required parameters.
func NewPhoto(chatID int64, photo interface{}) *SendPhotoParameters {
	return &SendPhotoParameters{
		ChatID: chatID,
		Photo:  photo,
	}
}

func NewPoll(chatID int64, question string, options ...string) SendPollConfig {
	return SendPollConfig{
		ChatID:   chatID,
		Question: question,
		Options:  options,
	}
}

// NewVenue creates SendVenueParameters only with required parameters.
func NewVenue(chatID int64, latitude, longitude float32, title, address string) *SendVenueParameters {
	return &SendVenueParameters{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// NewGame creates SendGameParameters only with required parameters.
func NewGame(chatID int64, gameShortName string) *SendGameParameters {
	return &SendGameParameters{
		ChatID:        chatID,
		GameShortName: gameShortName,
	}
}

// SendAnimation send animation files (GIF or H.264/MPEG-4 AVC video without
// sound). On success, the sent Message is returned. Bots can currently send
// animation files of up to 50 MB in size, this limit may be changed in the
// future.
func (bot *Bot) SendAnimation(params *SendAnimationParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(params.ChatID, 10))

	if params.Caption != "" {
		args.Add("caption", params.Caption)
	}

	if params.ReplyMarkup != nil {
		dst, err := parser.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		args.Add("reply_markup", string(dst))
	}

	if params.ReplyToMessageID != 0 {
		args.Add("reply_to_message_id", strconv.Itoa(params.ReplyToMessageID))
	}

	args.Add("disable_notification", strconv.FormatBool(params.DisableNotification))

	resp, err := bot.Upload(MethodSendAnimation, "animation", "", params.Animation, args)
	if err != nil {
		return nil, err
	}

	var result Message
	err = parser.Unmarshal(resp.Result, &result)
	return &result, err
}

// SendChatAction tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a
// noticeable amount of time to arrive.
func (bot *Bot) SendChatAction(chatID int64, action string) (bool, error) {
	dst, err := parser.Marshal(&SendChatActionParameters{
		ChatID: chatID,
		Action: action,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSendChatAction)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SendContact send phone contacts. On success, the sent Message is returned.
func (bot *Bot) SendContact(params *SendContactParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(*params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendContact)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendDocument send general files. On success, the sent Message is returned. Bots can currently send
// files of any type of up to 50 MB in size, this limit may be changed in the future.
func (bot *Bot) SendDocument(params *SendDocumentParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(params.ChatID, 10))

	if params.Caption != "" {
		args.Add("caption", params.Caption)
	}

	if params.ReplyMarkup != nil {
		dst, err := parser.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		args.Add("reply_markup", string(dst))
	}

	if params.ReplyToMessageID != 0 {
		args.Add("reply_to_message_id", strconv.Itoa(params.ReplyToMessageID))
	}

	args.Add("disable_notification", strconv.FormatBool(params.DisableNotification))

	resp, err := bot.Upload(MethodSendDocument, "document", "", params.Document, args)
	if err != nil {
		return nil, err
	}

	var result Message
	err = parser.Unmarshal(resp.Result, &result)
	return &result, err
}

// SendInvoice send invoices. On success, the sent Message is returned.
func (bot *Bot) SendInvoice(params *SendInvoiceParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendInvoice)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendLocation send point on the map. On success, the sent Message is returned.
func (bot *Bot) SendLocation(params *SendLocationParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendLocation)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendMediaGroup send a group of photos or videos as an album. On success, an array of the sent
// Messages is returned.
func (bot *Bot) SendMediaGroup(params *SendMediaGroupParameters) (album []Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendMediaGroup)
	if err != nil {
		return
	}

	err = parser.Unmarshal(resp.Result, &album)
	return
}

// SendMessage send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(params *SendMessageParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendMessage)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendPhoto send photos. On success, the sent Message is returned.
func (bot *Bot) SendPhoto(params *SendPhotoParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(params.ChatID, 10))

	if params.Caption != "" {
		args.Add("caption", params.Caption)
	}

	if params.ReplyMarkup != nil {
		dst, err := parser.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		args.Add("reply_markup", string(dst))
	}

	if params.ReplyToMessageID != 0 {
		args.Add("reply_to_message_id", strconv.Itoa(params.ReplyToMessageID))
	}

	args.Add("disable_notification", strconv.FormatBool(params.DisableNotification))

	resp, err := bot.Upload(MethodSendPhoto, "photo", "", params.Photo, args)
	if err != nil {
		return nil, err
	}

	var result Message
	err = parser.Unmarshal(resp.Result, &result)
	return &result, err
}

// SendPoll send a native poll. A native poll can't be sent to a private chat. On success, the sent Message is
// returned.
func (b *Bot) SendPoll(params SendPollConfig) (*Message, error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := b.request(dst, MethodSendPoll)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = parser.Unmarshal(resp.Result, &msg)
	return &msg, err
}

// SendVenue send information about a venue. On success, the sent Message is returned.
func (bot *Bot) SendVenue(params *SendVenueParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendVenue)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendGame send a game. On success, the sent Message is returned.
func (bot *Bot) SendGame(params *SendGameParameters) (msg *Message, err error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendGame)
	if err != nil {
		return
	}

	msg = new(Message)
	err = parser.Unmarshal(resp.Result, msg)
	return
}

// SendSticker send .webp stickers. On success, the sent Message is returned.
func (b *Bot) SendSticker(params *SendStickerParameters) (*Message, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Set("chat_id", strconv.FormatInt(params.ChatID, 10))
	args.Set("disable_notification", strconv.FormatBool(params.DisableNotification))
	if params.ReplyToMessageID > 0 {
		args.SetUint("reply_to_message_id", params.ReplyToMessageID)
	}
	if params.ReplyMarkup != nil {
		rm, err := parser.Marshal(params.ReplyMarkup)
		if err != nil {
			return nil, err
		}

		args.SetBytesV("reply_markup", rm)
	}

	resp, err := b.Upload(MethodSendSticker, TypeSticker, "sticker", params.Sticker, args)
	if err != nil {
		return nil, err
	}

	var result Message
	err = parser.Unmarshal(resp.Result, &result)
	return &result, err
}
