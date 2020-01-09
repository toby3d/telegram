package telegram

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	http "github.com/valyala/fasthttp"
	"golang.org/x/text/language"
)

type (
	// User represents a Telegram user or bot.
	User struct {
		// Unique identifier for this user or bot
		ID int `json:"id"`

		// True, if this user is a bot
		IsBot bool `json:"is_bot"`

		// User‘s or bot’s first name
		FirstName string `json:"first_name"`

		// User‘s or bot’s last name
		LastName string `json:"last_name,omitempty"`

		// User‘s or bot’s username
		Username string `json:"username,omitempty"`

		// IETF language tag of the user's language
		LanguageCode string `json:"language_code,omitempty"`
	}

	// Chat represents a chat.
	Chat struct {
		// Unique identifier for this chat.
		ID int64 `json:"id"`

		// Type of chat, can be either "private", "group", "supergroup" or
		// "channel"
		Type string `json:"type"`

		// Title, for supergroups, channels and group chats
		Title string `json:"title,omitempty"`

		// Username, for private chats, supergroups and channels if available
		Username string `json:"username,omitempty"`

		// First name of the other party in a private chat
		FirstName string `json:"first_name,omitempty"`

		// Last name of the other party in a private chat
		LastName string `json:"last_name,omitempty"`

		// Chat photo. Returned only in getChat.
		Photo *ChatPhoto `json:"photo,omitempty"`

		// Description, for groups, supergroups and channel chats. Returned only in getChat.
		Description string `json:"description,omitempty"`

		// Chat invite link, for supergroups and channel chats. Returned only in
		// getChat.
		InviteLink string `json:"invite_link,omitempty"`

		// Pinned message, for groups, supergroups and channels. Returned only in getChat.
		PinnedMessage *Message `json:"pinned_message,omitempty"`

		// Default chat member permissions, for groups and supergroups. Returned only in getChat.
		Permissions *ChatPermissions `json:"permissions,omitempty"`

		// For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user. Returned only in getChat.
		SlowModeDelay int `json:"slow_mode_delay,omitempty"`

		// For supergroups, name of Group sticker set. Returned only in getChat.
		StickerSetName string `json:"sticker_set_name,omitempty"`

		// True, if the bot can change group the sticker set. Returned only in
		// getChat.
		CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	}

	// Message represents a message.
	Message struct {
		// Unique message identifier inside this chat
		ID int `json:"message_id"`

		// Sender, empty for messages sent to channels
		From *User `json:"from,omitempty"`

		// Date the message was sent in Unix time
		Date int64 `json:"date"`

		// Conversation the message belongs to
		Chat *Chat `json:"chat"`

		// For forwarded messages, sender of the original message
		ForwardFrom *User `json:"forward_from,omitempty"`

		// For messages forwarded from channels, information about the original
		// channel
		ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

		// For messages forwarded from channels, identifier of the original
		// message in the channel
		ForwardFromMessageID int `json:"forward_from_message_id,omitempty"`

		// For messages forwarded from channels, signature of the post author if
		// present
		ForwardSignature string `json:"forward_signature,omitempty"`

		// Sender's name for messages forwarded from users who disallow adding a
		// link to their account in forwarded messages
		ForwardSenderName string `json:"forward_sender_name,omitempty"`

		// For forwarded messages, date the original message was sent in Unix
		// time
		ForwardDate int64 `json:"forward_date,omitempty"`

		// For replies, the original message. Note that the Message object in
		// this field will not contain further reply_to_message fields even if it
		// itself is a reply.
		ReplyToMessage *Message `json:"reply_to_message,omitempty"`

		// Date the message was last edited in Unix time
		EditDate int64 `json:"edit_date,omitempty"`

		// The unique identifier of a media message group this message belongs to
		MediaGroupID string `json:"media_group_id,omitempty"`

		// Signature of the post author for messages in channels
		AuthorSignature string `json:"author_signature,omitempty"`

		// For text messages, the actual UTF-8 text of the message, 0-4096
		// characters.
		Text string `json:"text,omitempty"`

		// For text messages, special entities like usernames, URLs, bot
		// commands, etc. that appear in the text
		Entities []*MessageEntity `json:"entities,omitempty"`

		// For messages with a caption, special entities like usernames, URLs,
		// bot commands, etc. that appear in the caption
		CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

		// Message is an audio file, information about the file
		Audio *Audio `json:"audio,omitempty"`

		// Message is a general file, information about the file
		Document *Document `json:"document,omitempty"`

		// Message is an animation, information about the animation. For backward
		// compatibility, when this field is set, the document field will also be
		// set
		Animation *Animation `json:"animation,omitempty"`

		// Message is a game, information about the game.
		Game *Game `json:"game,omitempty"`

		// Message is a photo, available sizes of the photo
		Photo []*PhotoSize `json:"photo,omitempty"`

		// Message is a sticker, information about the sticker
		Sticker *Sticker `json:"sticker,omitempty"`

		// Message is a video, information about the video
		Video *Video `json:"video,omitempty"`

		// Message is a voice message, information about the file
		Voice *Voice `json:"voice,omitempty"`

		// Message is a video note, information about the video message
		VideoNote *VideoNote `json:"video_note,omitempty"`

		// Caption for the document, photo or video, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Message is a shared contact, information about the contact
		Contact *Contact `json:"contact,omitempty"`

		// Message is a shared location, information about the location
		Location *Location `json:"location,omitempty"`

		// Message is a venue, information about the venue
		Venue *Venue `json:"venue,omitempty"`

		// Message is a native poll, information about the poll
		Poll *Poll `json:"poll,omitempty"`

		// New members that were added to the group or supergroup and information
		// about them (the bot itself may be one of these members)
		NewChatMembers []*User `json:"new_chat_members,omitempty"`

		// A member was removed from the group, information about them (this
		// member may be the bot itself)
		LeftChatMember *User `json:"left_chat_member,omitempty"`

		// A chat title was changed to this value
		NewChatTitle string `json:"new_chat_title,omitempty"`

		// A chat photo was change to this value
		NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`

		// Service message: the chat photo was deleted
		DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

		// Service message: the group has been created
		GroupChatCreated bool `json:"group_chat_created,omitempty"`

		// Service message: the supergroup has been created. This field can‘t be
		// received in a message coming through updates, because bot can’t be a
		// member of a supergroup when it is created. It can only be found in
		// reply_to_message if someone replies to a very first message in a
		// directly created supergroup.
		SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

		// Service message: the channel has been created. This field can‘t be
		// received in a message coming through updates, because bot can’t be a
		// member of a channel when it is created. It can only be found in
		// reply_to_message if someone replies to a very first message in a
		// channel.
		ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

		// The group has been migrated to a supergroup with the specified
		// identifier.
		MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

		// The supergroup has been migrated from a group with the specified
		// identifier.
		MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

		// Specified message was pinned. Note that the Message object in this
		// field will not contain further reply_to_message fields even if it is
		// itself a reply.
		PinnedMessage *Message `json:"pinned_message,omitempty"`

		// Message is an invoice for a payment, information about the invoice.
		Invoice *Invoice `json:"invoice,omitempty"`

		// Message is a service message about a successful payment, information
		// about the payment.
		SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

		// The domain name of the website on which the user has logged in.
		ConnectedWebsite string `json:"connected_website,omitempty"`

		// Telegram Passport data
		PassportData *PassportData `json:"passport_data,omitempty"`

		// Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}

	// MessageEntity represents one special entity in a text message. For
	// example, hashtags, usernames, URLs, etc.
	MessageEntity struct {
		// Type of the entity. Can be mention (@username), hashtag, bot_command,
		// url, email, bold (bold text), italic (italic text), code (monowidth
		// string), pre (monowidth block), text_link (for clickable text URLs),
		// text_mention (for users without usernames)
		Type string `json:"type"`

		// For "text_link" only, url that will be opened after user taps on the
		// text
		URL string `json:"url,omitempty"`

		// Offset in UTF-16 code units to the start of the entity
		Offset int `json:"offset"`

		// Length of the entity in UTF-16 code units
		Length int `json:"length"`

		// For "text_mention" only, the mentioned user
		User *User `json:"user,omitempty"`
	}

	// PhotoSize represents one size of a photo or a file / sticker thumbnail.
	PhotoSize struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots.
		// Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Photo width
		Width int `json:"width"`

		// Photo height
		Height int `json:"height"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Audio represents an audio file to be treated as music by the Telegram
	// clients.
	Audio struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// Performer of the audio as defined by sender or by audio tags
		Performer string `json:"performer,omitempty"`

		// Title of the audio as defined by sender or by audio tags
		Title string `json:"title,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`

		// Thumbnail of the album cover to which the music file belongs
		Thumb *PhotoSize `json:"thumb,omitempty"`
	}

	// Document represents a general file (as opposed to photos, voice messages
	// and audio files).
	Document struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Document thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// Original filename as defined by sender
		FileName string `json:"file_name,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Video represents a video file.
	Video struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Video width as defined by sender
		Width int `json:"width"`

		// Video height as defined by sender
		Height int `json:"height"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// Mime type of a file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Animation provide an animation for your game so that it looks stylish in
	// chats (check out Lumberjack for an example). This object represents an
	// animation file to be displayed in the message containing a game.
	Animation struct {
		// Unique file identifier
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Video width as defined by sender
		Width int `json:"width"`

		// Video height as defined by sender
		Height int `json:"height"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// Animation thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// Original animation filename as defined by sender
		FileName string `json:"file_name,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Voice represents a voice note.
	Voice struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// VideoNote represents a video message (available in Telegram apps as of
	// v.4.0).
	VideoNote struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// Video width and height (diameter of the video message) as defined by sender
		Length int `json:"length"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Contact represents a phone contact.
	Contact struct {
		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name,omitempty"`

		// Contact's user identifier in Telegram
		UserID int `json:"user_id,omitempty"`

		// Additional data about the contact in the form of a vCard
		VCard string `json:"vcard,omitempty"`
	}

	// Location represents a point on the map.
	Location struct {
		// Longitude as defined by sender
		Longitude float32 `json:"longitude"`

		// Latitude as defined by sender
		Latitude float32 `json:"latitude"`
	}

	// Venue represents a venue.
	Venue struct {
		// Venue location
		Location *Location `json:"location"`

		// Name of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue
		FoursquareID string `json:"foursquare_id,omitempty"`

		// Foursquare type of the venue. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
		FoursquareType string `json:"foursquare_type,omitempty"`
	}

	// This object contains information about one answer option in a poll.
	PollOption struct {
		// Option text, 1-100 characters
		Text string `json:"text"`

		// Number of users that voted for this option
		VoterCount int `json:"voter_count"`
	}

	// This object contains information about a poll.
	Poll struct {
		// Unique poll identifier
		ID string `json:"id"`

		// Poll question, 1-255 characters
		Question string `json:"question"`

		// List of poll options
		Options []*PollOption `json:"options"`

		// True, if the poll is closed
		IsClosed bool `json:"is_closed"`
	}

	// UserProfilePhotos represent a user's profile pictures.
	UserProfilePhotos struct {
		// Total number of profile pictures the target user has
		TotalCount int `json:"total_count"`

		// Requested profile pictures (in up to 4 sizes each)
		Photos [][]*PhotoSize `json:"photos"`
	}

	// File represents a file ready to be downloaded. The file can be downloaded
	// via the link https://api.telegram.org/file/bot<token>/<file_path>. It is
	// guaranteed that the link will be valid for at least 1 hour. When the link
	// expires, a new one can be requested by calling getFile.
	//
	// Maximum file size to download is 20 MB
	File struct {
		// Identifier for this file, which can be used to download or reuse the file
		FileID string `json:"file_id"`

		// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		FileUniqueID string `json:"file_unique_id"`

		// File size, if known
		FileSize int `json:"file_size,omitempty"`

		// File path. Use https://api.telegram.org/file/bot<token>/<file_path> to
		// get the file.
		FilePath string `json:"file_path,omitempty"`
	}

	// ReplyKeyboardMarkup represents a custom keyboard with reply options (see
	// Introduction to bots for details and examples).
	ReplyKeyboardMarkup struct {
		// Array of button rows, each represented by an Array of KeyboardButton
		// objects
		Keyboard [][]*KeyboardButton `json:"keyboard"`

		// Requests clients to resize the keyboard vertically for optimal fit
		// (e.g., make the keyboard smaller if there are just two rows of
		// buttons). Defaults to false, in which case the custom keyboard is
		// always of the same height as the app's standard keyboard.
		ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

		// Requests clients to hide the keyboard as soon as it's been used. The
		// keyboard will still be available, but clients will automatically
		// display the usual letter-keyboard in the chat – the user can press a
		// special button in the input field to see the custom keyboard again.
		// Defaults to false.
		OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

		// Use this parameter if you want to show the keyboard to specific users
		// only. Targets: 1) users that are @mentioned in the text of the Message
		// object; 2) if the bot's message is a reply (has reply_to_message_id),
		// sender of the original message.
		//
		// Example: A user requests to change the bot‘s language, bot replies to
		// the request with a keyboard to select the new language. Other users in
		// the group don’t see the keyboard.
		Selective bool `json:"selective,omitempty"`
	}

	// KeyboardButton represents one button of the reply keyboard. For simple
	// text buttons String can be used instead of this object to specify text of
	// the button. Optional fields are mutually exclusive.
	KeyboardButton struct {
		// Text of the button. If none of the optional fields are used, it will
		// be sent to the bot as a message when the button is pressed
		Text string `json:"text"`

		// If True, the user's phone number will be sent as a contact when the
		// button is pressed. Available in private chats only
		RequestContact bool `json:"request_contact,omitempty"`

		// If True, the user's current location will be sent when the button is
		// pressed. Available in private chats only
		RequestLocation bool `json:"request_location,omitempty"`
	}

	// ReplyKeyboardRemove will remove the current custom keyboard and display
	// the default letter-keyboard. By default, custom keyboards are displayed
	// until a new keyboard is sent by a b. An exception is made for one-time
	// keyboards that are hidden immediately after the user presses a button
	// (see ReplyKeyboardMarkup).
	ReplyKeyboardRemove struct {
		// Requests clients to remove the custom keyboard (user will not be able
		// to summon this keyboard; if you want to hide the keyboard from sight
		// but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
		RemoveKeyboard bool `json:"remove_keyboard"`

		// Use this parameter if you want to remove the keyboard for specific
		// users only. Targets: 1) users that are @mentioned in the text of the
		// Message object; 2) if the bot's message is a reply (has
		// reply_to_message_id), sender of the original message.
		//
		// Example: A user votes in a poll, bot returns confirmation message in
		// reply to the vote and removes the keyboard for that user, while still
		// showing the keyboard with poll options to users who haven't voted yet.
		Selective bool `json:"selective,omitempty"`
	}

	// InlineKeyboardMarkup represents an inline keyboard that appears right next
	// to the message it belongs to.
	InlineKeyboardMarkup struct {
		// Array of button rows, each represented by an Array of
		// InlineKeyboardButton objects
		InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
	}

	// InlineKeyboardButton represents one button of an inline keyboard. You
	// must use exactly one of the optional fields.
	InlineKeyboardButton struct {
		// Label text on the button
		Text string `json:"text"`

		// HTTP url to be opened when button is pressed
		URL string `json:"url,omitempty"`

		// An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram
		// Login Widget.
		LoginURL *LoginURL `json:"login_url,omitempty"`

		// Data to be sent in a callback query to the bot when button is pressed,
		// 1-64 bytes
		CallbackData string `json:"callback_data,omitempty"`

		// If set, pressing the button will prompt the user to select one of
		// their chats, open that chat and insert the bot‘s username and the
		// specified inline query in the input field. Can be empty, in which
		// case just the bot’s username will be inserted.
		//
		// Note: This offers an easy way for users to start using your bot in
		// inline mode when they are currently in a private chat with it.
		// Especially useful when combined with switch_pm… actions – in this case
		// the user will be automatically returned to the chat they switched
		// from, skipping the chat selection screen.
		SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

		// If set, pressing the button will insert the bot‘s username and the
		// specified inline query in the current chat's input field. Can be
		// empty, in which case only the bot’s username will be inserted.
		//
		// This offers a quick way for the user to open your bot in inline mode
		// in the same chat – good for selecting something from multiple options.
		SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

		// Description of the game that will be launched when the user presses
		// the button.
		//
		// NOTE: This type of button must always be the first button in the
		// first row.
		CallbackGame *CallbackGame `json:"callback_game,omitempty"`

		// Specify True, to send a Pay button.
		//
		// NOTE: This type of button must always be the first button in the
		// first row.
		Pay bool `json:"pay,omitempty"`
	}

	// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
	LoginURL struct {
		// An HTTP URL to be opened with user authorization data added to the query string when the button is
		// pressed. If the user refuses to provide authorization data, the original URL without information
		// about the user will be opened. The data added is the same as described in Receiving authorization
		// data.
		//
		// NOTE: You must always check the hash of the received data to verify the authentication and the
		// integrity of the data as described in Checking authorization.
		URL string `json:"url"`

		// New text of the button in forwarded messages.
		ForwardText string `json:"forward_text,omitempty"`

		// Username of a bot, which will be used for user authorization. See Setting up a bot for more
		// details. If not specified, the current bot's username will be assumed. The url's domain must be the
		// same as the domain linked with the b. See Linking your domain to the bot for more details.
		BotUsername string `json:"bot_username,omitempty"`

		// Pass true to request the permission for your bot to send messages to the user.
		RequestWriteAccess bool `json:"request_write_access,omitempty"`
	}

	// CallbackQuery represents an incoming callback query from a callback button
	// in an inline keyboard. If the button that originated the query was
	// attached to a message sent by the bot, the field message will be present.
	// If the button was attached to a message sent via the bot (in inline mode),
	// the field inline_message_id will be present. Exactly one of the fields
	// data or game_short_name will be present.
	//
	// NOTE: After the user presses a callback button, Telegram clients will
	// display a progress bar until you call answerCallbackQuery. It is,
	// therefore, necessary to react by calling answerCallbackQuery even if no
	// notification to the user is needed (e.g., without specifying any of the
	// optional ).
	CallbackQuery struct {
		// Unique identifier for this query
		ID string `json:"id"`

		// Identifier of the message sent via the bot in inline mode, that
		// originated the query.
		InlineMessageID string `json:"inline_message_id,omitempty"`

		// Global identifier, uniquely corresponding to the chat to which the
		// message with the callback button was sent. Useful for high scores in
		// games.
		ChatInstance string `json:"chat_instance"`

		// Data associated with the callback button. Be aware that a bad client
		// can send arbitrary data in this field.
		Data string `json:"data,omitempty"`

		// Short name of a Game to be returned, serves as the unique identifier
		// for the game
		GameShortName string `json:"game_short_name,omitempty"`

		// Sender
		From *User `json:"from"`

		// Message with the callback button that originated the query. Note that
		// message content and message date will not be available if the message
		// is too old
		Message *Message `json:"message,omitempty"`
	}

	// ForceReply display a reply interface to the user (act as if the user has
	// selected the bot‘s message and tapped ’Reply'). This can be extremely
	// useful if you want to create user-friendly step-by-step interfaces without
	// having to sacrifice privacy mode.
	ForceReply struct {
		// Shows reply interface to the user, as if they manually selected the
		// bot‘s message and tapped ’Reply'
		ForceReply bool `json:"force_reply"`

		// Use this parameter if you want to force reply from specific users
		// only. Targets: 1) users that are @mentioned in the text of the Message
		// object; 2) if the bot's message is a reply (has reply_to_message_id),
		// sender of the original message.
		Selective bool `json:"selective,omitempty"`
	}

	// ChatPhoto represents a chat photo.
	ChatPhoto struct {
		// File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
		SmallFileID string `json:"small_file_id"`

		// Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		SmallFileUniqueID string `json:"small_file_unique_id"`

		// File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
		BigFileID string `json:"big_file_id"`

		// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
		BigFileUniqueID string `json:"big_file_unique_id"`
	}

	// ChatMember contains information about one member of a chat.
	ChatMember struct {
		// Information about the user
		User *User `json:"user"`

		// The member's status in the chat. Can be "creator", "administrator", "member", "restricted", "left"
		// or "kicked"
		Status string `json:"status"`

		// Owner and administrators only. Custom title for this user
		CustomTitle string `json:"custom_title,omitempty"`

		// Restictred and kicked only. Date when restrictions will be lifted for this user, unix time
		UntilDate int64 `json:"until_date,omitempty"`

		// Administrators only. True, if the bot is allowed to edit administrator privileges of that user
		CanBeEdited bool `json:"can_be_edited,omitempty"`

		// Administrators only. True, if the administrator can post in the channel, channels only
		CanPostMessages bool `json:"can_post_messages,omitempty"`

		// Administrators only. True, if the administrator can edit messages of other users, channels only
		CanEditMessages bool `json:"can_edit_messages,omitempty"`

		// Administrators only. True, if the administrator can delete messages of other users
		CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

		// Administrators only. True, if the administrator can restrict, ban or
		// unban chat members
		CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

		// Administrators only. True, if the administrator can add new
		// administrators with a subset of his own privileges or demote
		// administrators that he has promoted, directly or indirectly (promoted
		// by administrators that were appointed by the user)
		CanPromoteMembers bool `json:"can_promote_members,omitempty"`

		// Restricted only. True, if the user is a member of the chat at the moment of the request
		IsMember bool `json:"is_member,omitempty"`

		ChatPermissions
	}

	// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
	ChatPermissions struct {
		// True, if the user is allowed to send text messages, contacts, locations and venues
		CanSendMessages bool `json:"can_send_messages,omitempty"`

		// True, if the user is allowed to send audios, documents, photos, videos, video notes and voice
		// notes, implies can_send_messages
		CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

		// True, if the user is allowed to send polls, implies can_send_messages
		CanSendPolls bool `json:"can_send_polls,omitempty"`

		// True, if the user is allowed to send animations, games, stickers and use inline bots, implies
		// can_send_media_messages
		CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

		// True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
		CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

		// True, if the user is allowed to change the chat title, photo and other settings. Ignored in public
		// supergroups
		CanChangeInfo bool `json:"can_change_info,omitempty"`

		// True, if the user is allowed to invite new users to the chat
		CanInviteUsers bool `json:"can_invite_users,omitempty"`

		// True, if the user is allowed to pin messages. Ignored in public supergroups
		CanPinMessages bool `json:"can_pin_messages,omitempty"`
	}

	// ResponseParameters contains information about why a request was
	// unsuccessful.
	ResponseParameters struct {
		// The group has been migrated to a supergroup with the specified
		// identifier.
		MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

		// In case of exceeding flood control, the number of seconds left to wait
		// before the request can be repeated
		RetryAfter int `json:"retry_after,omitempty"`
	}

	// InputMedia represents the content of a media message to be sent.
	InputMedia interface {
		GetMedia() *InputFile
	}

	AlbumMedia interface {
		GetMedia() *InputFile
		isAlbumMedia()
	}

	// InputMediaPhoto represents a photo to be sent.
	InputMediaPhoto struct {
		// Type of the result, must be photo
		Type string `json:"type"`

		// File to send. Pass a file_id to send a file that exists on the
		// Telegram servers (recommended), pass an HTTP URL for Telegram to get
		// a file from the Internet, or pass "attach://<file_attach_name>" to
		// upload a new one using multipart/form-data under <file_attach_name>
		// name.
		Media *InputFile `json:"media"`

		// Caption of the photo to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`
	}

	// InputMediaVideo represents a video to be sent.
	InputMediaVideo struct {
		// Type of the result, must be video
		Type string `json:"type"`

		// File to send. Pass a file_id to send a file that exists on the
		// Telegram servers (recommended), pass an HTTP URL for Telegram to get
		// a file from the Internet, or pass "attach://<file_attach_name>" to
		// upload a new one using multipart/form-data under <file_attach_name>
		// name.
		Media *InputFile `json:"media"`

		// Caption of the video to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Video width
		Width int `json:"width,omitempty"`

		// Video height
		Height int `json:"height,omitempty"`

		// Video duration
		Duration int `json:"duration,omitempty"`

		// Pass true, if the uploaded video is suitable for streaming
		SupportsStreaming bool `json:"supports_streaming,omitempty"`
	}

	// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC
	// video without sound) to be sent.
	InputMediaAnimation struct {
		// Type of the result, must be animation
		Type string `json:"type"`

		// File to send. Pass a file_id to send a file that exists on the
		// Telegram servers (recommended), pass an HTTP URL for Telegram to get
		// a file from the Internet, or pass "attach://<file_attach_name>" to
		// upload a new one using multipart/form-data under <file_attach_name
		// name.
		Media *InputFile `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and
		// less than 200 kB in size. A thumbnail‘s width and height should not
		// exceed 90. Ignored if the file is not uploaded using
		// multipart/form-data. Thumbnails can’t be reused and can be only
		// uploaded as a new file, so you can pass "attach://<file_attach_name>"
		// if the thumbnail was uploaded using multipart/form-data under
		// <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Caption of the animation to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Animation width
		Width int `json:"width,omitempty"`

		// Animation height
		Height int `json:"height,omitempty"`

		// Animation duration
		Duration int `json:"duration,omitempty"`
	}

	// InputMediaAudio represents an audio file to be treated as music to be sent.
	InputMediaAudio struct {
		// Type of the result, must be audio
		Type string `json:"type"`

		// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.
		Media *InputFile `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Caption of the audio to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`

		// Duration of the audio in seconds
		Duration int `json:"duration,omitempty"`

		// Performer of the audio
		Performer string `json:"performer,omitempty"`

		// Title of the audio
		Title string `json:"title,omitempty"`
	}

	// InputMediaDocument represents a general file to be sent.
	InputMediaDocument struct {
		// Type of the result, must be document
		Type string `json:"type"`

		// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.
		Media *InputFile `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
		Thumb *InputFile `json:"thumb,omitempty"`

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`
	}

	// InputFile represents the contents of a file to be uploaded. Must be poste using multipart/form-data in the usual way that files are uploaded via the browser.
	InputFile struct {
		ID         string    `json:"-"`
		URI        *http.URI `json:"-"`
		Attachment *os.File  `json:"-"`
	}
)

// Language parse LanguageCode of current user and returns language.Tag.
func (u User) Language() language.Tag {
	tag, err := language.Parse(u.LanguageCode)
	if err != nil {
		tag = language.Und
	}

	return tag
}

// FullName returns the full name of user or FirstName if LastName is not available.
func (u User) FullName() string {
	if u.FirstName == "" {
		return ""
	}

	if u.HasLastName() {
		return u.FirstName + " " + u.LastName
	}

	return u.FirstName
}

// HaveLastName checks what the current user has a LastName.
func (u User) HasLastName() bool { return u.LastName != "" }

// HaveUsername checks what the current user has a username.
func (u User) HasUsername() bool { return u.Username != "" }

// IsPrivate checks that the current chat is a private chat with single user.
func (c Chat) IsPrivate() bool { return strings.EqualFold(c.Type, ChatPrivate) }

// IsGroup checks that the current chat is a group.
func (c Chat) IsGroup() bool { return strings.EqualFold(c.Type, ChatGroup) }

// IsSuperGroup checks that the current chat is a supergroup.
func (c Chat) IsSuperGroup() bool { return strings.EqualFold(c.Type, ChatSuperGroup) }

// IsChannel checks that the current chat is a channel.
func (c Chat) IsChannel() bool { return strings.EqualFold(c.Type, ChatChannel) }

// HasPinnedMessage checks that the current chat has a pinned message.
func (c Chat) HasPinnedMessage() bool { return c.PinnedMessage != nil }

// HasStickerSet checks that the current chat has a sticker set.
func (c Chat) HasStickerSet() bool { return c.StickerSetName != "" }

// FullName returns the full name of chat or FirstName if LastName is not available.
func (c Chat) FullName() string {
	if c.FirstName == "" {
		return ""
	}

	if c.HasLastName() {
		return c.FirstName + " " + c.LastName
	}

	return c.FirstName
}

// HaveLastName checks what the current user has a LastName.
func (c Chat) HasLastName() bool { return c.LastName != "" }

// HaveUsername checks what the current user has a username.
func (c Chat) HasUsername() bool { return c.Username != "" }

func (c Chat) HasDescription() bool { return c.Description != "" }

func (c Chat) HasInviteLink() bool { return c.InviteLink != "" }

// IsCommand checks that the current message is a bot command.
func (m Message) IsCommand() bool {
	return m.HasEntities() && m.Entities[0].IsBotCommand() && m.Entities[0].Offset == 0
}

// IsCommandEqual checks that the current message is a specific bot command.
func (m Message) IsCommandEqual(command string) bool {
	return m.IsCommand() && strings.EqualFold(m.Command(), command)
}

// Command returns identifier of the bot command without bot username, if it was available
func (m Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	return strings.Split(m.RawCommand(), "@")[0]
}

// RawCommand returns identifier of the bot command with bot username, if it was available
func (m Message) RawCommand() string {
	if !m.IsCommand() {
		return ""
	}

	return string([]rune(m.Text)[1:m.Entities[0].Length])
}

// HasCommandArgument checks that the current command message contains argument.
func (m Message) HasCommandArgument() bool {
	return m.IsCommand() && m.Entities[0].IsBotCommand() && len([]rune(m.Text)) != m.Entities[0].Length
}

// CommandArgument returns raw command argument.
func (m Message) CommandArgument() string {
	if !m.HasCommandArgument() {
		return ""
	}

	return string([]rune(m.Text)[m.Entities[0].Length+1:])
}

// IsReply checks that the current message is a reply on other message.
func (m Message) IsReply() bool { return m.ReplyToMessage != nil }

// IsForward checks that the current message is a forward of other message.
func (m Message) IsForward() bool { return m.ForwardDate > 0 }

// Time parse current message Date and returns time.Time.
func (m Message) Time() time.Time {
	if m.Date <= 0 {
		return time.Time{}
	}

	return time.Unix(m.Date, 0)
}

// ForwardTime parse current message ForwardDate and returns time.Time.
func (m Message) ForwardTime() time.Time {
	if !m.IsForward() {
		return time.Time{}
	}

	return time.Unix(m.ForwardDate, 0)
}

// EditTime parse current message EditDate and returns time.Time.
func (m Message) EditTime() time.Time {
	if !m.HasBeenEdited() {
		return time.Time{}
	}

	return time.Unix(m.EditDate, 0)
}

// HasBeenEdited checks that the current message has been edited.
func (m Message) HasBeenEdited() bool { return m.EditDate > 0 }

// IsText checks that the current message is just a text message.
func (m Message) IsText() bool { return m.Text != "" }

// IsAudio checks that the current message is a audio.
func (m Message) IsAudio() bool { return m.Audio != nil }

// IsDocument checks that the current message is a document.
func (m Message) IsDocument() bool { return m.Document != nil }

// IsGame checks that the current message is a game.
func (m Message) IsGame() bool { return m.Game != nil }

// IsPhoto checks that the current message is a photo.
func (m Message) IsPhoto() bool { return len(m.Photo) > 0 }

// IsSticker checks that the current message is a sticker.
func (m Message) IsSticker() bool { return m.Sticker != nil }

// IsVideo checks that the current message is a video.
func (m Message) IsVideo() bool { return m.Video != nil }

// IsVoice checks that the current message is a voice.
func (m Message) IsVoice() bool { return m.Voice != nil }

// IsVideoNote checks that the current message is a video note.
func (m Message) IsVideoNote() bool { return m.VideoNote != nil }

// IsContact checks that the current message is a contact.
func (m Message) IsContact() bool { return m.Contact != nil }

// IsLocation checks that the current message is a location.
func (m Message) IsLocation() bool { return m.Location != nil }

// IsVenue checks that the current message is a venue.
func (m Message) IsVenue() bool { return m.Venue != nil }

// IsAnimation checks that the current message is a animation.
func (m Message) IsAnimation() bool { return m.Animation != nil }

// IsNewChatMembersEvent checks that the current message is a event of entry of new members.
func (m Message) IsNewChatMembersEvent() bool { return len(m.NewChatMembers) > 0 }

// IsLeftChatMemberEvent checks that the current message is a event of members exit.
func (m Message) IsLeftChatMemberEvent() bool { return m.LeftChatMember != nil }

// IsNewChatTitleEvent checks that the current message is a event of setting a new chat title.
func (m Message) IsNewChatTitleEvent() bool { return m.NewChatTitle != "" }

// IsNewChatPhotoEvent checks that the current message is a event of setting a new chat avatar.
func (m Message) IsNewChatPhotoEvent() bool { return len(m.NewChatPhoto) > 0 }

// IsDeleteChatPhotoEvent checks that the current message is a event of deleting a chat avatar.
func (m Message) IsDeleteChatPhotoEvent() bool { return m.DeleteChatPhoto }

// IsGroupChatCreatedEvent checks that the current message is a event of creating a new group.
func (m Message) IsGroupChatCreatedEvent() bool { return m.GroupChatCreated }

// IsSupergroupChatCreatedEvent checks that the current message is a event of creating a new supergroup.
func (m Message) IsSupergroupChatCreatedEvent() bool { return m.SupergroupChatCreated }

// IsChannelChatCreatedEvent checks that the current message is a event of creating a new channel.
func (m Message) IsChannelChatCreatedEvent() bool { return m.ChannelChatCreated }

// IsPinnedMessage checks that the current message is a event of pinning another message.
func (m Message) IsPinnedMessage() bool { return m.PinnedMessage != nil }

// IsInvoice checks that the current message is a invoice.
func (m Message) IsInvoice() bool { return m.Invoice != nil }

// IsSuccessfulPayment checks that the current message is a event of successful payment.
func (m Message) IsSuccessfulPayment() bool { return m.SuccessfulPayment != nil }

// IsPoll checks that the current message is a poll.
func (m Message) IsPoll() bool { return m.Poll != nil }

// HasEntities checks that the current message contains entities.
func (m Message) HasEntities() bool { return len(m.Entities) > 0 }

// HasCaptionEntities checks that the current media contains entities in caption.
func (m Message) HasCaptionEntities() bool { return len(m.CaptionEntities) > 0 }

// HasMentions checks that the current message contains mentions.
func (m Message) HasMentions() bool {
	if !m.HasEntities() {
		return false
	}

	for _, entity := range m.Entities {
		if !entity.IsMention() && !entity.IsTextMention() {
			continue
		}

		return true
	}

	return false
}

// HasCaptionMentions checks that the current media contains mentions in caption.
func (m Message) HasCaptionMentions() bool {
	if !m.HasCaptionEntities() {
		return false
	}

	for _, entity := range m.CaptionEntities {
		if !entity.IsMention() && !entity.IsTextMention() {
			continue
		}

		return true
	}

	return false
}

// HasCaption checks that the current media has caption.
func (m Message) HasCaption() bool { return m.Caption != "" }

// HasAuthorSignature checks that the current channel post has author signature.
func (m Message) HasAuthorSignature() bool { return m.AuthorSignature != "" }

// IsEvent checks what current message is a any chat event.
func (m Message) IsEvent() bool {
	return m.IsChannelChatCreatedEvent() || m.IsDeleteChatPhotoEvent() || m.IsGroupChatCreatedEvent() ||
		m.IsLeftChatMemberEvent() || m.IsNewChatMembersEvent() || m.IsNewChatTitleEvent() ||
		m.IsSupergroupChatCreatedEvent() || m.IsNewChatPhotoEvent()
}

// IsBold checks that the current entity is a bold tag.
func (e MessageEntity) IsBold() bool { return strings.EqualFold(e.Type, EntityBold) }

// IsBotCommand checks that the current entity is a bot command.
func (e MessageEntity) IsBotCommand() bool { return strings.EqualFold(e.Type, EntityBotCommand) }

// IsCashtag checks that the current entity is a cashtag.
func (e MessageEntity) IsCashtag() bool { return strings.EqualFold(e.Type, EntityCashtag) }

// IsCode checks that the current entity is a code tag.
func (e MessageEntity) IsCode() bool { return strings.EqualFold(e.Type, EntityCode) }

// IsEmail checks that the current entity is a email.
func (e MessageEntity) IsEmail() bool { return strings.EqualFold(e.Type, EntityEmail) }

// IsHashtag checks that the current entity is a hashtag.
func (e MessageEntity) IsHashtag() bool { return strings.EqualFold(e.Type, EntityHashtag) }

// IsItalic checks that the current entity is a italic tag.
func (e MessageEntity) IsItalic() bool { return strings.EqualFold(e.Type, EntityItalic) }

// IsMention checks that the current entity is a username mention.
func (e MessageEntity) IsMention() bool { return strings.EqualFold(e.Type, EntityMention) }

// IsMPhoneNumberchecks that the current entity is a phone number.
func (e MessageEntity) IsPhoneNumber() bool { return strings.EqualFold(e.Type, EntityPhoneNumber) }

// IsPre checks that the current entity is a pre tag.
func (e MessageEntity) IsPre() bool { return strings.EqualFold(e.Type, EntityPre) }

// IsPre checks that the current entity is a pre tag.
func (e MessageEntity) IsStrikethrough() bool { return strings.EqualFold(e.Type, EntityStrikethrough) }

// IsTextLink checks that the current entity is a text link.
func (e MessageEntity) IsTextLink() bool { return strings.EqualFold(e.Type, EntityTextLink) }

// IsTextMention checks that the current entity is a mention without username.
func (e MessageEntity) IsTextMention() bool { return strings.EqualFold(e.Type, EntityTextMention) }

// IsUnderline checks that the current entity is a underline.
func (e MessageEntity) IsUnderline() bool { return strings.EqualFold(e.Type, EntityUnderline) }

// IsURL checks that the current entity is a URL.
func (e MessageEntity) IsURL() bool { return strings.EqualFold(e.Type, EntityURL) }

// ParseURL selects URL from message text/caption and parse it as fasthttp.URI.
func (e MessageEntity) ParseURL(text string) *http.URI {
	if !e.IsURL() || text == "" {
		return nil
	}

	link := http.AcquireURI()
	to := e.Offset + e.Length
	src := []rune(text)

	if len(src) < to {
		return nil
	}

	link.Update(string(src[e.Offset:to]))

	return link
}

// TextLink parse current text link entity as fasthttp.URI.
func (e MessageEntity) TextLink() *http.URI {
	if !e.IsTextLink() || e.URL == "" {
		return nil
	}

	link := http.AcquireURI()

	link.Update(e.URL)

	return link
}

func (a Audio) FullName(separator string) (name string) {
	if a.HasPerformer() {
		if separator == "" {
			separator = DefaultAudioSeparator
		}

		name += a.Performer + separator
	}

	title := DefaultAudioTitle
	if a.HasTitle() {
		title = a.Title
	}

	name += title

	return
}

func (a Audio) HasPerformer() bool { return a.Performer != "" }

func (a Audio) HasTitle() bool { return a.Title != "" }

func (a Audio) HasThumb() bool { return a.Thumb != nil }

// File returns File structure without FilePath parameter.
func (a Audio) File() File {
	return File{
		FileID:       a.FileID,
		FileSize:     a.FileSize,
		FileUniqueID: a.FileUniqueID,
	}
}

func (d Document) HasThumb() bool { return d.Thumb != nil }

func (d Document) File() File {
	return File{
		FileID:       d.FileID,
		FileUniqueID: d.FileUniqueID,
		FileSize:     d.FileSize,
	}
}

func (v Video) HasThumb() bool { return v.Thumb != nil }

func (v Video) File() File {
	return File{
		FileID:       v.FileID,
		FileUniqueID: v.FileUniqueID,
		FileSize:     v.FileSize,
	}
}

func (a Animation) HasThumb() bool { return a.Thumb != nil }

func (a Animation) File() File {
	return File{
		FileID:       a.FileID,
		FileUniqueID: a.FileUniqueID,
		FileSize:     a.FileSize,
	}
}

func (v Voice) File() File {
	return File{
		FileID:       v.FileID,
		FileUniqueID: v.FileUniqueID,
		FileSize:     v.FileSize,
	}
}

func (vn VideoNote) HasThumb() bool { return vn.Thumb != nil }

func (vn VideoNote) File() File {
	return File{
		FileID:       vn.FileID,
		FileUniqueID: vn.FileUniqueID,
		FileSize:     vn.FileSize,
	}
}

// FullName returns the full name of contact or FirstName if LastName is not available.
func (c Contact) FullName() string {
	if c.FirstName == "" {
		return ""
	}

	if c.HasLastName() {
		return c.FirstName + " " + c.LastName
	}

	return c.FirstName
}

// HaveLastName checks what the current contact has a LastName.
func (c Contact) HasLastName() bool { return c.LastName != "" }

func (c Contact) InTelegram() bool { return c.UserID != 0 }

func (c Contact) HasVCard() bool { return c.VCard != "" }

// VotesCount returns the total number of votes.
func (p Poll) VotesCount() int {
	var v int
	for i := range p.Options {
		v += p.Options[i].VoterCount
	}

	return v
}

func (p ChatPhoto) SmallFile() File {
	return File{
		FileID:       p.SmallFileID,
		FileUniqueID: p.SmallFileUniqueID,
	}
}

func (p ChatPhoto) BigFile() File {
	return File{
		FileID:       p.BigFileID,
		FileUniqueID: p.BigFileUniqueID,
	}
}

// IsAdministrator checks that current member is administrator.
func (m ChatMember) IsAdministrator() bool { return strings.EqualFold(m.Status, StatusAdministrator) }

// IsCreator checks that current member is creator.
func (m ChatMember) IsCreator() bool { return strings.EqualFold(m.Status, StatusCreator) }

// IsKicked checks that current member has been kicked.
func (m ChatMember) IsKicked() bool { return strings.EqualFold(m.Status, StatusKicked) }

// IsLeft checks that current member has left the chat.
func (m ChatMember) IsLeft() bool { return strings.EqualFold(m.Status, StatusLeft) }

// IsRestricted checks that current member has been restricted.
func (m ChatMember) IsRestricted() bool { return strings.EqualFold(m.Status, StatusRestricted) }

// UntilTime parse UntilDate of restrictions and returns time.Time.
func (m ChatMember) UntilTime() time.Time { return time.Unix(m.UntilDate, 0) }

func (m *InputMediaAnimation) GetMedia() *InputFile { return m.Media }

func (m *InputMediaAudio) GetMedia() *InputFile { return m.Media }

func (m *InputMediaDocument) GetMedia() *InputFile { return m.Media }

func (m *InputMediaPhoto) GetMedia() *InputFile { return m.Media }

func (InputMediaPhoto) isAlbumMedia() {}

func (m *InputMediaVideo) GetMedia() *InputFile { return m.Media }

func (InputMediaVideo) isAlbumMedia() {}

func (f *InputFile) IsFileID() bool { return f.ID != "" }

func (f *InputFile) IsURI() bool { return f.URI != nil }

func (f *InputFile) IsAttachment() bool { return f.Attachment != nil }

func (f *InputFile) MarshalJSON() ([]byte, error) {
	switch {
	case f.IsFileID():
		return []byte(f.ID), nil
	case f.IsURI():
		return f.URI.FullURI(), nil
	case f.IsAttachment():
		_, fileName := filepath.Split(f.Attachment.Name())

		u := http.AcquireURI()
		defer http.ReleaseURI(u)
		u.SetScheme(SchemeAttach)
		u.SetHost(fileName)
		u.SetPathBytes(nil)

		uri := u.FullURI() // NOTE(toby3d): remove slash on the end

		return uri[:len(uri)-1], nil
	default:
		return nil, nil
	}
}
