package telegram

import (
	gojson "encoding/json"
	"errors"
	"path"

	json "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

// Response represents a response from the Telegram API with the result
// stored raw. If ok equals true, the request was successful, and the result
// of the query can be found in the result field. In case of an unsuccessful
// request, ok equals false, and the error is explained in the error field.
type Response struct {
	Ok          bool                `json:"ok"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Result      gojson.RawMessage   `json:"result,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

var (
	defaultClient = http.Client{}
	parser        = json.ConfigFastest
)

func (b *Bot) request(dst []byte, method string) (*Response, error) {
	if b.Client == nil {
		b.SetClient(&defaultClient)
	}

	requestURI := http.AcquireURI()
	requestURI.SetScheme("https")
	requestURI.SetHost("api.telegram.org")
	requestURI.SetPath(path.Join("bot"+b.AccessToken, method))

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(http.MethodPost)
	if dst == nil {
		req.Header.SetMethod(http.MethodGet)
	}
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent(path.Join("telegram", Version))
	req.Header.SetHostBytes(requestURI.Host())
	req.SetBody(dst)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := b.Client.Do(req, resp); err != nil {
		return nil, err
	}

	var data Response
	if err := parser.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data, nil
}
