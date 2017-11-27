// Version of the bot API: 3.5 (November 17, 2017)
package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const (
	APIEndpoint  = "https://api.telegram.org/bot%s/%s"
	FileEndpoind = "https://api.telegram.org/file/bot%s/%s"
)

var ErrBadFileType = errors.New("bad file type")

func (bot *Bot) request(
	dst []byte,
	method string,
	args *http.Args,
) (*Response, error) {
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
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.SetBody(dst)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	if err := http.Do(req, resp); err != nil {
		return nil, err
	}

	log.Ln("Response:")
	log.D(*resp)

	var data Response
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data, nil
}

func (bot *Bot) upload(
	file InputFile,
	fieldName, fileName, method string,
	args *http.Args,
) (*Response, error) {
	var buffer bytes.Buffer
	multi := multipart.NewWriter(&buffer)
	defer multi.Close()

	switch source := file.(type) {
	case string:
		f, err := os.Open(source)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		formFile, err := multi.CreateFormFile(fieldName, f.Name())
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(formFile, f); err != nil {
			return nil, err
		}
	case []byte:
		formFile, err := multi.CreateFormFile(fieldName, fileName)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(formFile, bytes.NewReader(source)); err != nil {
			return nil, err
		}
	case *url.URL:
		if err := multi.WriteField(fieldName, source.String()); err != nil {
			return nil, err
		}
	case io.Reader:
		multi.CreateFormFile(fieldName, fileName)
	default:
		return nil, ErrBadFileType
	}

	requestURI := &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("/file/bot", bot.AccessToken, "/", method),
	}

	if args != nil {
		requestURI.RawQuery = args.String()
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("multipart/form-data; charset=utf-8")
	req.Header.SetMultipartFormBoundary(multi.Boundary())
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.SetBody(buffer.Bytes())

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	if err := http.Do(req, resp); err != nil {
		return nil, err
	}

	log.Ln("Response:")
	log.D(*resp)

	var data Response
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data, nil
}
