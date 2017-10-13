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
)

func (bot *Bot) request(dst []byte, method string, args *http.Args) (*Response, error) {
	requestURI := fmt.Sprintf(APIEndpoint, bot.AccessToken, method)
	if args != nil {
		requestURI += args.String()
	}

	var req http.Request
	var resp http.Response

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(dst)
	req.SetRequestURI(requestURI)

	if err := http.Do(&req, &resp); err != nil {
		return nil, err
	}

	var data Response
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return &data, errors.New(data.Description)
	}

	return &data, nil
}
