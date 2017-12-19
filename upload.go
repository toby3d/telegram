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

var ErrBadFileType = errors.New("bad file type")

// upload is a helper method which provide are three ways to send files
// (photos, stickers, audio, media, etc.):
//
// 1. If the file is already stored somewhere on the Telegram servers, you don't
// need to reupload it: each file object has a file_id field, simply pass this
// file_id as a parameter instead of uploading. There are no limits for files
// sent this way.
// 2. Provide Telegram with an *url.URL for the file to be sent. Telegram will
// download and send the file. 5 MB max size for photos and 20 MB max for other
// types of content.
// 3. Post the file using multipart/form-data in the usual way that files are
// uploaded via the browser. Use path string, []byte or io.Reader for this. 10 MB
// max size for photos, 50 MB for other files.
//
// Sending by file_id
//
// - It is not possible to change the file type when resending by file_id. I.e.
// a video can't be sent as a photo, a photo can't be sent as a document, etc.
// - It is not possible to resend thumbnails.
// - Resending a photo by file_id will send all of its sizes.
// - file_id is unique for each individual bot and can't be transferred from one
// bot to another.
//
// Sending by URL
//
// - When sending by *url.URL the target file must have the correct MIME type
// (e.g., audio/mpeg for sendAudio, etc.).
// - In sendDocument, sending by URL will currently only work for gif, pdf and
// zip files.
// - To use SendVoice, the file must have the type audio/ogg and be no more than
// 1MB in size. 1–20MB voice notes will be sent as files.
// - Other configurations may work but we can't guarantee that they will.
func (bot *Bot) upload(file InputFile, fieldName, fileName, method string, args *http.Args) (*Response, error) {
	buffer := bytes.NewBuffer(nil)
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
	case string: // Send by 'file_id'
		err := multi.WriteField(fieldName, f)
		if err != nil {
			return nil, err
		}

		/*
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
		*/
	case []byte: // Upload new
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
	// req.Header.SetHost("api.telegram.org")

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
