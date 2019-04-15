package telegram

import "encoding/json"

type (
	// Response represents a response from the Telegram API with the result
	// stored raw. If ok equals true, the request was successful, and the result
	// of the query can be found in the result field. In case of an unsuccessful
	// request, ok equals false, and the error is explained in the error field.
	Response struct {
		Ok          bool                `json:"ok"`
		ErrorCode   int                 `json:"error_code,omitempty"`
		Description string              `json:"description,omitempty"`
		Result      *json.RawMessage    `json:"result,omitempty"`
		Parameters  *ResponseParameters `json:"parameters,omitempty"`
	}

	// Update represents an incoming update.
	//
	// At most one of the optional parameters can be present in any given update.
	Update struct {
		// The update‘s unique identifier. Update identifiers start from a
		// certain positive number and increase sequentially. This ID becomes
		// especially handy if you’re using Webhooks, since it allows you to
		// ignore repeated updates or to restore the correct update sequence,
		// should they get out of order.
		ID int `json:"update_id"`

		// New incoming message of any kind — text, photo, sticker, etc.
		Message *Message `json:"message,omitempty"`

		// New version of a message that is known to the bot and was edited
		EditedMessage *Message `json:"edited_message,omitempty"`

		// New incoming channel post of any kind — text, photo, sticker, etc.
		ChannelPost *Message `json:"channel_post,omitempty"`

		// New version of a channel post that is known to the bot and was edited
		EditedChannelPost *Message `json:"adited_channel_post,omitempty"`

		// New incoming inline query
		InlineQuery *InlineQuery `json:"inline_query,omitempty"`

		// The result of an inline query that was chosen by a user and sent to
		// their chat partner.
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

		// New incoming callback query
		CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

		// New incoming shipping query. Only for invoices with flexible price
		ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

		// New incoming pre-checkout query. Contains full information about
		// checkout
		PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

		// New poll state. Bots receive only updates about polls, which are sent or stopped by the bot
		Poll *Poll `json:"poll,omitempty"`
	}

	// WebhookInfo contains information about the current status of a webhook.
	WebhookInfo struct {
		// Webhook URL, may be empty if webhook is not set up
		URL string `json:"url"`

		// Error message in human-readable format for the most recent error that
		// happened when trying to deliver an update via webhook
		LastErrorMessage string `json:"last_error_message,omitempty"`

		// True, if a custom certificate was provided for webhook certificate
		// checks
		HasCustomCertificate bool `json:"has_custom_certificate"`

		// Number of updates awaiting delivery
		PendingUpdateCount int `json:"pending_update_count"`

		// Maximum allowed number of simultaneous HTTPS connections to the
		// webhook for update delivery
		MaxConnections int `json:"max_connections,omitempty"`

		// Unix time for the most recent error that happened when trying to
		// deliver an update via webhook
		LastErrorDate int64 `json:"last_error_date,omitempty"`

		// A list of update types the bot is subscribed to. Defaults to all
		// update types
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

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

		// True if a group has ‘All Members Are Admins’ enabled.
		AllMembersAreAdministrators bool `json:"all_members_are_administrators,omitempty"`

		// Chat photo. Returned only in getChat.
		Photo *ChatPhoto `json:"photo,omitempty"`

		// Description, for supergroups and channel chats. Returned only in
		// getChat.
		Description string `json:"description,omitempty"`

		// Chat invite link, for supergroups and channel chats. Returned only in
		// getChat.
		InviteLink string `json:"invite_link,omitempty"`

		// Pinned message, for groups, supergroups and channels. Returned only in getChat.
		PinnedMessage *Message `json:"pinned_message,omitempty"`

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
		Entities []MessageEntity `json:"entities,omitempty"`

		// For messages with a caption, special entities like usernames, URLs,
		// bot commands, etc. that appear in the caption
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

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
		Photo []PhotoSize `json:"photo,omitempty"`

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
		Poll *Poll `json:"poll,omitempry"`

		// New members that were added to the group or supergroup and information
		// about them (the bot itself may be one of these members)
		NewChatMembers []User `json:"new_chat_members,omitempty"`

		// A member was removed from the group, information about them (this
		// member may be the bot itself)
		LeftChatMember *User `json:"left_chat_member,omitempty"`

		// A chat title was changed to this value
		NewChatTitle string `json:"new_chat_title,omitempty"`

		// A chat photo was change to this value
		NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

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
		// Unique identifier for this file
		FileID string `json:"file_id"`

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
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Performer of the audio as defined by sender or by audio tags
		Performer string `json:"performer,omitempty"`

		// Title of the audio as defined by sender or by audio tags
		Title string `json:"title,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// File size
		FileSize int `json:"file_size,omitempty"`

		// Thumbnail of the album cover to which the music file belongs
		Thumb *PhotoSize `json:"thumb,omitempty"`
	}

	// Document represents a general file (as opposed to photos, voice messages
	// and audio files).
	Document struct {
		// Unique file identifier
		FileID string `json:"file_id"`

		// Original filename as defined by sender
		FileName string `json:"file_name,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// Document thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// Video represents a video file.
	Video struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Mime type of a file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// Video width as defined by sender
		Width int `json:"width"`

		// Video height as defined by sender
		Height int `json:"height"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// File size
		FileSize int `json:"file_size,omitempty"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb,omitempty"`
	}

	// Voice represents a voice note.
	Voice struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}

	// VideoNote represents a video message (available in Telegram apps as of
	// v.4.0).
	VideoNote struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Video width and height (diameter of the video message) as defined by sender
		Length int `json:"length"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// File size
		FileSize int `json:"file_size,omitempty"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb,omitempty"`
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

		// Foursquare type of the venue. (For example,
		// "arts_entertainment/default", "arts_entertainment/aquarium" or
		// "food/icecream".)
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
		Options []PollOption `json:"options"`

		// True, if the poll is closed
		IsClosed bool `json:"is_closed"`
	}

	// UserProfilePhotos represent a user's profile pictures.
	UserProfilePhotos struct {
		// Total number of profile pictures the target user has
		TotalCount int `json:"total_count"`

		// Requested profile pictures (in up to 4 sizes each)
		Photos [][]PhotoSize `json:"photos"`
	}

	// File represents a file ready to be downloaded. The file can be downloaded
	// via the link https://api.telegram.org/file/bot<token>/<file_path>. It is
	// guaranteed that the link will be valid for at least 1 hour. When the link
	// expires, a new one can be requested by calling getFile.
	//
	// Maximum file size to download is 20 MB
	File struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// File path. Use https://api.telegram.org/file/bot<token>/<file_path> to
		// get the file.
		FilePath string `json:"file_path,omitempty"`

		// File size, if known
		FileSize int `json:"file_size,omitempty"`
	}

	// ReplyKeyboardMarkup represents a custom keyboard with reply options (see
	// Introduction to bots for details and examples).
	ReplyKeyboardMarkup struct {
		// Array of button rows, each represented by an Array of KeyboardButton
		// objects
		Keyboard [][]KeyboardButton `json:"keyboard"`

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
	// until a new keyboard is sent by a bot. An exception is made for one-time
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
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	}

	// InlineKeyboardButton represents one button of an inline keyboard. You
	// must use exactly one of the optional fields.
	InlineKeyboardButton struct {
		// Label text on the button
		Text string `json:"text"`

		// HTTP url to be opened when button is pressed
		URL string `json:"url,omitempty"`

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
	// optional parameters).
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
		// Unique file identifier of small (160x160) chat photo. This file_id can
		// be used only for photo download.
		SmallFileID string `json:"small_file_id"`

		// Unique file identifier of big (640x640) chat photo. This file_id can
		// be used only for photo download.
		BigFileID string `json:"big_file_id"`
	}

	// ChatMember contains information about one member of a chat.
	ChatMember struct {
		// Information about the user
		User *User `json:"user"`

		// The member's status in the chat. Can be "creator", "administrator",
		// "member", "restricted", "left" or "kicked"
		Status string `json:"status"`

		// Restictred and kicked only. Date when restrictions will be lifted for
		// this user, unix time
		UntilDate int64 `json:"until_date,omitempty"`

		// Administrators only. True, if the bot is allowed to edit administrator
		// privileges of that user
		CanBeEdited bool `json:"can_be_edited,omitempty"`

		// Administrators only. True, if the administrator can change the chat
		// title, photo and other settings
		CanChangeInfo bool `json:"can_change_info,omitempty"`

		// Administrators only. True, if the administrator can post in the
		// channel, channels only
		CanPostMessages bool `json:"can_post_messages,omitempty"`

		// Administrators only. True, if the administrator can edit messages of
		// other users, channels only
		CanEditMessages bool `json:"can_edit_messages,omitempty"`

		// Administrators only. True, if the administrator can delete messages of
		// other users
		CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

		// Administrators only. True, if the administrator can invite new users
		// to the chat
		CanInviteUsers bool `json:"can_invite_users,omitempty"`

		// Administrators only. True, if the administrator can restrict, ban or
		// unban chat members
		CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

		// Administrators only. True, if the administrator can pin messages,
		// supergroups only
		CanPinMessages bool `json:"can_pin_messages,omitempty"`

		// Administrators only. True, if the administrator can add new
		// administrators with a subset of his own privileges or demote
		// administrators that he has promoted, directly or indirectly (promoted
		// by administrators that were appointed by the user)
		CanPromoteMembers bool `json:"can_promote_members,omitempty"`

		// Restricted only. True, if the user is a member of the chat at the moment of the request
		IsMember bool `json:"is_member,omitempty"`

		// Restricted only. True, if the user can send text messages, contacts,
		// locations and venues
		CanSendMessages bool `json:"can_send_messages,omitempty"`

		// Restricted only. True, if the user can send audios, documents, photos,
		// videos, video notes and voice notes, implies can_send_messages
		CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

		// Restricted only. True, if the user can send animations, games,
		// stickers and use inline bots, implies can_send_media_messages
		CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

		// Restricted only. True, if user may add web page previews to his
		// messages, implies can_send_media_messages
		CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
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
		File() string
		InputMediaCaption() string
		InputMediaParseMode() string
		InputMediaType() string
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
		Media string `json:"media"`

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
		Media string `json:"media"`

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
		Media string `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and
		// less than 200 kB in size. A thumbnail‘s width and height should not
		// exceed 90. Ignored if the file is not uploaded using
		// multipart/form-data. Thumbnails can’t be reused and can be only
		// uploaded as a new file, so you can pass "attach://<file_attach_name>"
		// if the thumbnail was uploaded using multipart/form-data under
		// <file_attach_name>.
		Thumb InputFile `json:"thumb,omitempty"`

		// Caption of the animation to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
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

		// File to send. Pass a file_id to send a file that exists on the
		// Telegram servers (recommended), pass an HTTP URL for Telegram to get
		// a file from the Internet, or pass "attach://<file_attach_name>" to
		// upload a new one using multipart/form-data under <file_attach_name>
		// name.
		Media string `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and
		// less than 200 kB in size. A thumbnail‘s width and height should not
		// exceed 90. Ignored if the file is not uploaded using
		// multipart/form-data. Thumbnails can’t be reused and can be only
		// uploaded as a new file, so you can pass "attach://<file_attach_name>"
		// if the thumbnail was uploaded using multipart/form-data under
		// <file_attach_name>.
		Thumb InputFile `json:"thumb,omitempty"`

		// Caption of the audio to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
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

		// File to send. Pass a file_id to send a file that exists on the
		// Telegram servers (recommended), pass an HTTP URL for Telegram to get
		// a file from the Internet, or pass "attach://<file_attach_name>" to
		// upload a new one using multipart/form-data under <file_attach_name>
		// name.
		Media string `json:"media"`

		// Thumbnail of the file sent. The thumbnail should be in JPEG format and
		// less than 200 kB in size. A thumbnail‘s width and height should not
		// exceed 90. Ignored if the file is not uploaded using
		// multipart/form-data. Thumbnails can’t be reused and can be only
		// uploaded as a new file, so you can pass "attach://<file_attach_name>"
		// if the thumbnail was uploaded using multipart/form-data under
		// <file_attach_name>.
		Thumb InputFile `json:"thumb,omitempty"`

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption,omitempty"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic,
		// fixed-width text or inline URLs in the media caption.
		ParseMode string `json:"parse_mode,omitempty"`
	}

	// InputFile represents the contents of a file to be uploaded. Must be posted
	// using multipart/form-data in the usual way that files are uploaded via the
	// browser.
	InputFile interface{}

	// Animation provide an animation for your game so that it looks stylish in
	// chats (check out Lumberjack for an example). This object represents an
	// animation file to be displayed in the message containing a game.
	Animation struct {
		// Unique file identifier
		FileID string `json:"file_id"`

		// Original animation filename as defined by sender
		FileName string `json:"file_name,omitempty"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type,omitempty"`

		// Animation thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb,omitempty"`

		// File size
		FileSize int `json:"file_size,omitempty"`
	}
)
