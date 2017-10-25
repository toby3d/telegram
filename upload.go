package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// There are three ways to send files (photos, stickers, audio, media, etc.):
//
// 1. If the file is already stored somewhere on the Telegram servers, you don't need to reupload it: each file object has a file_id field, simply pass this file_id as a parameter instead of uploading. There are no limits for files sent this way.
// 2. Provide Telegram with an HTTP URL for the file to be sent. Telegram will download and send the file. 5 MB max size for photos and 20 MB max for other types of content.
// 3. Post the file using multipart/form-data in the usual way that files are uploaded via the browser. 10 MB max size for photos, 50 MB for other files.
func (bot *Bot) upload(file InputFile, fieldName, fileName, method string, args *http.Args) (*Response, error) {
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
		return nil, errors.New("bad file type")
	}

	requestURI := fmt.Sprintf(APIEndpoint, bot.AccessToken, method)
	if args != nil {
		requestURI += fmt.Sprint("?", args.String())
	}

	var req http.Request
	var resp http.Response

	req.Header.SetMethod("POST")
	req.Header.SetContentType("multipart/form-data")
	req.Header.SetMultipartFormBoundary(multi.Boundary())
	args.WriteTo(req.BodyWriter())
	req.SetRequestURI(requestURI)
	req.SetBody(buffer.Bytes())

	if err := http.Do(&req, &resp); err != nil {
		return nil, err
	}

	var data Response
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data, nil
}
