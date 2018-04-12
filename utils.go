package telegram

import (
	"net/url"
	"strconv"
)

func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

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
