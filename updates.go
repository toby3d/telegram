package telegram

import (
	"strconv"
	"strings"
	"time"

	http "github.com/valyala/fasthttp"
)

type (
	// Update represents an incoming update.
	//
	// At most one of the optional parameters can be present in any given update.
	Update struct {
		// The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
		UpdateID int `json:"update_id"`

		// New incoming message of any kind — text, photo, sticker, etc.
		Message *Message `json:"message,omitempty"`

		// New version of a message that is known to the bot and was edited
		EditedMessage *Message `json:"edited_message,omitempty"`

		// New incoming channel post of any kind — text, photo, sticker, etc.
		ChannelPost *Message `json:"channel_post,omitempty"`

		// New version of a channel post that is known to the bot and was edited
		EditedChannelPost *Message `json:"adited_channel_post,omitempty"`

		// New incoming inline query
		InlineQuery *InlineQuery `json:"inline_query,omitempty"`

		// The result of an inline query that was chosen by a user and sent to their chat partner.
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

		// New incoming callback query
		CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

		// New incoming shipping query. Only for invoices with flexible price
		ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

		// New incoming pre-checkout query. Contains full information about checkout
		PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

		// New poll state. Bots receive only updates about polls, which are sent or stopped by the bot
		Poll *Poll `json:"poll,omitempty"`

		// A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were
		// sent by the bot itself.
		PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	}

	// WebhookInfo contains information about the current status of a webhook.
	WebhookInfo struct {
		// Webhook URL, may be empty if webhook is not set up
		URL string `json:"url"`

		// Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
		LastErrorMessage string `json:"last_error_message,omitempty"`

		// True, if a custom certificate was provided for webhook certificate checks
		HasCustomCertificate bool `json:"has_custom_certificate"`

		// Number of updates awaiting delivery
		PendingUpdateCount int `json:"pending_update_count"`

		// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
		MaxConnections int `json:"max_connections,omitempty"`

		// Unix time for the most recent error that happened when trying to deliver an update via webhook
		LastErrorDate int64 `json:"last_error_date,omitempty"`

		// A list of update types the bot is subscribed to. Defaults to all update types
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

	// GetUpdatesParameters represents data for GetUpdates method.
	GetUpdates struct {
		// Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
		Offset int `json:"offset,omitempty"`

		// Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`

		// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
		Timeout int `json:"timeout,omitempty"`

		// List the types of updates you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.
		//
		// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

	// SetWebhookParameters represents data for SetWebhook method.
	SetWebhook struct {
		// HTTPS url to send updates to. Use an empty string to remove webhook integration
		URL string `json:"url"`

		// Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
		Certificate InputFile `json:"certificate,omitempty"`

		// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
		MaxConnections int `json:"max_connections,omitempty"`

		// List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.
		//
		// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}
)

// GetUpdates receive incoming updates using long polling. An Array of Update objects is returned.
func (b Bot) GetUpdates(p *GetUpdates) ([]*Update, error) {
	src, err := b.Do(MethodGetUpdates, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := make([]*Update, 0)
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetWebhook specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns true.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL, e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
func (b Bot) SetWebhook(p SetWebhook) (bool, error) {
	if p.Certificate.IsAttachment() {
		_, err := b.Upload(MethodSetWebhook, map[string]string{
			"allowed_updates": strings.Join(p.AllowedUpdates, ","),
			"max_connections": strconv.Itoa(p.MaxConnections),
			"url":             p.URL,
		}, &p.Certificate)

		return err == nil, err
	}

	src, err := b.Do(MethodSetWebhook, p)
	if err != nil {
		return false, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return false, err
	}

	var result bool
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return false, err
	}

	return result, nil
}

// DeleteWebhook remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
func (b Bot) DeleteWebhook() (bool, error) {
	src, err := b.Do(MethodDeleteWebhook, nil)
	if err != nil {
		return false, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return false, err
	}

	var result bool
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return false, err
	}

	return result, nil
}

// GetWebhookInfo get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (b Bot) GetWebhookInfo() (*WebhookInfo, error) {
	src, err := b.Do(MethodGetWebhookInfo, nil)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(WebhookInfo)
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// IsMessage checks that the current update is a message creation event.
func (u Update) IsMessage() bool { return u.Message != nil }

// IsEditedMessage checks that the current update is a editing message event.
func (u Update) IsEditedMessage() bool { return u.EditedMessage != nil }

// IsChannelPost checks that the current update is a post channel creation event.
func (u Update) IsChannelPost() bool { return u.ChannelPost != nil }

// IsEditedChannelPost checks that the current update is a editing post channel event.
func (u Update) IsEditedChannelPost() bool { return u.EditedChannelPost != nil }

// IsInlineQuery checks that the current update is a inline query update.
func (u Update) IsInlineQuery() bool { return u.InlineQuery != nil }

// IsChosenInlineResult checks that the current update is a chosen inline result update.
func (u Update) IsChosenInlineResult() bool { return u.ChosenInlineResult != nil }

// IsCallbackQuery checks that the current update is a callback query update.
func (u Update) IsCallbackQuery() bool { return u.CallbackQuery != nil }

// IsShippingQuery checks that the current update is a shipping query update.
func (u Update) IsShippingQuery() bool { return u.ShippingQuery != nil }

// IsPreCheckoutQuery checks that the current update is a pre checkout query update.
func (u Update) IsPreCheckoutQuery() bool { return u.PreCheckoutQuery != nil }

// IsPoll checks that the current update is a poll update.
func (u Update) IsPoll() bool { return u.Poll != nil }

// Type return update type for current update.
func (u Update) Type() string {
	switch {
	case u.IsCallbackQuery():
		return UpdateCallbackQuery
	case u.IsChannelPost():
		return UpdateChannelPost
	case u.IsChosenInlineResult():
		return UpdateChosenInlineResult
	case u.IsEditedChannelPost():
		return UpdateEditedChannelPost
	case u.IsEditedMessage():
		return UpdateEditedMessage
	case u.IsInlineQuery():
		return UpdateInlineQuery
	case u.IsMessage():
		return UpdateMessage
	case u.IsPreCheckoutQuery():
		return UpdatePreCheckoutQuery
	case u.IsShippingQuery():
		return UpdateShippingQuery
	case u.IsPoll():
		return UpdatePoll
	default:
		return ""
	}
}

func (w WebhookInfo) LastErrorTime() time.Time { return time.Unix(w.LastErrorDate, 0) }

func (w WebhookInfo) HasURL() bool { return w.URL != "" }

func (w WebhookInfo) URI() *http.URI {
	if !w.HasURL() {
		return nil
	}

	u := http.AcquireURI()
	u.Update(w.URL)

	return u
}
