package telegram

import (
	"strconv"
	"strings"

	http "github.com/valyala/fasthttp"
)

type (
	// SetChatDescriptionParameters represents data for SetChatDescription method.
	SetChatDescriptionParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// New chat description, 0-255 characters
		Description string `json:"description"`
	}

	// SetChatPhotoParameters represents data for SetChatPhoto method.
	SetChatPhotoParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// New chat photo, uploaded using multipart/form-data
		ChatPhoto interface{} `json:"chat_photo"`
	}

	// SetChatStickerSetParameters represents data for SetChatStickerSet method.
	SetChatStickerSetParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Name of the sticker set to be set as the group sticker set
		StickerSetName string `json:"sticker_set_name"`
	}

	// SetChatTitleParameters represents data for SetChatTitle method.
	SetChatTitleParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// New chat title, 1-255 characters
		Title string `json:"title"`
	}

	SetPassportDataErrorsParameters struct {
		// User identifier
		UserID int `json:"user_id"`

		// A JSON-serialized array describing the errors
		Errors []PassportElementError `json:"errors"`
	}

	// SetWebhookParameters represents data for SetWebhook method.
	SetWebhookParameters struct {
		// HTTPS url to send updates to. Use an empty string to remove webhook
		// integration
		URL string `json:"url"`

		// Upload your public key certificate so that the root certificate in use can
		// be checked. See our self-signed guide for details.
		Certificate InputFile `json:"certificate,omitempty"`

		// Maximum allowed number of simultaneous HTTPS connections to the webhook
		// for update delivery, 1-100. Defaults to 40. Use lower values to limit the
		// load on your bot‘s server, and higher values to increase your bot’s
		// throughput.
		MaxConnections int `json:"max_connections,omitempty"`

		// List the types of updates you want your bot to receive. For example,
		// specify [“message”, “edited_channel_post”, “callback_query”] to only
		// receive updates of these types. See Update for a complete list of
		// available update types. Specify an empty list to receive all updates
		// regardless of type (default). If not specified, the previous setting will
		// be used.
		//
		// Please note that this parameter doesn't affect updates created before the
		// call to the setWebhook, so unwanted updates may be received for a short
		// period of time.
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

	// SetGameScoreParameters represents data for SetGameScore method.
	SetGameScoreParameters struct {
		// User identifier
		UserID int `json:"user_id"`

		// New score, must be non-negative
		Score int `json:"score"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Pass True, if the high score is allowed to decrease. This can be useful
		// when fixing mistakes or banning cheaters
		Force bool `json:"force,omitempty"`

		// Pass True, if the game message should not be automatically edited to
		// include the current scoreboard
		DisableEditMessage bool `json:"disable_edit_message,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}

	// SetStickerPositionInSetParameters represents data for SetStickerPositionInSet
	// method.
	SetStickerPositionInSetParameters struct {
		// File identifier of the sticker
		Sticker string `json:"sticker"`

		// New sticker position in the set, zero-based
		Position int `json:"position"`
	}
)

// NewWebhook creates new SetWebhookParameters only with required parameters.
func NewWebhook(url string, file interface{}) *SetWebhookParameters {
	return &SetWebhookParameters{
		URL:         url,
		Certificate: file,
	}
}

// NewGameScore creates SetGameScoreParameters only with required parameters.
func NewGameScore(userID, score int) *SetGameScoreParameters {
	return &SetGameScoreParameters{
		UserID: userID,
		Score:  score,
	}
}

// SetChatDescription change the description of a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatDescription(chatID int64, description string) (bool, error) {
	dst, err := parser.Marshal(&SetChatDescriptionParameters{
		ChatID:      chatID,
		Description: description,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSetChatDescription)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetChatPhoto set a new profile photo for the chat. Photos can't be changed for private chats. The
// bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the 'All Members Are
// Admins' setting is off in the target group.
func (bot *Bot) SetChatPhoto(chatID int64, chatPhoto interface{}) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("chat_id", strconv.FormatInt(chatID, 10))

	resp, err := bot.Upload(MethodSetChatPhoto, TypePhoto, "chat_photo", chatPhoto, args)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetChatStickerSet set a new group sticker set for a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success.
func (bot *Bot) SetChatStickerSet(chatID int64, stickerSetName string) (bool, error) {
	dst, err := parser.Marshal(&SetChatStickerSetParameters{
		ChatID:         chatID,
		StickerSetName: stickerSetName,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSetChatStickerSet)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetChatTitle change the title of a chat. Titles can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the
// 'All Members Are Admins' setting is off in the target group.
func (bot *Bot) SetChatTitle(chatID int64, title string) (bool, error) {
	dst, err := parser.Marshal(&SetChatTitleParameters{
		ChatID: chatID,
		Title:  title,
	})
	if err != nil {
		return false, err
	}

	resp, err := bot.request(dst, MethodSetChatTitle)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetPassportDataErrors informs a user that some of the Telegram Passport
// elements they provided contains errors. The user will not be able to re-submit
// their Passport to you until the errors are fixed (the contents of the field
// for which you returned the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your
// service requires for any reason. For example, if a birthday date seems
// invalid, a submitted document is blurry, a scan shows evidence of tampering,
// etc. Supply some details in the error message to make sure the user knows how
// to correct the issues.
func (b *Bot) SetPassportDataErrors(userId int, errors []PassportElementError) (bool, error) {
	dst, err := parser.Marshal(&SetPassportDataErrorsParameters{
		UserID: userId,
		Errors: errors,
	})
	if err != nil {
		return false, err
	}

	resp, err := b.request(dst, MethodSetPassportDataErrors)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetWebhook specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to
// the specified url, containing a JSON-serialized Update. In case of an
// unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns true.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we
// recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
func (bot *Bot) SetWebhook(params *SetWebhookParameters) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("url", params.URL)

	if len(params.AllowedUpdates) > 0 {
		args.Add("allowed_updates", strings.Join(params.AllowedUpdates, ","))
	}
	if params.MaxConnections > 0 &&
		params.MaxConnections <= 100 {
		args.Add("max_connections", strconv.Itoa(params.MaxConnections))
	}

	var resp *Response
	var err error
	if params.Certificate != nil {
		resp, err = bot.Upload(MethodSetWebhook, "certificate", "cert.pem", params.Certificate, args)
	} else {
		var dst []byte
		dst, err = parser.Marshal(params)
		if err != nil {
			return false, err
		}

		resp, err = bot.request(dst, MethodSetWebhook)
	}
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}

// SetGameScore set the score of the specified user in a game. On success, if the
// message was sent by the bot, returns the edited Message, otherwise returns
// True. Returns an error, if the new score is not greater than the user's
// current score in the chat and force is False.
func (bot *Bot) SetGameScore(params *SetGameScoreParameters) (*Message, error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodSetGameScore)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = parser.Unmarshal(resp.Result, &msg)
	return &msg, err
}

// SetStickerPositionInSet move a sticker in a set created by the bot to a
// specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	dst, err := parser.Marshal(&SetStickerPositionInSetParameters{
		Sticker:  sticker,
		Position: position,
	})
	if err != nil {
		return false, err
	}

	resp, err := b.request(dst, MethodSetStickerPositionInSet)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
