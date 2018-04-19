package telegram

import "net/url"

var defaultURI = &url.URL{
	Scheme: "https",
	Host:   "api.telegram.org",
}
