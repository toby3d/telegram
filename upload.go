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

/*
Upload is a helper method which provide are three ways to send files (photos, stickers, audio,
media, etc.):

1. If the file is already stored somewhere on the Telegram servers, you don't need to reupload it:
each file object has a file_id field, simply pass this file_id as a parameter instead of uploading.
There are no limits for files sent this way.
2. Provide Telegram with an *url.URL for the file to be sent. Telegram will download and send the
file. 5 MB max size for photos and 20 MB max for other types of content.
3. Post the file using multipart/form-data in the usual way that files are uploaded via the
browser. Use []byte or io.Reader for this. 10 MB max size for photos, 50 MB for other files.

Sending by file_id

- It is not possible to change the file type when resending by file_id. I.e. a video can't be sent
as a photo, a photo can't be sent as a document, etc.
- It is not possible to resend thumbnails.
- Resending a photo by file_id will send all of its sizes.
- file_id is unique for each individual bot and can't be transferred from one bot to another.

Sending by URL

- When sending by *url.URL the target file must have the correct MIME type (e.g., audio/mpeg for
sendAudio, etc.).
- In sendDocument, sending by URL will currently only work for gif, pdf and zip files.
- To use SendVoice, the file must have the type audio/ogg and be no more than 1MB in size. 1–20MB
voice notes will be sent as files.
- Other configurations may work but we can't guarantee that they will.
*/
func (bot *Bot) Upload(method, key, name string, file InputFile, args fmt.Stringer) (*Response, error) {
	buffer := bytes.NewBuffer(nil)
	multi := multipart.NewWriter(buffer)

	requestURI := defaultURI
	requestURI.Path = fmt.Sprint("/bot", bot.AccessToken, "/", method)

	query, err := url.ParseQuery(args.String())
	if err != nil {
		return nil, err
	}

	for key, val := range query {
		if err = multi.WriteField(key, val[0]); err != nil {
			return nil, err
		}
	}

	if err = createFileField(multi, file, key, name); err != nil {
		return nil, err
	}

	if err = multi.Close(); err != nil {
		return nil, err
	}

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetBody(buffer.Bytes())
	req.Header.SetContentType(multi.FormDataContentType())
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(requestURI.String())
	req.Header.SetUserAgent(fmt.Sprint("telegram/", Version))
	req.Header.SetHost(requestURI.Hostname())

	log.Ln("Request:")
	log.D(req)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	err = http.Do(req, resp)
	log.Ln("Resp:")
	log.D(resp)
	if err != nil {
		return nil, err
	}

	var data Response
	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data, nil
}

func createFileField(w *multipart.Writer, file interface{}, key, val string) error {
	var err error
	switch src := file.(type) {
	case string: // Send FileID of file on disk path
		err = createFileFieldString(w, key, src)
	case *url.URL: // Send by URL
		err = w.WriteField(key, src.String())
	case []byte: // Upload new
		err = createFileFieldRaw(w, key, val, bytes.NewReader(src))
	case io.Reader: // Upload new
		err = createFileFieldRaw(w, key, val, src)
	default:
		return ErrBadFileType
	}
	return err
}

func createFileFieldString(w *multipart.Writer, key, src string) error {
	_, err := os.Stat(src)

	switch {
	case os.IsNotExist(err):
		err = w.WriteField(key, src)
	case os.IsExist(err):
		err = uploadFromDisk(w, key, src)
	}

	return err
}

func uploadFromDisk(w *multipart.Writer, key, src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	var formFile io.Writer
	formFile, err = w.CreateFormFile(key, file.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(formFile, file)
	return err
}

func createFileFieldRaw(w *multipart.Writer, key, value string, src io.Reader) error {
	field, err := w.CreateFormFile(key, value)
	if err != nil {
		return err
	}

	_, err = io.Copy(field, src)
	return err
}
