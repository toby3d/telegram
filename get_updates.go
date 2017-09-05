package telegram

import (
	"fmt"
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type GetUpdatesParameters struct {
	// Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Offset int `json:"offset"` // optional

	// Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit"` // optional

	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout"` // optional

	// List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates"` // optional
}

// GetUpdates receive incoming updates using long polling (wiki). An Array of Update objects is returned.
func (bot *Bot) GetUpdates(params *GetUpdatesParameters) ([]Update, error) {
	var args http.Args

	if params.Offset != 0 {
		args.Add("offset", strconv.Itoa(params.Offset))
	}

	if params.Limit > 0 && params.Limit <= 100 {
		args.Add("limit", strconv.Itoa(params.Limit))
	} else {
		args.Add("limit", strconv.Itoa(100))
	}

	if params.Timeout > 0 {
		args.Add("timeout", strconv.Itoa(params.Timeout))
	}

	if len(params.AllowedUpdates) > 0 {
		args.Add("allowed_updates", fmt.Sprint(`["`, strings.Join(params.AllowedUpdates, `","`), `"]`))
	}

	resp, err := bot.request("getUpdates", &args)
	if err != nil {
		return nil, err
	}

	var updates []Update
	err = json.Unmarshal(*resp.Result, &updates)
	return updates, err
}
