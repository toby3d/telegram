package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"

	http "github.com/valyala/fasthttp"
)

// Response represents a response from the Telegram API with the result
// stored raw. If ok equals true, the request was successful, and the result
// of the query can be found in the result field. In case of an unsuccessful
// request, ok equals false, and the error is explained in the error field.
type Response struct {
	Description string                `json:"description,omitempty"`
	ErrorCode   int                   `json:"error_code,omitempty"`
	Ok          bool                  `json:"ok"`
	Parameters  []*ResponseParameters `json:"parameters,omitempty"`
	Result      json.RawMessage       `json:"result,omitempty"`
}

func (b *Bot) Do(method string, payload interface{}) ([]byte, error) {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegram.org")
	u.SetPath(path.Join("bot"+b.AccessToken, method))

	var buf bytes.Buffer
	if err := b.marshler.NewEncoder(&buf).Encode(payload); err != nil {
		return nil, err
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURIBytes(u.RequestURI())
	req.Header.SetContentType("application/json")
	req.SetBody(buf.Bytes())

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := b.client.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (b *Bot) Upload(method string, payload map[string]string, files ...*InputFile) ([]byte, error) {
	if len(files) == 0 {
		return b.Do(method, payload)
	}

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	for i := range files {
		_, fileName := filepath.Split(files[i].Attachment.Name())

		part, err := w.CreateFormFile(fileName, fileName)
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, files[i].Attachment); err != nil {
			return nil, err
		}
	}

	for key, val := range payload {
		if err := w.WriteField(key, val); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegram.org")
	u.SetPath(path.Join("bot"+b.AccessToken, method))

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURIBytes(u.RequestURI())
	req.Header.SetContentType(w.FormDataContentType())
	req.Header.SetMultipartFormBoundary(w.Boundary())

	if _, err := body.WriteTo(req.BodyWriter()); err != nil {
		return nil, err
	}

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := b.client.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
