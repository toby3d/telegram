package telegram

import (
	"net/url"
	"strconv"
)

// NewForceReply calls the response interface to the message.
func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

// NewInlineMentionURL creates a url.URL for the mention user without username.
func NewInlineMentionURL(userID int) *url.URL {
	link := &url.URL{
		Scheme: SchemeTelegram,
		Path:   "user",
	}

	q := link.Query()
	q.Add("id", strconv.Itoa(userID))
	link.RawQuery = q.Encode()

	return link
}
