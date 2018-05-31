package telegram

import (
	"net/url"

	http "github.com/valyala/fasthttp"
)

var (
	defaultClient = new(http.Client)
	defaultURI    = &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
	}
)
