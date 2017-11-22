package telegram

// UploadFile is a helper method which provide are three ways to send files
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
func (bot *Bot) UploadFile(name string, file InputFile) (*Response, error) {
	return bot.upload(file, "file", name, "", nil)
}
