package telegram

type (
	// InlineQuery represents an incoming inline query. When the user sends an
	// empty query, your bot could return some default or trending results.
	InlineQuery struct {
		// Unique identifier for this query
		ID string `json:"id"`

		// Text of the query (up to 512 characters)
		Query string `json:"query"`

		// Offset of the results to be returned, can be controlled by the bot
		Offset string `json:"offset"`

		// Sender
		From *User `json:"from"`

		// Sender location, only for bots that request user location
		Location *Location `json:"location,omitempty"`
	}

	// InlineQueryResult represents one result of an inline query.
	InlineQueryResult interface {
		ResultID() string
		ResultType() string
		ResultReplyMarkup() *InlineKeyboardMarkup
	}

	// InlineQueryResultArticle represents a link to an article or web page.
	InlineQueryResultArticle struct {
		// Type of the result, must be article
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Title of the result
		Title string `json:"title"`

		// URL of the result
		URL string `json:"url,omitempty"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url,omitempty"`

		// Content of the message to be sent
		InputMessageContent interface{} `json:"input_message_content"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Pass True, if you don't want the URL to be shown in the message
		HideURL bool `json:"hide_url,omitempty"`

		// Thumbnail width
		ThumbWidth int `json:"thumb_width,omitempty"`

		// Thumbnail height
		ThumbHeight int `json:"thumb_height,omitempty"`
	}

	// InlineQueryResultPhoto represents a link to a photo. By default, this
	// photo will be sent by the user with optional caption. Alternatively, you
	// can use input_message_content to send a message with the specified content
	// instead of the photo.
	InlineQueryResultPhoto struct {
		// Type of the result, must be photo
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL of the photo. Photo must be in jpeg format. Photo size
		// must not exceed 5MB
		PhotoURL string `json:"photo_url"`

		// URL of the thumbnail for the photo
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Caption of the photo to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Width of the photo
		PhotoWidth int `json:"photo_width,omitempty"`

		// Height of the photo
		PhotoHeight int `json:"photo_height,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the photo
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultGif represents a link to an animated GIF file. By
	// default, this animated GIF file will be sent by the user with optional
	// caption. Alternatively, you can use input_message_content to send a
	// message with the specified content instead of the animation.
	InlineQueryResultGif struct {
		// Type of the result, must be gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the GIF file. File size must not exceed 1MB
		GifURL string `json:"gif_url"`

		// URL of the static thumbnail for the result (jpeg or gif)
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Caption of the GIF file to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Width of the GIF
		GifWidth int `json:"gif_width,omitempty"`

		// Height of the GIF
		GifHeight int `json:"gif_height,omitempty"`

		// Duration of the GIF
		GifDuration int `json:"gif_duration,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the GIF animation
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultMpeg4Gif represents a link to a video animation
	// (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4
	// file will be sent by the user with optional caption. Alternatively, you
	// can use input_message_content to send a message with the specified content
	// instead of the animation.
	InlineQueryResultMpeg4Gif struct {
		// Type of the result, must be mpeg4_gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the MP4 file. File size must not exceed 1MB
		Mpeg4URL string `json:"mpeg4_url"`

		// URL of the static thumbnail (jpeg or gif) for the result
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Caption of the MPEG-4 file to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Video width
		Mpeg4Width int `json:"mpeg4_width,omitempty"`

		// Video height
		Mpeg4Height int `json:"mpeg4_height,omitempty"`

		// Video duration
		Mpeg4Duration int `json:"mpeg4_duration,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the video animation
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultVideo represents a link to a page containing an embedded
	// video player or a video file. By default, this video file will be sent by
	// the user with an optional caption. Alternatively, you can use
	// input_message_content to send a message with the specified content
	// instead of the video.
	//
	// If an InlineQueryResultVideo message contains an embedded video (e.g.,
	// YouTube), you must replace its content using input_message_content.
	InlineQueryResultVideo struct {
		// Type of the result, must be video
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the embedded video player or video file
		VideoURL string `json:"video_url"`

		// Mime type of the content of video url, "text/html" or "video/mp4"
		MimeType string `json:"mime_type"`

		// URL of the thumbnail (jpeg only) for the video
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title"`

		// Caption of the video to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Video width
		VideoWidth int `json:"video_width,omitempty"`

		// Video height
		VideoHeight int `json:"video_height,omitempty"`

		// Video duration in seconds
		VideoDuration int `json:"video_duration,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the video. This field is
		// required if InlineQueryResultVideo is used to send an HTML-page as a
		// result (e.g., a YouTube video).
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultAudio represents a link to an mp3 audio file. By default,
	// this audio file will be sent by the user. Alternatively, you can use
	// input_message_content to send a message with the specified content
	// instead of the audio.
	InlineQueryResultAudio struct {
		// Type of the result, must be audio
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the audio file
		AudioURL string `json:"audio_url"`

		// Title
		Title string `json:"title"`

		// Caption, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Performer
		Performer string `json:"performer,omitempty"`

		// Audio duration in seconds
		AudioDuration int `json:"audio_duration,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the audio
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultVoice represents a link to a voice recording in an .ogg
	// container encoded with OPUS. By default, this voice recording will be
	// sent by the user. Alternatively, you can use input_message_content to
	// send a message with the specified content instead of the the voice message.
	InlineQueryResultVoice struct {
		// Type of the result, must be voice
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the voice recording
		VoiceURL string `json:"voice_url"`

		// Recording title
		Title string `json:"title"`

		// Caption, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Recording duration in seconds
		VoiceDuration int `json:"voice_duration,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the voice recording
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultDocument represents a link to a file. By default, this
	// file will be sent by the user with an optional caption. Alternatively,
	// you can use input_message_content to send a message with the specified
	// content instead of the file. Currently, only .PDF and .ZIP files can be
	// sent using this method.
	InlineQueryResultDocument struct {
		// Type of the result, must be document
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// Title for the result
		Title string `json:"title"`

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// A valid URL for the file
		DocumentURL string `json:"document_url"`

		// Mime type of the content of the file, either "application/pdf" or
		// "application/zip"
		MimeType string `json:"mime_type"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// URL of the thumbnail (jpeg only) for the file
		ThumbURL string `json:"thumb_url,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the file
		InputMessageContent interface{} `json:"input_message_content,omitempty"`

		// Thumbnail width
		ThumbWidth int `json:"thumb_width,omitempty"`

		// Thumbnail height
		ThumbHeight int `json:"thumb_height,omitempty"`
	}

	// InlineQueryResultLocation represents a location on a map. By default, the
	// location will be sent by the user. Alternatively, you can use
	// input_message_content to send a message with the specified content
	// instead of the location.
	InlineQueryResultLocation struct {
		// Type of the result, must be location
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Location title
		Title string `json:"title"`

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url,omitempty"`

		// Location latitude in degrees
		Latitude float32 `json:"latitude"`

		// Location longitude in degrees
		Longitude float32 `json:"longitude"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the location
		InputMessageContent interface{} `json:"input_message_content,omitempty"`

		// Thumbnail width
		ThumbWidth int `json:"thumb_width,omitempty"`

		// Thumbnail height
		ThumbHeight int `json:"thumb_height,omitempty"`
	}

	// InlineQueryResultVenue represents a venue. By default, the venue will be
	// sent by the user. Alternatively, you can use input_message_content to
	// send a message with the specified content instead of the venue.
	InlineQueryResultVenue struct {
		// Type of the result, must be venue
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Title of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue if known
		FoursquareID string `json:"foursquare_id,omitempty"`

		// Foursquare type of the venue, if known. (For example,
		// "arts_entertainment/default", "arts_entertainment/aquarium" or
		// "food/icecream".)
		FoursquareType string `json:"foursquare_type,omitempty"`

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url,omitempty"`

		// Latitude of the venue location in degrees
		Latitude float32 `json:"latitude"`

		// Longitude of the venue location in degrees
		Longitude float32 `json:"longitude"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the venue
		InputMessageContent interface{} `json:"input_message_content,omitempty"`

		// Thumbnail width
		ThumbWidth int `json:"thumb_width,omitempty"`

		// Thumbnail height
		ThumbHeight int `json:"thumb_height,omitempty"`
	}

	// InlineQueryResultContact represents a contact with a phone number. By
	// default, this contact will be sent by the user. Alternatively, you can
	// use input_message_content to send a message with the specified content
	// instead of the contact.
	InlineQueryResultContact struct {
		// Type of the result, must be contact
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name,omitempty"`

		// Additional data about the contact in the form of a vCard, 0-2048 bytes
		VCard string `json:"vcard,omitempty"`

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the contact
		InputMessageContent interface{} `json:"input_message_content,omitempty"`

		// Thumbnail width
		ThumbWidth int `json:"thumb_width,omitempty"`

		// Thumbnail height
		ThumbHeight int `json:"thumb_height,omitempty"`
	}

	// InlineQueryResultGame represents a Game.
	InlineQueryResultGame struct {
		// Type of the result, must be game
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// Short name of the game
		GameShortName string `json:"game_short_name"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// InlineQueryResultCachedPhoto represents a link to a photo stored on the
	// Telegram servers. By default, this photo will be sent by the user with an
	// optional caption. Alternatively, you can use input_message_content to
	// send a message with the specified content instead of the photo.
	InlineQueryResultCachedPhoto struct {
		// Type of the result, must be photo
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier of the photo
		PhotoFileID string `json:"photo_file_id"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Caption of the photo to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the photo
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedGif represents a link to an animated GIF file
	// stored on the Telegram servers. By default, this animated GIF file will
	// be sent by the user with an optional caption. Alternatively, you can use
	// input_message_content to send a message with specified content instead of
	// the animation.
	InlineQueryResultCachedGif struct {
		// Type of the result, must be gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the GIF file
		GifFileID string `json:"gif_file_id"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Caption of the GIF file to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the GIF animation
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedMpeg4Gif represents a link to a video animation
	// (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By
	// default, this animated MPEG-4 file will be sent by the user with an
	// optional caption. Alternatively, you can use input_message_content to
	// send a message with the specified content instead of the animation.
	InlineQueryResultCachedMpeg4Gif struct {
		// Type of the result, must be mpeg4_gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the MP4 file
		Mpeg4FileID string `json:"mpeg4_file_id"`

		// Title for the result
		Title string `json:"title,omitempty"`

		// Caption of the MPEG-4 file to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the video animation
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedSticker represents a link to a sticker stored on
	// the Telegram servers. By default, this sticker will be sent by the user.
	// Alternatively, you can use input_message_content to send a message with
	// the specified content instead of the sticker.
	InlineQueryResultCachedSticker struct {
		// Type of the result, must be sticker
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier of the sticker
		StickerFileID string `json:"sticker_file_id"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the sticker
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedDocument represents a link to a file stored on the
	// Telegram servers. By default, this file will be sent by the user with an
	// optional caption. Alternatively, you can use input_message_content to
	// send a message with the specified content instead of the file.
	InlineQueryResultCachedDocument struct {
		// Type of the result, must be document
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// Title for the result
		Title string `json:"title"`

		// A valid file identifier for the file
		DocumentFileID string `json:"document_file_id"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the file
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedVideo represents a link to a video file stored on
	// the Telegram servers. By default, this video file will be sent by the
	// user with an optional caption. Alternatively, you can use
	// input_message_content to send a message with the specified content
	// instead of the video.
	InlineQueryResultCachedVideo struct {
		// Type of the result, must be video
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the video file
		VideoFileID string `json:"video_file_id"`

		// Title for the result
		Title string `json:"title"`

		// Short description of the result
		Description string `json:"description,omitempty"`

		// Caption of the video to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the video
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedVoice represents a link to a voice message stored
	// on the Telegram servers. By default, this voice message will be sent by
	// the user. Alternatively, you can use input_message_content to send a
	// message with the specified content instead of the voice message.
	InlineQueryResultCachedVoice struct {
		// Type of the result, must be voice
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the voice message
		VoiceFileID string `json:"voice_file_id"`

		// Voice message title
		Title string `json:"title"`

		// Caption, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the voice message
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InlineQueryResultCachedAudio represents a link to an mp3 audio file
	// stored on the Telegram servers. By default, this audio file will be sent
	// by the user. Alternatively, you can use input_message_content to send a
	// message with the specified content instead of the audio.
	InlineQueryResultCachedAudio struct {
		// Type of the result, must be audio
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the audio file
		AudioFileID string `json:"audio_file_id"`

		// Caption, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

		// Content of the message to be sent instead of the audio
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	}

	// InputMessageContent represents the content of a message to be sent as a result
	// of an inline query.
	InputMessageContent interface {
		IsInputMessageContent() bool
	}

	// InputTextMessageContent represents the content of a text message to be
	// sent as the result of an inline query.
	InputTextMessageContent struct {
		// Text of the message to be sent, 1-4096 characters
		MessageText string `json:"message_text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode,omitempty"`

		// Disables link previews for links in the sent message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	}

	// InputLocationMessageContent represents the content of a location message
	// to be sent as the result of an inline query.
	InputLocationMessageContent struct {
		// Latitude of the location in degrees
		Latitude float32 `json:"latitude"`

		// Longitude of the location in degrees
		Longitude float32 `json:"longitude"`
	}

	// InputVenueMessageContent represents the content of a venue message to be
	// sent as the result of an inline query.
	InputVenueMessageContent struct {
		// Latitude of the venue in degrees
		Latitude float32 `json:"latitude"`

		// Longitude of the venue in degrees
		Longitude float32 `json:"longitude"`

		// Name of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue, if known
		FoursquareID string `json:"foursquare_id,omitempty"`

		// Foursquare type of the venue, if known. (For example,
		// "arts_entertainment/default", "arts_entertainment/aquarium" or
		// "food/icecream".)
		FoursquareType string `json:"foursquare_type,omitempty"`
	}

	// InputContactMessageContent represents the content of a contact message to
	// be sent as the result of an inline query.
	InputContactMessageContent struct {
		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name,omitempty"`

		// Additional data about the contact in the form of a vCard, 0-2048 bytes
		VCard string `json:"vcard,omitempty"`
	}

	// ChosenInlineResult represents a result of an inline query that was chosen
	// by the user and sent to their chat partner.
	ChosenInlineResult struct {
		// The unique identifier for the result that was chosen
		ResultID string `json:"result_id"`

		// Identifier of the sent inline message. Available only if there is an
		// inline keyboard attached to the message. Will be also received in
		// callback queries and can be used to edit the message.
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// The query that was used to obtain the result
		Query string `json:"query"`

		// The user that chose the result
		From *User `json:"from"`

		// Sender location, only for bots that require user location
		Location *Location `json:"location,omitempty"`
	}
)
