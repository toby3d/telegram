package telegram

import (
	"errors"
	"fmt"
	"net/url"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

func (bot *Bot) request(dst []byte, method string, args *http.Args) (*Response, error) {
	requestURI := &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("/bot", bot.AccessToken, "/", method),
	}

	if args != nil {
		requestURI.RawQuery = args.String()
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.Header.SetHost("api.telegram.org")
	req.SetBody(dst)

	log.Ln("Request:")
	log.D(*req)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	if err := http.Do(req, resp); err != nil {
		log.Ln("Resp:")
		log.D(*resp)

		return nil, err
	}

	log.Ln("Resp:")
	log.D(*resp)

	var data Response
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return &data, errors.New(data.Description)
	}

	return &data, nil
}
