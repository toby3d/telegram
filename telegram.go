package telegram

import (
	"errors"
	"fmt"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const (
	APIEndpoint  = "https://api.telegram.org/bot%s/%s"
	FileEndpoind = "https://api.telegram.org/file/bot%s/%s"

	StyleMarkdown = "markdown"
	StyleHTML     = "html"
)

type Bot struct {
	AccessToken string
}

func NewBot(accessToken string) *Bot {
	return &Bot{accessToken}
}

func (bot *Bot) get(method string, args *http.Args) (*Response, error) {
	method = fmt.Sprintf(APIEndpoint, bot.AccessToken, method)
	if args != nil {
		method += fmt.Sprint("?", args.String())
	}

	_, body, err := http.Get(nil, method)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if !resp.Ok {
		return &resp, errors.New(resp.Description)
	}

	return &resp, nil
}

func (bot *Bot) post(method string, args *http.Args) (*Response, error) {
	method = fmt.Sprintf(APIEndpoint, bot.AccessToken, method)
	_, body, err := http.Post(nil, method, args)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if !resp.Ok {
		return &resp, errors.New(resp.Description)
	}

	return &resp, nil
}

func (bot *Bot) upload(dst interface{}, method string, args *http.Args) (*Response, error) {
	return nil, nil
}
