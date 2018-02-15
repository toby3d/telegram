package telegram

import (
	"errors"
	"fmt"
	"net/url"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

func (bot *Bot) request(dst []byte, method string) (*Response, error) {
	requestURI := &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("/bot", bot.AccessToken, "/", method),
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.Header.SetHost("api.telegram.org")
	req.SetBody(dst)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	err := http.Do(req, resp)
	log.Ln("Request:")
	log.D(req)
	log.Ln("Response:")
	log.D(resp)
	if err != nil {
		return nil, err
	}

	var data Response
	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return &data, errors.New(data.Description)
	}

	return &data, nil
}
