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
	req.Header.SetMethod("POST")
	req.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.Header.SetHost("api.telegram.org")
	req.Header.SetContentType("application/json; charset=utf-8")
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

func (bot *Bot) upload(file InputFile, fieldName, fileName, method string, args *http.Args) (*Response, error) {
	buffer := &bytes.Buffer{}
	multi := multipart.NewWriter(buffer)

	requestURI := &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("/bot", bot.AccessToken, "/", method),
	}

	query, err := url.ParseQuery(args.String())
	if err != nil {
		return nil, err
	}

	for key, val := range query {
		if err := multi.WriteField(key, val[0]); err != nil {
			return nil, err
		}
	}

	switch f := file.(type) {
	case string:
		src, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer src.Close()

		formFile, err := multi.CreateFormFile(fieldName, src.Name())
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(formFile, src); err != nil {
			return nil, err
		}
	case []byte:
		formFile, err := multi.CreateFormFile(fieldName, fileName)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(formFile, bytes.NewReader(f)); err != nil {
			return nil, err
		}
	case *url.URL:
		if err := multi.WriteField(fieldName, f.String()); err != nil {
			return nil, err
		}
	case io.Reader:
		if _, err := multi.CreateFormFile(fieldName, fileName); err != nil {
			return nil, err
		}
	default:
		return nil, ErrBadFileType
	}
	if err := multi.Close(); err != nil {
		return nil, err
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetBody(buffer.Bytes())
	req.Header.SetContentType(multi.FormDataContentType())
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent("go-telegram/3.5")
	req.Header.SetHost("api.telegram.org")

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
		return nil, errors.New(data.Description)
	}

	return &data, nil
}
