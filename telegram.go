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

func (bot *Bot) request(dst []byte, method string, args *http.Args) (*Response, error) {
	method = fmt.Sprintf(APIEndpoint, bot.AccessToken, method)
	_, body, err := http.Post(dst, method, args)
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
