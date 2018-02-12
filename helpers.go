package telegram

import (
	"net/url"
	"strconv"
)

func NewForceReply(selective bool) *ForceReply {
	return &ForceReply{
		ForceReply: true,
		Selective:  selective,
	}
}

func NewInlineMentionURL(id int) *url.URL {
	link := &url.URL{
		Scheme: "tg",
		Path:   "user",
	}

	q := link.Query()
	q.Add("id", strconv.Itoa(id))
	link.RawQuery = q.Encode()

	return link
}
