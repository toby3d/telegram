package telegram

import (
	"errors"
	"log"
	"path"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

var defaultClient = new(http.Client)

func (bot *Bot) request(dst []byte, method string) (response *Response, err error) {
	if bot.Client == nil {
		bot.SetClient(defaultClient)
	}

	requestURI := http.AcquireURI()
	requestURI.SetScheme("https")
	requestURI.SetHost("api.telegram.org")
	requestURI.SetPath(path.Join("bot"+bot.AccessToken, method))

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("POST")
	if dst == nil {
		req.Header.SetMethod("GET")
	}
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent(path.Join("telegram", strconv.FormatInt(Version, 10)))
	req.Header.SetHostBytes(requestURI.Host())
	req.SetBody(dst)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	err = bot.Client.Do(req, resp)
	log.Ln("Request:")
	log.D(req)
	log.Ln("Response:")
	log.D(resp)
	if err != nil {
		return
	}

	response = new(Response)
	if err = json.Unmarshal(resp.Body(), response); err != nil {
		return
	}

	if !response.Ok {
		err = errors.New(response.Description)
	}

	return
}
