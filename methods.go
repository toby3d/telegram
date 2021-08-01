package telegram

import (
	"strconv"
	"strings"
)

type (
	// SendMessage represents data for SendMessage method.
	SendMessage struct {
		ChatID ChatID `json:"chat_id"`

		// Text of the message to be sent
		Text string `json:"text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		Entities []*MessageEntity `json:"entities,omitempty"`

		// Disables link previews for links in this message
		DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// ForwardMessage represents data for ForwardMessage method.
	ForwardMessage struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
		FromChatID ChatID `json:"from_chat_id"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// Message identifier in the chat specified in from_chat_id
		MessageID int64 `json:"message_id"`
	}

	// CopyMessage represents data for CopyMessage method.
	CopyMessage struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier for the chat where the original message was sent
		FromChatID ChatID `json:"from_chat_id"`

		// Message identifier in the chat specified in from_chat_id
		MessageID int64 `json:"message_id"`

		// New caption for media, 0-1024 characters after entities parsing. If not specified, the original
		// caption is kept
		Caption string `json:"caption,omitempty"`

		// Mode for parsing entities in the new caption. See formatting options for more details.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the new caption, which can be specified instead of
		// parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
		// keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendPhoto represents data for SendPhoto method.
	SendPhoto struct {
		ChatID ChatID `json:"chat_id"`

		// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
		Photo *InputFile `json:"photo"`

		// Photo caption (may also be used when resending photos by file_id), 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendAudio represents data for SendAudio method.
	SendAudio struct {
		ChatID ChatID `json:"chat_id"`

		// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data.
		Audio *InputFile `json:"audio"`

		// Audio caption, 0-1024 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Duration of the audio in seconds
		Duration int `json:"duration,omitempty"`

		// Performer
		Performer string `json:"performer,omitempty"`

		// Track name
		Title string `json:"title,omitempty"`

		// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendDocument represents data for SendDocument method.
	SendDocument struct {
		ChatID ChatID `json:"chat_id"`

		// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
		Document *InputFile `json:"document"`

		// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
		// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
		// width and height should not exceed 320. Ignored if the file is not uploaded using
		// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can
		// pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under
		// <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Document caption (may also be used when resending documents by file_id), 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Disables automatic server-side content type detection for files uploaded using multipart/form-data
		DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendDocument represents data for SendVideo method.
	SendVideo struct {
		ChatID ChatID `json:"chat_id"`

		// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
		Video *InputFile `json:"video"`

		// Duration of sent video in seconds
		Duration int `json:"duration,omitempty"`

		// Video width
		Width int `json:"width,omitempty"`

		// Video height
		Height int `json:"height,omitempty"`

		// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Video caption (may also be used when resending videos by file_id), 0-1024 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Pass True, if the uploaded video is suitable for streaming
		SupportsStreaming bool `json:"supports_streaming,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendAnimation represents data for SendAnimation method.
	SendAnimation struct {
		ChatID ChatID `json:"chat_id"`

		// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data.
		Animation *InputFile `json:"animation"`

		// Duration of sent animation in seconds
		Duration int `json:"duration,omitempty"`

		// Animation width
		Width int `json:"width,omitempty"`

		// Animation height
		Height int `json:"height,omitempty"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Animation caption (may also be used when resending animation by file_id), 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendVoice represents data for SendVoice method.
	SendVoice struct {
		ChatID ChatID `json:"chat_id"`

		// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
		Voice *InputFile `json:"voice"`

		// Voice message caption, 0-1024 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Duration of the voice message in seconds
		Duration int `json:"duration,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendVideoNote represents data for SendVideoNote method.
	SendVideoNote struct {
		ChatID ChatID `json:"chat_id"`

		// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data.. Sending video notes by a URL is currently unsupported
		VideoNote *InputFile `json:"video_note"`

		// Duration of sent video in seconds
		Duration int `json:"duration,omitempty"`

		// Video width and height, i.e. diameter of the video message
		Length int `json:"length,omitempty"`

		// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendMediaGroup represents data for SendMediaGroup method.
	SendMediaGroup struct {
		ChatID ChatID `json:"chat_id" form:"chat_id"`

		// A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
		Media []AlbumMedia `json:"media" form:"media"`

		// Sends the messages silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty" form:"disable_notification"`

		// If the messages are a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty" form:"reply_to_message_id"`
	}

	// SendLocation represents data for SendLocation method.
	SendLocation struct {
		ChatID ChatID `json:"chat_id"`

		// Latitude of the location
		Latitude float64 `json:"latitude"`

		// Longitude of the location
		Longitude float64 `json:"longitude"`

		// The radius of uncertainty for the location, measured in meters; 0-1500
		HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

		// Period in seconds for which the location will be updated (see Live Locations), should be between 60
		// and 86400.
		LivePeriod int `json:"live_period,omitempty"`

		// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360
		// if specified.
		Heading int `json:"heading,omitempty"`

		// For live locations, a maximum distance for proximity alerts about approaching another chat member,
		// in meters. Must be between 1 and 100000 if specified.
		ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be
		// shown. If not empty, the first button must be a Pay button.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// EditMessageLiveLocation represents data for EditMessageLiveLocation method.
	EditMessageLiveLocation struct {
		ChatID ChatID `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the sent message
		MessageID int64 `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// Latitude of new location
		Latitude float64 `json:"latitude"`

		// Longitude of new location
		Longitude float64 `json:"longitude"`

		// The radius of uncertainty for the location, measured in meters; 0-1500
		HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

		// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
		Heading int `json:"heading,omitempty"`

		// Maximum distance for proximity alerts about approaching another chat member, in meters. Must be
		// between 1 and 100000 if specified.
		ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

		// A JSON-serialized object for a new inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// StopMessageLiveLocation represents data for StopMessageLiveLocation method.
	StopMessageLiveLocation struct {
		ChatID ChatID `json:"chat_id,omitempty"`

		// Required if inline_message_id is not specified. Identifier of the message with live location to stop
		MessageID int64 `json:"message_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// A JSON-serialized object for a new inline keyboard.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// SendVenue represents data for SendVenue method.
	SendVenue struct {
		ChatID ChatID `json:"chat_id"`

		// Latitude of the venue
		Latitude float64 `json:"latitude"`

		// Longitude of the venue
		Longitude float64 `json:"longitude"`

		// Name of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue
		FoursquareID string `json:"foursquare_id,omitempty"`

		// Foursquare type of the venue, if known. (For example, "arts_entertainment/default",
		// "arts_entertainment/aquarium" or "food/icecream".)
		FoursquareType string `json:"foursquare_type,omitempty"`

		// Google Places identifier of the venue
		GooglePlaceID string `json:"google_place_id,omitempty"`

		// Google Places type of the venue.
		GooglePlaceType string `json:"google_place_type,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendContact represents data for SendContact method.
	SendContact struct {
		ChatID ChatID `json:"chat_id"`

		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name"`

		// Additional data about the contact in the form of a vCard, 0-2048 bytes
		VCard string `json:"vcard,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendPoll represents data for SendPoll method.
	SendPoll struct {
		ChatID ChatID `json:"chat_id"`

		// Poll question, 1-300 characters
		Question string `json:"question"`

		// List of answer options, 2-10 strings 1-100 characters each
		Options []string `json:"options"`

		// True, if the poll needs to be anonymous, defaults to True
		IsAnonymous bool `json:"is_anonymous,omitempty"`

		// Poll type, “quiz” or “regular”, defaults to “regular”
		Type string `json:"type,omitempty"`

		// True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
		AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

		// 0-based identifier of the correct answer option, required for polls in quiz mode
		CorrectOptionID int64 `json:"correct_option_id,omitempty"`

		// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style
		// poll, 0-200 characters with at most 2 line feeds after entities parsing
		Explanation string `json:"explanation,omitempty"`

		// Mode for parsing entities in the explanation. See formatting options for more details.
		ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

		// List of special entities that appear in the caption, which can be specified instead of parse_mode
		ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`

		// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
		OpenPeriod int `json:"open_period,omitempty"`

		// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
		CloseDate int64 `json:"close_date,omitempty"`

		// Pass True, if the poll needs to be immediately closed
		IsClosed bool `json:"is_closed,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendDice represents data for SendDice method.
	SendDice struct {
		ChatID ChatID `json:"chat_id"`

		// Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”,
		// “⚽”, or “🎰”. Dice can have values 1-6 for “🎲” and “🎯”, values 1-5 for “🏀” and “⚽”, and values
		// 1-64 for “🎰”. Defaults to “🎲”
		Emoji string `json:"emoji,omitempty"`

		// Sends the message silently. Users will receive a notification with no sound.
		DisableNotification bool `json:"disable_notification,omitempty"`

		// If the message is a reply, ID of the original message
		ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

		// Pass True, if the message should be sent even if the specified replied-to message is not found
		AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

		// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
		ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	}

	// SendChatAction represents data for SendChatAction method.
	SendChatAction struct {
		ChatID ChatID `json:"chat_id"`

		// Type of action to broadcast
		Action string `json:"action"`
	}

	// GetUserProfilePhotos represents data for GetUserProfilePhotos method.
	GetUserProfilePhotos struct {
		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// Sequential number of the first photo to be returned. By default, all photos are returned.
		Offset int `json:"offset,omitempty"`

		// Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`
	}

	// GetFile represents data for GetFile method.
	GetFile struct {
		// File identifier to get info about
		FileID string `json:"file_id"`
	}

	// BanChatMember represents data for BanChatMember method.
	BanChatMember struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
		UntilDate int64 `json:"until_date"`

		// Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
		RevokeMessages bool `json:"revoke_messages,omitempty"`
	}

	// UnbanChatMember represents data for UnbanChatMember method.
	UnbanChatMember struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// Do nothing if the user is not banned
		OnlyIfBanned bool `json:"only_if_banned,omitempty"`
	}

	// RestrictChatMember represents data for RestrictChatMember method.
	RestrictChatMember struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// New user permissions
		Permissions *ChatPermissions `json:"permissions"`

		// Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
		UntilDate int64 `json:"until_date,omitempty"`
	}

	// PromoteChatMember represents data for PromoteChatMember method.
	PromoteChatMember struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// Pass True, if the administrator's presence in the chat is hidden
		IsAnonymous bool `json:"is_anonymous,omitempty"`

		// Pass True, if the administrator can access the chat event log, chat statistics, message statistics
		// in channels, see channel members, see anonymous administrators in supergoups and ignore slow mode.
		// Implied by any other administrator privilege
		CanManageChat bool `json:"can_manage_chat,omitempty"`

		// Pass True, if the administrator can change chat title, photo and other settings
		CanChangeInfo bool `json:"can_change_info,omitempty"`

		// Pass True, if the administrator can create channel posts, channels only
		CanPostMessages bool `json:"can_post_messages,omitempty"`

		// Pass True, if the administrator can edit messages of other users and can pin messages, channels only
		CanEditMessages bool `json:"can_edit_messages,omitempty"`

		// Pass True, if the administrator can delete messages of other users
		CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

		// Pass True, if the administrator can manage voice chats, supergroups only
		CanManageVoiceChats bool `json:"can_manage_voice_chats,omitempty"`

		// Pass True, if the administrator can invite new users to the chat
		CanInviteUsers bool `json:"can_invite_users,omitempty"`

		// Pass True, if the administrator can restrict, ban or unban chat members
		CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

		// Pass True, if the administrator can pin messages, supergroups only
		CanPinMessages bool `json:"can_pin_messages,omitempty"`

		// Pass True, if the administrator can add new administrators with a subset of his own privileges or
		// demote administrators that he has promoted, directly or indirectly (promoted by administrators that
		// were appointed by him)
		CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	}

	// SetChatAdministratorCustomTitle represents data for SetChatAdministratorCustomTitle method.
	SetChatAdministratorCustomTitle struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`

		// New custom title for the administrator; 0-16 characters, emoji are not allowed
		CustomTitle string `json:"custom_title"`
	}

	// SetChatPermissions represents data for SetChatPermissions method.
	SetChatPermissions struct {
		ChatID ChatID `json:"chat_id"`

		// New default chat permissions
		Permissions ChatPermissions `json:"permissions"`
	}

	// ExportChatInviteLink represents data for ExportChatInviteLink method.
	ExportChatInviteLink struct {
		ChatID ChatID `json:"chat_id"`
	}

	// CreateChatInviteLink represents data for CreateChatInviteLink method.
	CreateChatInviteLink struct {
		ChatID ChatID `json:"chat_id"`

		// Point in time (Unix timestamp) when the link will expire
		ExpireDate int64 `json:"expire_date,omitempty"`

		// Maximum number of users that can be members of the chat simultaneously after joining the chat via
		// this invite link; 1-99999
		MemberLimit int `json:"member_limit,omitempty"`
	}

	// EditChatInviteLink represents data for EditChatInviteLink method.
	EditChatInviteLink struct {
		ChatID ChatID `json:"chat_id"`

		// The invite link to edit
		InviteLink string `json:"invite_link"`

		// Point in time (Unix timestamp) when the link will expire
		ExpireDate int64 `json:"expire_date,omitempty"`

		// Maximum number of users that can be members of the chat simultaneously after joining the chat via
		// this invite link; 1-99999
		MemberLimit int `json:"member_limit,omitempty"`
	}

	// RevokeChatInviteLink represents data for RevokeChatInviteLink method.
	RevokeChatInviteLink struct {
		ChatID ChatID `json:"chat_id"`

		// The invite link to revoke
		InviteLink string `json:"invite_link"`
	}

	// SetChatPhoto represents data for SetChatPhoto method.
	SetChatPhoto struct {
		ChatID ChatID `json:"chat_id"`

		// New chat photo, uploaded using multipart/form-data
		ChatPhoto InputFile `json:"chat_photo"`
	}

	// DeleteChatPhoto represents data for DeleteChatPhoto method.
	DeleteChatPhoto struct {
		ChatID ChatID `json:"chat_id"`
	}

	// SetChatTitle represents data for SetChatTitle method.
	SetChatTitle struct {
		ChatID ChatID `json:"chat_id"`

		// New chat title, 1-255 characters
		Title string `json:"title"`
	}

	// SetChatDescription represents data for SetChatDescription method.
	SetChatDescription struct {
		ChatID ChatID `json:"chat_id"`

		// New chat description, 0-255 characters
		Description string `json:"description"`
	}

	// PinChatMessage represents data for PinChatMessage method.
	PinChatMessage struct {
		ChatID ChatID `json:"chat_id"`

		// Identifier of a message to pin
		MessageID int64 `json:"message_id"`

		// Pass true, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
		DisableNotification bool `json:"disable_notification"`
	}

	// UnpinChatMessage represents data for UnpinChatMessage method.
	UnpinChatMessage struct {
		ChatID ChatID `json:"chat_id"`

		// Identifier of a message to unpin. If not specified, the most recent pinned message (by sending
		// date) will be unpinned.
		MessageID int64 `json:"messge_id,omitempty"`
	}

	// UnpinAllChatMessages represents data for UnpinAllChatMessages method.
	UnpinAllChatMessages struct {
		ChatID ChatID `json:"chat_id"`
	}

	// LeaveChat represents data for LeaveChat method.
	LeaveChat struct {
		ChatID ChatID `json:"chat_id"`
	}

	// GetChat represents data for GetChat method.
	GetChat struct {
		ChatID ChatID `json:"chat_id"`
	}

	// GetChatAdministrators represents data for GetChatAdministrators method.
	GetChatAdministrators struct {
		ChatID ChatID `json:"chat_id"`
	}

	// GetChatMemberCount represents data for GetChatMemberCount method.
	GetChatMemberCount struct {
		ChatID ChatID `json:"chat_id"`
	}

	// GetChatMember represents data for GetChatMember method.
	GetChatMember struct {
		ChatID ChatID `json:"chat_id"`

		// Unique identifier of the target user
		UserID int64 `json:"user_id"`
	}

	// SetChatStickerSet represents data for SetChatStickerSet method.
	SetChatStickerSet struct {
		ChatID ChatID `json:"chat_id"`

		// Name of the sticker set to be set as the group sticker set
		StickerSetName string `json:"sticker_set_name"`
	}

	// DeleteChatStickerSet represents data for DeleteChatStickerSet method.
	DeleteChatStickerSet struct {
		ChatID ChatID `json:"chat_id"`
	}

	// AnswerCallbackQuery represents data for AnswerCallbackQuery method.
	AnswerCallbackQuery struct {
		// Unique identifier for the query to be answered
		CallbackQueryID string `json:"callback_query_id"`

		// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
		Text string `json:"text,omitempty"`

		// URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.
		//
		// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
		URL string `json:"url,omitempty"`

		// If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
		ShowAlert bool `json:"show_alert,omitempty"`

		// The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
		CacheTime int `json:"cache_time,omitempty"`
	}

	// SetMyCommands represents data for SetMyCommands method.
	SetMyCommands struct {
		// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100
		// commands can be specified.
		Commands []*BotCommand `json:"commands"`

		// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to
		// BotCommandScopeDefault.
		Scope BotCommandScope `json:"scope,omitempty"`

		// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given
		// scope, for whose language there are no dedicated commands
		LanguageCode string `json:"language_code,omitempty"`
	}

	// DeleteMyCommands represents data for DeleteMyCommands method.
	DeleteMyCommands struct {
		// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to
		// BotCommandScopeDefault.
		Scope BotCommandScope `json:"scope,omitempty"`

		// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given
		// scope, for whose language there are no dedicated commands
		LanguageCode string `json:"language_code,omitempty"`
	}

	// GetMyCommands represents data for GetMyCommands method.
	GetMyCommands struct {
		// A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
		Scope BotCommandScope `json:"scope,omitempty"`

		// A two-letter ISO 639-1 language code or an empty string
		LanguageCode string `json:"language_code,omitempty"`
	}
)

// GetMe testing your bot's auth token. Returns basic information about the bot in form of a User object.
func (b Bot) GetMe() (*User, error) {
	src, err := b.Do(MethodGetMe, nil)
	if err != nil {
		return nil, err
	}

	result := new(User)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// LogOut method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot
// before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful
// call, you will not be able to log in again using the same token for 10 minutes. Returns True on success. Requires
// no parameters.
func (b Bot) LogOut() (ok bool, err error) {
	src, err := b.Do(MethodLogOut, nil)
	if err != nil {
		return false, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// Close method to close the bot instance before moving it from one local server to another. You need to delete the
// webhook before calling this method to ensure that the bot isn't launched again after server restart. The method
// will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no
// parameters.
func (b Bot) Close() (ok bool, err error) {
	src, err := b.Do(MethodClose, nil)
	if err != nil {
		return false, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

func NewMessage(chatID ChatID, text string) SendMessage {
	return SendMessage{
		ChatID: chatID,
		Text:   text,
	}
}

// SendMessage send text messages. On success, the sent Message is returned.
func (b Bot) SendMessage(p SendMessage) (*Message, error) {
	src, err := b.Do(MethodSendMessage, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewForward(fromChatID, toChatID ChatID, messageID int64) ForwardMessage {
	return ForwardMessage{
		FromChatID: fromChatID,
		ChatID:     toChatID,
		MessageID:  messageID,
	}
}

// ForwardMessage forward messages of any kind. On success, the sent Message is returned.
func (b Bot) ForwardMessage(p ForwardMessage) (*Message, error) {
	src, err := b.Do(MethodForwardMessage, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, err
}

// CopyMessage copy messages of any kind. The method is analogous to the method forwardMessages, but the copied
// message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (b Bot) CopyMessage(p CopyMessage) (*MessageID, error) {
	src, err := b.Do(MethodCopyMessage, p)
	if err != nil {
		return nil, err
	}

	result := new(MessageID)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, err
}

func NewPhoto(chatID ChatID, photo *InputFile) SendPhoto {
	return SendPhoto{
		ChatID: chatID,
		Photo:  photo,
	}
}

// SendPhoto send photos. On success, the sent Message is returned.
func (b Bot) SendPhoto(p SendPhoto) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["allow_sending_without_reply"] = strconv.FormatBool(p.AllowSendingWithoutReply)
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["photo"], err = b.marshler.MarshalToString(p.Photo); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Photo.IsAttachment() {
		files = append(files, p.Photo)
	}

	src, err := b.Upload(MethodSendPhoto, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewAudio(chatID ChatID, audio *InputFile) SendAudio {
	return SendAudio{
		ChatID: chatID,
		Audio:  audio,
	}
}

// SendAudio send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
func (b Bot) SendAudio(p SendAudio) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["duration"] = strconv.Itoa(p.Duration)
	params["performer"] = p.Performer
	params["title"] = p.Title
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["audio"], err = b.marshler.MarshalToString(p.Audio); err != nil {
		return nil, err
	}

	if params["thumb"], err = b.marshler.MarshalToString(p.Thumb); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Audio.IsAttachment() {
		files = append(files, p.Audio)
	}

	if p.Thumb.IsAttachment() {
		files = append(files, p.Thumb)
	}

	src, err := b.Upload(MethodSendAudio, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewDocument(chatID ChatID, document *InputFile) SendDocument {
	return SendDocument{
		ChatID:   chatID,
		Document: document,
	}
}

// SendDocument send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (b Bot) SendDocument(p SendDocument) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["document"], err = b.marshler.MarshalToString(p.Document); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Document.IsAttachment() {
		files = append(files, p.Document)
	}

	src, err := b.Upload(MethodSendDocument, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewVideo(chatID ChatID, video *InputFile) SendVideo {
	return SendVideo{
		ChatID: chatID,
		Video:  video,
	}
}

// SendVideo send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (b Bot) SendVideo(p SendVideo) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["duration"] = strconv.Itoa(p.Duration)
	params["width"] = strconv.Itoa(p.Width)
	params["height"] = strconv.Itoa(p.Height)
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["supports_streaming"] = strconv.FormatBool(p.SupportsStreaming)
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["video"], err = b.marshler.MarshalToString(p.Video); err != nil {
		return nil, err
	}

	if params["thumb"], err = b.marshler.MarshalToString(p.Thumb); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Video.IsAttachment() {
		files = append(files, p.Video)
	}

	if p.Thumb.IsAttachment() {
		files = append(files, p.Thumb)
	}

	src, err := b.Upload(MethodSendVideo, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewAnimation(chatID ChatID, animation *InputFile) SendAnimation {
	return SendAnimation{
		ChatID:    chatID,
		Animation: animation,
	}
}

// SendAnimation send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (b Bot) SendAnimation(p SendAnimation) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["duration"] = strconv.Itoa(p.Duration)
	params["width"] = strconv.Itoa(p.Width)
	params["height"] = strconv.Itoa(p.Height)
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["animation"], err = b.marshler.MarshalToString(p.Animation); err != nil {
		return nil, err
	}

	if params["thumb"], err = b.marshler.MarshalToString(p.Thumb); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Animation.IsAttachment() {
		files = append(files, p.Animation)
	}

	if p.Thumb.IsAttachment() {
		files = append(files, p.Thumb)
	}

	src, err := b.Upload(MethodSendAnimation, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewVoice(chatID ChatID, voice *InputFile) SendVoice {
	return SendVoice{
		ChatID: chatID,
		Voice:  voice,
	}
}

// SendVoice send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (b Bot) SendVoice(p SendVoice) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["duration"] = strconv.Itoa(p.Duration)
	params["caption"] = p.Caption
	params["parse_mode"] = p.ParseMode
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["voice"], err = b.marshler.MarshalToString(p.Voice); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.Voice.IsAttachment() {
		files = append(files, p.Voice)
	}

	src, err := b.Upload(MethodSendVoice, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewVideoNote(chatID ChatID, videoNote *InputFile) SendVideoNote {
	return SendVideoNote{
		ChatID:    chatID,
		VideoNote: videoNote,
	}
}

// SendVideoNote send video messages. On success, the sent Message is returned.
func (b Bot) SendVideoNote(p SendVideoNote) (*Message, error) {
	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["duration"] = strconv.Itoa(p.Duration)
	params["length"] = strconv.Itoa(p.Length)
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)

	var err error
	if params["video_note"], err = b.marshler.MarshalToString(p.VideoNote); err != nil {
		return nil, err
	}

	if params["thumb"], err = b.marshler.MarshalToString(p.Thumb); err != nil {
		return nil, err
	}

	if params["reply_markup"], err = b.marshler.MarshalToString(p.ReplyMarkup); err != nil {
		return nil, err
	}

	files := make([]*InputFile, 0)
	if p.VideoNote.IsAttachment() {
		files = append(files, p.VideoNote)
	}

	if p.Thumb.IsAttachment() {
		files = append(files, p.Thumb)
	}

	src, err := b.Upload(MethodSendVideoNote, params, files...)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewMediaGroup(chatID ChatID, media ...AlbumMedia) SendMediaGroup {
	return SendMediaGroup{
		ChatID: chatID,
		Media:  media,
	}
}

// SendMediaGroup send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
func (b Bot) SendMediaGroup(p SendMediaGroup) ([]*Message, error) {
	media := make([]string, len(p.Media), 10)
	files := make([]*InputFile, 0)

	for i := range p.Media {
		m := p.Media[i].GetMedia()

		if m.IsAttachment() {
			files = append(files, m)
		}

		src, err := b.marshler.MarshalToString(p.Media[i])
		if err != nil {
			return nil, err
		}

		media = append(media, src)
	}

	params := make(map[string]string)
	params["chat_id"] = p.ChatID.String()
	params["disable_notification"] = strconv.FormatBool(p.DisableNotification)
	params["reply_to_message_id"] = strconv.FormatInt(p.ReplyToMessageID, 10)
	params["media"] = "[" + strings.Join(media, ",") + "]"

	src, err := b.Upload(MethodSendMediaGroup, params, files...)
	if err != nil {
		return nil, err
	}

	result := make([]*Message, 0)
	if err = parseResponseError(b.marshler, src, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewLocation(chatID ChatID, latitude, longitude float64) SendLocation {
	return SendLocation{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// SendLocation send point on the map. On success, the sent Message is returned.
func (b Bot) SendLocation(p SendLocation) (*Message, error) {
	src, err := b.Do(MethodSendLocation, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewLiveLocation(latitude, longitude float64) EditMessageLiveLocation {
	return EditMessageLiveLocation{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// EditMessageLiveLocation edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (b Bot) EditMessageLiveLocation(p EditMessageLiveLocation) (*Message, error) {
	src, err := b.Do(MethodEditMessageLiveLocation, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// StopMessageLiveLocation stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
func (b Bot) StopMessageLiveLocation(p StopMessageLiveLocation) (*Message, error) {
	src, err := b.Do(MethodStopMessageLiveLocation, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewVenue(chatID ChatID, lat, long float64, title, address string) SendVenue {
	return SendVenue{
		ChatID:    chatID,
		Latitude:  lat,
		Longitude: long,
		Title:     title,
		Address:   address,
	}
}

// SendVenue send information about a venue. On success, the sent Message is returned.
func (b Bot) SendVenue(p SendVenue) (*Message, error) {
	src, err := b.Do(MethodSendVenue, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewContact(chatID ChatID, phoneNumber, firstName string) SendContact {
	return SendContact{
		ChatID:      chatID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// SendContact send phone contacts. On success, the sent Message is returned.
func (b Bot) SendContact(p SendContact) (*Message, error) {
	src, err := b.Do(MethodSendContact, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewPoll(chatID ChatID, question string, options ...string) SendPoll {
	return SendPoll{
		ChatID:   chatID,
		Question: question,
		Options:  options,
	}
}

// SendPoll send a native poll. A native poll can't be sent to a private chat. On success, the sent Message is returned.
func (b Bot) SendPoll(p SendPoll) (*Message, error) {
	src, err := b.Do(MethodSendPoll, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// SendDice send a dice, which will have a random value from 1 to 6. On success, the sent Message is returned. (Yes,
// we're aware of the “proper” singular of die. But it's awkward, and we decided to help it change. One dice at a
// time!)
func (b Bot) SendDice(p SendDice) (*Message, error) {
	src, err := b.Do(MethodSendDice, p)
	if err != nil {
		return nil, err
	}

	result := new(Message)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// SendChatAction tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (b Bot) SendChatAction(p SendChatAction) (ok bool, err error) {
	src, err := b.Do(MethodSendChatAction, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (b Bot) GetUserProfilePhotos(p GetUserProfilePhotos) (*UserProfilePhotos, error) {
	src, err := b.Do(MethodGetUserProfilePhotos, p)
	if err != nil {
		return nil, err
	}

	result := new(UserProfilePhotos)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetFile get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (b Bot) GetFile(fid string) (*File, error) {
	src, err := b.Do(MethodGetFile, GetFile{FileID: fid})
	if err != nil {
		return nil, err
	}

	result := new(File)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewKick(chatID ChatID, userID int64) BanChatMember {
	return BanChatMember{
		ChatID: chatID,
		UserID: userID,
	}
}

// BanChatMember kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the 'All Members Are Admins' setting is off in the target group. Otherwise members may only be removed by the group's creator or by the member that added them.
func (b Bot) BanChatMember(p BanChatMember) (ok bool, err error) {
	src, err := b.Do(MethodBanChatMember, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// UnbanChatMember unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
func (b Bot) UnbanChatMember(p UnbanChatMember) (ok bool, err error) {
	src, err := b.Do(MethodUnbanChatMember, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

func NewRestrict(chatID ChatID, userID int64, permissions ChatPermissions) RestrictChatMember {
	return RestrictChatMember{
		ChatID:      chatID,
		UserID:      userID,
		Permissions: &permissions,
	}
}

// restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (b Bot) RestrictChatMember(p RestrictChatMember) (ok bool, err error) {
	src, err := b.Do(MethodRestrictChatMember, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

func NewPromote(chatID ChatID, userID int64) PromoteChatMember {
	return PromoteChatMember{
		ChatID: chatID,
		UserID: userID,
	}
}

// PromoteChatMember promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean  to demote a user. Returns True on success.
func (b Bot) PromoteChatMember(p PromoteChatMember) (ok bool, err error) {
	src, err := b.Do(MethodPromoteChatMember, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// SetChatAdministratorCustomTitle method to set a custom title for an administrator in a supergroup promoted by the b. Returns True on success.
func (b Bot) SetChatAdministratorCustomTitle(p SetChatAdministratorCustomTitle) (ok bool, err error) {
	src, err := b.Do(MethodSetChatAdministratorCustomTitle, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// SetChatPermissions set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success.
func (b Bot) SetChatPermissions(p SetChatPermissions) (ok bool, err error) {
	src, err := b.Do(MethodSetChatPermissions, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// ExportChatInviteLink export an invite link to a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns exported invite link as String on success.
func (b Bot) ExportChatInviteLink(p ExportChatInviteLink) (string, error) {
	src, err := b.Do(MethodExportChatInviteLink, p)
	if err != nil {
		return "", err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return "", err
	}

	var result string
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CreateChatInviteLink create an additional invite link for a chat. The bot must be an administrator in the chat for
// this to work and must have the appropriate admin rights. The link can be revoked using the method
// revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (b Bot) CreateChatInviteLink(p CreateChatInviteLink) (*ChatInviteLink, error) {
	src, err := b.Do(MethodCreateChatInviteLink, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(ChatInviteLink)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// EditChatInviteLink method to edit a non-primary invite link created by the bot. The bot must be an administrator in
// the chat for this to work and must have the appropriate admin rights. Returns the edited invite link as a
// ChatInviteLink object.
func (b Bot) EditChatInviteLink(p EditChatInviteLink) (*ChatInviteLink, error) {
	src, err := b.Do(MethodEditChatInviteLink, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(ChatInviteLink)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// RevokeChatInviteLink method to revoke an invite link created by the bot. If the primary link is revoked, a new link
// is automatically generated. The bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns the revoked invite link as ChatInviteLink object.
func (b Bot) RevokeChatInviteLink(p RevokeChatInviteLink) (*ChatInviteLink, error) {
	src, err := b.Do(MethodRevokeChatInviteLink, p)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return nil, err
	}

	result := new(ChatInviteLink)
	if err = b.marshler.Unmarshal(resp.Result, result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetChatPhoto set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (b Bot) SetChatPhoto(cid int64, photo *InputFile) (ok bool, err error) {
	params := make(map[string]string)
	params["chat_id"] = strconv.FormatInt(cid, 10)

	if params["photo"], err = b.marshler.MarshalToString(photo); err != nil {
		return
	}

	files := make([]*InputFile, 0)
	if photo.IsAttachment() {
		files = append(files, photo)
	}

	src, err := b.Upload(MethodSetChatPhoto, params, files...)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// DeleteChatPhoto delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (b Bot) DeleteChatPhoto(p DeleteChatPhoto) (ok bool, err error) {
	src, err := b.Do(MethodDeleteChatPhoto, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// SetChatTitle change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (b Bot) SetChatTitle(p SetChatTitle) (ok bool, err error) {
	src, err := b.Do(MethodSetChatTitle, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// SetChatDescription change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (b Bot) SetChatDescription(p SetChatDescription) (ok bool, err error) {
	src, err := b.Do(MethodSetChatDescription, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

func NewPin(chatID ChatID, messageID int64) PinChatMessage {
	return PinChatMessage{
		ChatID:    chatID,
		MessageID: messageID,
	}
}

// PinChatMessage pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (b Bot) PinChatMessage(p PinChatMessage) (ok bool, err error) {
	src, err := b.Do(MethodPinChatMessage, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// UnpinChatMessage unpin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (b Bot) UnpinChatMessage(p UnpinChatMessage) (ok bool, err error) {
	src, err := b.Do(MethodUnpinChatMessage, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// UnpinChatMessage method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot
// must be an administrator in the chat for this to work and must have the 'can_pin_messages' admin right in a
// supergroup or 'can_edit_messages' admin right in a channel. Returns True on success.
func (b Bot) UnpinAllChatMessages(p UnpinAllChatMessages) (ok bool, err error) {
	src, err := b.Do(MethodUnpinAllChatMessages, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// LeaveChat leave a group, supergroup or channel. Returns True on success.
func (b Bot) LeaveChat(p LeaveChat) (ok bool, err error) {
	src, err := b.Do(MethodLeaveChat, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// GetChat get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (b Bot) GetChat(p GetChat) (*Chat, error) {
	src, err := b.Do(MethodGetChat, p)
	if err != nil {
		return nil, err
	}

	result := new(Chat)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetChatAdministrators get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (b Bot) GetChatAdministrators(p GetChatAdministrators) ([]*ChatMember, error) {
	src, err := b.Do(MethodGetChatAdministrators, p)
	if err != nil {
		return nil, err
	}

	result := make([]*ChatMember, 0)
	if err = parseResponseError(b.marshler, src, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetChatMemberCount get the number of members in a chat. Returns Int on success.
func (b Bot) GetChatMemberCount(p GetChatMemberCount) (int, error) {
	src, err := b.Do(MethodGetChatMemberCount, p)
	if err != nil {
		return 0, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return 0, err
	}

	var result int
	if err = b.marshler.Unmarshal(resp.Result, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// GetChatMember get information about a member of a chat. Returns a ChatMember object on success.
func (b Bot) GetChatMember(p GetChatMember) (*ChatMember, error) {
	src, err := b.Do(MethodGetChatMember, p)
	if err != nil {
		return nil, err
	}

	result := new(ChatMember)
	if err = parseResponseError(b.marshler, src, result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetChatStickerSet set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b Bot) SetChatStickerSet(p SetChatStickerSet) (ok bool, err error) {
	src, err := b.Do(MethodSetChatStickerSet, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// DeleteChatStickerSet delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b Bot) DeleteChatStickerSet(p DeleteChatStickerSet) (ok bool, err error) {
	src, err := b.Do(MethodDeleteChatStickerSet, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

func NewAnswerCallback(callbackQueryID string) AnswerCallbackQuery {
	return AnswerCallbackQuery{CallbackQueryID: callbackQueryID}
}

// AnswerCallbackQuery send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (b Bot) AnswerCallbackQuery(p AnswerCallbackQuery) (ok bool, err error) {
	src, err := b.Do(MethodAnswerCallbackQuery, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// SetMyCommands change the list of the bot's commands. Returns True on success.
func (b Bot) SetMyCommands(p SetMyCommands) (ok bool, err error) {
	src, err := b.Do(MethodSetMyCommands, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// DeleteMyCommands delete the list of the bot's commands for the given scope and user language. After deletion, higher
// level commands will be shown to affected users. Returns True on success.
func (b Bot) DeleteMyCommands(p DeleteMyCommands) (ok bool, err error) {
	src, err := b.Do(MethodDeleteMyCommands, p)
	if err != nil {
		return ok, err
	}

	resp := new(Response)
	if err = b.marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if err = b.marshler.Unmarshal(resp.Result, &ok); err != nil {
		return
	}

	return
}

// GetMyCommands get the current list of the bot's commands. Requires no parameters. Returns Array of BotCommand on
// success.
func (b Bot) GetMyCommands(p GetMyCommands) ([]*BotCommand, error) {
	src, err := b.Do(MethodGetMyCommands, p)
	if err != nil {
		return nil, err
	}

	result := make([]*BotCommand, 0)
	if err = parseResponseError(b.marshler, src, &result); err != nil {
		return nil, err
	}

	return result, nil
}
