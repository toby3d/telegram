package telegram

import (
	"encoding/json"
	"time"
)

type (
	// Response represents a response from the Telegram API with the result stored raw. If ok equals true, the request was successful, and the result of the query can be found in the result field. In case of an unsuccessful request, ok equals false, and the error is explained in the error field.
	Response struct {
		Ok          bool                `json:"ok"`
		ErrorCode   string              `json:"error_code"`
		Description string              `json:"description"`
		Result      *json.RawMessage    `json:"result"`
		Parameters  *ResponseParameters `json:"parameters"`
	}

	// Update represents an incoming update.
	//
	// At most one of the optional parameters can be present in any given update.
	Update struct {
		// The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
		ID int `json:"update_id"`

		// New incoming message of any kind — text, photo, sticker, etc.
		Message *Message `json:"message"` // optional

		// New version of a message that is known to the bot and was edited
		EditedMessage *Message `json:"edited_message"` // optional

		// New incoming channel post of any kind — text, photo, sticker, etc.
		ChannelPost *Message `json:"channel_post"` // optional

		// New version of a channel post that is known to the bot and was edited
		EditedChannelPost *Message `json:"adited_channel_post"` // optional

		// New incoming inline query
		InlineQuery *InlineQuery `json:"inline_query"` // optional

		// The result of an inline query that was chosen by a user and sent to their chat partner.
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` // optional

		// New incoming callback query
		CallbackQuery *CallbackQuery `json:"callback_query"` // optional

		// New incoming shipping query. Only for invoices with flexible price
		ShippingQuery *ShippingQuery `json:"shipping_query"` // optional

		// New incoming pre-checkout query. Contains full information about checkout
		PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query"` // optional
	}

	// WebhookInfo contains information about the current status of a webhook.
	WebhookInfo struct {
		// Webhook URL, may be empty if webhook is not set up
		URL string `json:"url"`

		// True, if a custom certificate was provided for webhook certificate checks
		HasCustomCertificate bool `json:"has_custom_certificate"`

		// Number of updates awaiting delivery
		PendingUpdateCount int `json:"pending_update_count"`

		//  Unix time for the most recent error that happened when trying to deliver an update via webhook
		LastErrorDate int `json:"last_error_date"` // optional

		// Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
		LastErrorMessage string `json:"last_error_message"` // optional

		// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
		MaxConnections int `json:"max_connections"` // optional

		// A list of update types the bot is subscribed to. Defaults to all update types
		AllowedUpdates []string `json:"allowed_updates"` // optional
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
		LastName string `json:"last_name"` // optional

		// User‘s or bot’s username
		Username string `json:"username"` // optional

		// IETF language tag of the user's language
		LanguageCode string `json:"language_code"` // optional
	}

	// Chat represents a chat.
	Chat struct {
		// Unique identifier for this chat.
		ID int64 `json:"id"`

		// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
		Type string `json:"type"`

		//Title, for supergroups, channels and group chats
		Title string `json:"title"` // optional

		// Username, for private chats, supergroups and channels if available
		Username string `json:"username"` // optional

		// First name of the other party in a private chat
		FirstName string `json:"first_name"` // optional

		// Last name of the other party in a private chat
		LastName string `json:"last_name"` // optional

		// True if a group has ‘All Members Are Admins’ enabled.
		AllMembersAreAdministrators bool `json:"all_members_are_administrators"` // optional

		// Chat photo. Returned only in getChat.
		Photo *ChatPhoto `json:"photo"` // optional

		// Description, for supergroups and channel chats. Returned only in getChat.
		Description string `json:"description"` // optional

		// Chat invite link, for supergroups and channel chats. Returned only in getChat.
		InviteLink string `json:"invite_link"` // optional

		// Pinned message, for supergroups. Returned only in getChat.
		PinnedMessage *Message `json:"pinned_message"` // optional
	}

	// Message represents a message.
	Message struct {
		ID int `json:"message_id"` // Unique message identifier inside this chat

		// Sender, empty for messages sent to channels
		From *User `json:"from"` // optional

		// Date the message was sent in Unix time
		Date int `json:"date"`

		// Conversation the message belongs to
		Chat *Chat `json:"chat"`

		// For forwarded messages, sender of the original message
		ForwardFrom *User `json:"forward_from"` // optional

		// For messages forwarded from channels, information about the original channel
		ForwardFromChat *Chat `json:"forward_from_chat"` // optional

		// For messages forwarded from channels, identifier of the original message in the channel
		ForwardFromMessageID int `json:"forward_from_message_id"` // optional

		// For messages forwarded from channels, signature of the post author if present
		ForwardSignature string `json:"forward_signature"` // optional

		// For forwarded messages, date the original message was sent in Unix time
		ForwardDate int `json:"forward_date"` // optional

		// For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
		ReplyToMessage *Message `json:"reply_to_message"` // optional

		// Date the message was last edited in Unix time
		EditDate int `json:"edit_date"` // optional

		// Signature of the post author for messages in channels
		AuthorSignature string `json:"author_signature"` // optional

		// For text messages, the actual UTF-8 text of the message, 0-4096 characters.
		Text string `json:"text"` // optional

		// For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
		Entities *[]MessageEntity `json:"entities"` // optional

		// Message is an audio file, information about the file
		Autdio *Audio `json:"audio"` // optional

		// Message is a general file, information about the file
		Document *Document `json:"document"` // optional

		// Message is a game, information about the game. More about games »
		Game *Game `json:"game"` // optional

		// Message is a photo, available sizes of the photo
		Photo *[]PhotoSize `json:"photo"` // optional

		// Message is a sticker, information about the sticker
		Sticker *Sticker `json:"sticker"` // optional

		// Message is a video, information about the video
		Video *Video `json:"video"` // optional

		// Message is a voice message, information about the file
		Voice *Voice `json:"voice"` // optional

		// Message is a video note, information about the video message
		VideoNote *VideoNote `json:"video_note"` // optional

		// New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
		NewChatMembers *[]User `json:"new_chat_members"` // optional

		// Caption for the document, photo or video, 0-200 characters
		Caption string `json:"caption"` // optional

		// Message is a shared contact, information about the contact
		Contact *Contact `json:"contact"` // optional

		// Message is a shared location, information about the location
		Location *Location `json:"location"` // optional

		// Message is a venue, information about the venue
		Venue *Venue `json:"venue"` // optional

		// A new member was added to the group, information about them (this member may be the bot itself)
		NewChatMember *User `json:"new_chat_member"` // optional

		// A member was removed from the group, information about them (this member may be the bot itself)
		LeftChatMember *User `json:"left_chat_member"` // optional

		// A chat title was changed to this value
		NewChatTitle string `json:"new_chat_title"` // optional

		// A chat photo was change to this value
		NewChatPhoto *[]PhotoSize `json:"new_chat_photo"` // optional

		// Service message: the chat photo was deleted
		DeleteChatPhoto bool `json:"delete_chat_photo"` // optional

		// Service message: the group has been created
		GroupChatCreated bool `json:"group_chat_created"` // optional

		// Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
		SupergroupChatCreated bool `json:"supergroup_chat_created"` // optional

		// Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
		ChannelChatCreated bool `json:"channel_chat_created"` // optional

		// The group has been migrated to a supergroup with the specified identifier.
		MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional

		// The supergroup has been migrated from a group with the specified identifier.
		MigrateFromChatID int64 `json:"migrate_from_chat_id"` // optional

		// Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
		PinnedMessage *Message `json:"pinned_message"` // optional

		// Message is an invoice for a payment, information about the invoice. More about payments »
		Invoice *Invoice `json:"invoice"` // optional

		// Message is a service message about a successful payment, information about the payment. More about payments »
		SuccessfulPayment *SuccessfulPayment `json:"successful_payment"` // optional
	}

	// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
	MessageEntity struct {
		// Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
		Type string `json:"type"`

		// Offset in UTF-16 code units to the start of the entity
		Offset int `json:"offset"`

		// Length of the entity in UTF-16 code units
		Length int `json:"length"`

		// For “text_link” only, url that will be opened after user taps on the text
		Url string `json:"url"` // optional

		// For “text_mention” only, the mentioned user
		User *User `json:"user"` // optional
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
		FileSize int `json:"file_size"` // optional
	}

	// Audio represents an audio file to be treated as music by the Telegram clients.
	Audio struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// Performer of the audio as defined by sender or by audio tags
		Performer string `json:"performer"` // optional

		// Title of the audio as defined by sender or by audio tags
		Title string `json:"title"` // optional

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// Document represents a general file (as opposed to photos, voice messages and audio files).
	Document struct {
		// Unique file identifier
		FileID string `json:"file_id"`

		// Document thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb"` // optional

		// Original filename as defined by sender
		FileName string `json:"file_name"` // optional

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// Video represents a video file.
	Video struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Video width as defined by sender
		Width int `json:"width"`

		// Video height as defined by sender
		Height int `json:"height"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb"` // optional

		// Mime type of a file as defined by sender
		MimeType string `json:"mime_type"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// Voice represents a voice note.
	Voice struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Duration of the audio in seconds as defined by sender
		Duration int `json:"duration"`

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// VideoNote represents a video message (available in Telegram apps as of v.4.0).
	VideoNote struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Video width and height as defined by sender
		Length int `json:"length"`

		// Duration of the video in seconds as defined by sender
		Duration int `json:"duration"`

		// Video thumbnail
		Thumb *PhotoSize `json:"thumb"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// Contact represents a phone contact.
	Contact struct {
		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name"` // optional

		// Contact's user identifier in Telegram
		UserID int `json:"user_id"` // optional
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
		FoursquareID string `json:"foursquare_id"` // optional
	}

	// UserProfilePhotos represent a user's profile pictures.
	UserProfilePhotos struct {
		// Total number of profile pictures the target user has
		TotalCount int `json:"total_count"`

		// Requested profile pictures (in up to 4 sizes each)
		Photos []*[]PhotoSize `json:"photos"`
	}

	// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
	//
	// Maximum file size to download is 20 MB
	File struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// File size, if known
		FileSize int `json:"file_size"` // optional

		// File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
		FilePath string `json:"file_path"` // optional
	}

	// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
	ReplyKeyboardMarkup struct {
		// Array of button rows, each represented by an Array of KeyboardButton objects
		keyboard [][]KeyboardButton `json:"keyboard"`

		// Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
		resize_keyboard bool `json:"resize_keyboard"` // optional

		// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
		one_time_keyboard bool `json:"one_time_keyboard"` // optional

		// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
		//
		// Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
		selective bool `json:"selective"` // optional
	}

	// KeyboardButton represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
	KeyboardButton struct {
		// Text of the button. If none of the optional fields are used, it will be sent to the bot as a message when the button is pressed
		Text string `json:"text"`

		// If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
		RequestContact bool `json:"request_contact"` // optional

		// If True, the user's current location will be sent when the button is pressed. Available in private chats only
		RequestLocation bool `json:"request_location"` // optional
	}

	// ReplyKeyboardRemove will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
	ReplyKeyboardRemove struct {
		// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
		SemoveKeyboard bool `json:"remove_keyboard"`

		// Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
		//
		// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
		Selective bool `json:"selective"` // optional
	}

	// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
	InlineKeyboardMarkup struct {
		// Array of button rows, each represented by an Array of InlineKeyboardButton objects
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	}

	// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
	InlineKeyboardButton struct {
		// Label text on the button
		Text string `json:"text"`

		// HTTP url to be opened when button is pressed
		URL string `json:"url"` // optional

		// Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
		CallbackData string `json:"callback_data"` // optional

		// If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
		//
		// Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
		SwitchInlineQuery string `json:"switch_inline_query"` // optional

		// If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.
		//
		// This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
		SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"` // optional

		// Description of the game that will be launched when the user presses the button.
		//
		// NOTE: This type of button must always be the first button in the first row.
		CallbackGame *CallbackGame `json:"callback_game"` // optional

		// Specify True, to send a Pay button.
		//
		// NOTE: This type of button must always be the first button in the first row.
		Pay bool `json:"pay"` // optional
	}

	// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
	//
	// NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any of the optional parameters).
	CallbackQuery struct {
		// Unique identifier for this query
		ID string `json:"id"`

		// Sender
		From *User `json:"from"`

		// Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
		Message *Message `json:"message"` // optional

		// Identifier of the message sent via the bot in inline mode, that originated the query.
		InlineMessageID string `json:"inline_message_id"` // optional

		// Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
		ChatInstance string `json:"chat_instance"`

		// Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
		Data string `json:"data"` // optional

		// Short name of a Game to be returned, serves as the unique identifier for the game
		GameShortName string `json:"game_short_name"` // optional
	}

	// ForceReply display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
	ForceReply struct {
		// Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
		ForceReply bool `json:"force_reply"`

		// Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
		Selective bool `json:"selective"` // optional
	}

	// ChatPhoto represents a chat photo.
	ChatPhoto struct {
		// Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
		SmallFileID string `json:"small_file_id"`

		// Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
		BigFileID string `json:"big_file_id"`
	}

	// ChatMember contains information about one member of a chat.
	ChatMember struct {
		// Information about the user
		User *User `json:"user"`

		// The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
		Status string `json:"status"`

		// Restictred and kicked only. Date when restrictions will be lifted for this user, unix time
		UntilDate int `json:"until_date"` // optional

		// Administrators only. True, if the bot is allowed to edit administrator privileges of that user
		CanBeEdited bool `json:"can_be_edited"` // optional

		// Administrators only. True, if the administrator can change the chat title, photo and other settings
		CanChangeInfo bool `json:"can_change_info"` // optional

		// Administrators only. True, if the administrator can post in the channel, channels only
		CanPostMessages bool `json:"can_post_messages"` // optional

		// Administrators only. True, if the administrator can edit messages of other users, channels only
		CanEditMessages bool `json:"can_edit_messages"` // optional

		// Administrators only. True, if the administrator can delete messages of other users
		CanDeleteMessages bool `json:"can_delete_messages"` // optional

		// Administrators only. True, if the administrator can invite new users to the chat
		CanInviteUsers bool `json:"can_invite_users"` // optional

		// Administrators only. True, if the administrator can restrict, ban or unban chat members
		CanRestrictMembers bool `json:"can_restrict_members"` // optional

		// Administrators only. True, if the administrator can pin messages, supergroups only
		CanPinMessages bool `json:"can_pin_messages"` // optional

		// Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
		CanPromoteMembers bool `json:"can_promote_members"` // optional

		// Restricted only. True, if the user can send text messages, contacts, locations and venues
		CanSendMessages bool `json:"can_send_messages"` // optional

		// Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
		CanSendMediaMessages bool `json:"can_send_media_messages"` // optional

		// Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
		CanSendOtherMessages bool `json:"can_send_other_messages"` // optional

		// Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
		CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // optional
	}

	// ResponseParameters contains information about why a request was unsuccessful.
	ResponseParameters struct {
		// The group has been migrated to a supergroup with the specified identifier.
		MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional

		// In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
		RetryAfter int `json:"retry_after"` // optional
	}

	// InputFile represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
	InputFile interface{}

	// Sticker represents a sticker.
	Sticker struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// Sticker width
		Width int `json:"width"`

		// Sticker height
		Height int `json:"height"`

		// Sticker thumbnail in the .webp or .jpg format
		Thumb *PhotoSize `json:"thumb"` // optional

		// Emoji associated with the sticker
		Emoji string `json:"emoji"` // optional

		// Name of the sticker set to which the sticker belongs
		SetName string `json:"set_name"` // optional

		// For mask stickers, the position where the mask should be placed
		MaskPosition *MaskPosition `json:"mask_position"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// StickerSet represents a sticker set.
	StickerSet struct {
		// Sticker set name
		Name string `json:"name"`

		// Sticker set title
		Title string `json:"title"`

		// True, if the sticker set contains masks
		ContainsMasks bool `json:"contains_masks"`

		// List of all set stickers
		Stickers *[]Sticker `json:"stickers"`
	}

	// MaskPosition describes the position on faces where a mask should be placed by default.
	MaskPosition struct {
		// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
		Point string `json:"point"`

		// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
		XShift float32 `json:"x_shift"`

		// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
		YShift float32 `json:"y_shift"`

		// Mask scaling coefficient. For example, 2.0 means double size.
		Scale float32 `json:"scale"`
	}

	// InlineQuery represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
	InlineQuery struct {
		// Unique identifier for this query
		ID string `json:"id"`

		// Sender
		From *User `json:"from"`

		// Sender location, only for bots that request user location
		Location *Location `json:"location"` // optional

		// Text of the query (up to 512 characters)
		Query string `json:"query"`

		// Offset of the results to be returned, can be controlled by the bot
		Offset string `json:"offset"`
	}

	// InlineQueryResultArticle represents a link to an article or web page.
	InlineQueryResultArticle struct {
		// Type of the result, must be article
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		//Title of the result
		Title string `json:"title"`

		// Content of the message to be sent
		InputMessageContent *InputMessageContent `json:"input_message_content"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// URL of the result
		URL string `json:"url"` // optional

		// Pass True, if you don't want the URL to be shown in the message
		HideURL bool `json:"hide_url"` // optional

		// Short description of the result
		Description string `json:"description"` // optional

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url"` // optional

		// Thumbnail width
		ThumbWidth int `json:"thumb_width"` // optional

		// Thumbnail height
		ThumbHeight int `json:"thumb_height"` // optional
	}

	// InlineQueryResultPhoto represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
	InlineQueryResultPhoto struct {
		// Type of the result, must be photo
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
		PhotoURL string `json:"photo_url"`

		// URL of the thumbnail for the photo
		thumb_url string `json:"thumb_url"`

		// Width of the photo
		photo_width int `json:"photo_width"` // optional

		// Height of the photo
		photo_height int `json:"photo_height"` // optional

		// Title for the result
		Title string `json:"title"` // optional

		// Short description of the result
		Description string `json:"description"` // optional

		// Caption of the photo to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the photo
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultGif represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
	InlineQueryResultGif struct {
		// Type of the result, must be gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the GIF file. File size must not exceed 1MB
		GifURL string `json:"gif_url"`

		// Width of the GIF
		GifWidth int `json:"gif_width"` // optional

		// Height of the GIF
		GifHeight int `json:"gif_height"` // optional

		// Duration of the GIF
		GifDuration int `json:"gif_duration"` // optional

		// URL of the static thumbnail for the result (jpeg or gif)
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title"` // optional

		// Caption of the GIF file to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the GIF animation
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
	InlineQueryResultMpeg4Gif struct {
		// Type of the result, must be mpeg4_gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the MP4 file. File size must not exceed 1MB
		Mpeg4URL string `json:"mpeg4_url"`

		// Video width
		Mpeg4Width int `json:"mpeg4_width"` // optional

		// Video height
		Mpeg4Height int `json:"mpeg4_height"` // optional

		// Video duration
		Mpeg4Duration int `json:"mpeg4_duration"` // optional

		// URL of the static thumbnail (jpeg or gif) for the result
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title"` // optional

		// Caption of the MPEG-4 file to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the video animation
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
	//
	// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using input_message_content.
	InlineQueryResultVideo struct {
		// Type of the result, must be video
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid URL for the embedded video player or video file
		VideoURL string `json:"video_url"`

		// Mime type of the content of video url, “text/html” or “video/mp4”
		MimeType string `json:"mime_type"`

		// URL of the thumbnail (jpeg only) for the video
		ThumbURL string `json:"thumb_url"`

		// Title for the result
		Title string `json:"title"`

		// Caption of the video to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Video width
		VideoWidth int `json:"video_width"` // optional

		// Video height
		VideoHeight int `json:"video_height"` // optional

		// Video duration in seconds
		VideoDuration int `json:"video_duration"` // optional

		// Short description of the result
		Description string `json:"description"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultAudio represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
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
		Caption string `json:"caption"` // optional

		// Performer
		Performer string `json:"performer"` // optional

		// Audio duration in seconds
		AudioDuration int `json:"audio_duration"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the audio
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultVoice represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
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
		Caption string `json:"caption"` // optional

		// Recording duration in seconds
		VoiceDuration int `json:"voice_duration"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the voice recording
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultDocument represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
	InlineQueryResultDocument struct {
		// Type of the result, must be document
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// Title for the result
		Title string `json:"title"`

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// A valid URL for the file
		DocumentURL string `json:"document_url"`

		// Mime type of the content of the file, either “application/pdf” or “application/zip”
		MimeType string `json:"mime_type"`

		// Short description of the result
		Description string `json:"description"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the file
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional

		// URL of the thumbnail (jpeg only) for the file
		ThumbURL string `json:"thumb_url"` // optional

		// Thumbnail width
		ThumbWidth int `json:"thumb_width"` // optional

		// Thumbnail height
		ThumbHeight int `json:"thumb_height"` // optional
	}

	// InlineQueryResultLocation represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
	InlineQueryResultLocation struct {
		// Type of the result, must be location
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Location latitude in degrees
		Latitude float32 `json:"latitude"`

		// Location longitude in degrees
		Longitude float32 `json:"longitude"`

		// Location title
		Title string `json:"title"`

		//Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		//Content of the message to be sent instead of the location
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional

		//Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url"` // optional

		//Thumbnail width
		ThumbWidth int `json:"thumb_width"` // optional

		//Thumbnail height
		ThumbHeight int `json:"thumb_height"` // optional
	}

	// InlineQueryResultVenue represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
	InlineQueryResultVenue struct {
		// Type of the result, must be venue
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 Bytes
		ID string `json:"id"`

		// Latitude of the venue location in degrees
		Latitude float32 `json:"latitude"`

		// Longitude of the venue location in degrees
		Longitude float32 `json:"longitude"`

		// Title of the venue
		Title string `json:"title"`

		// Address of the venue
		Address string `json:"address"`

		// Foursquare identifier of the venue if known
		FoursquareID string `json:"foursquare_id"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the venue
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url"` // optional

		// Thumbnail width
		ThumbWidth int `json:"thumb_width"` // optional

		// Thumbnail height
		ThumbHeight int `json:"thumb_height"` // optional
	}

	// InlineQueryResultContact represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
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
		LastName string `json:"last_name"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the contact
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional

		// Url of the thumbnail for the result
		ThumbURL string `json:"thumb_url"` // optional

		// Thumbnail width
		ThumbWidth int `json:"thumb_width"` // optional

		// Thumbnail height
		ThumbHeight int `json:"thumb_height"` // optional
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
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional
	}

	// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
	InlineQueryResultCachedPhoto struct {
		// Type of the result, must be photo
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier of the photo
		PhotoFileID string `json:"photo_file_id"`

		// Title for the result
		Title string `json:"title"` // optional

		// Short description of the result
		Description string `json:"description"` // optional

		// Caption of the photo to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the photo
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedGif represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
	InlineQueryResultCachedGif struct {
		// Type of the result, must be gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the GIF file
		GifFileID string `json:"gif_file_id"`

		// Title for the result
		Title string `json:"title"` // optional

		// Caption of the GIF file to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the GIF animation
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
	InlineQueryResultCachedMpeg4Gif struct {
		// Type of the result, must be mpeg4_gif
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the MP4 file
		Mpeg4FileID string `json:"mpeg4_file_id"`

		// Title for the result
		Title string `json:"title"` // optional

		// Caption of the MPEG-4 file to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the video animation
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
	InlineQueryResultCachedSticker struct {
		// Type of the result, must be sticker
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier of the sticker
		StickerFileID string `json:"sticker_file_id"`

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the sticker
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
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
		Description string `json:"description"` // optional

		// Caption of the document to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the file
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
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
		Description string `json:"description"` // optional

		// Caption of the video to be sent, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the video
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
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
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the voice message
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	// InlineQueryResultCachedAudio represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
	InlineQueryResultCachedAudio struct {
		// Type of the result, must be audio
		Type string `json:"type"`

		// Unique identifier for this result, 1-64 bytes
		ID string `json:"id"`

		// A valid file identifier for the audio file
		AudioFileID string `json:"audio_file_id"`

		// Caption, 0-200 characters
		Caption string `json:"caption"` // optional

		// Inline keyboard attached to the message
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // optional

		// Content of the message to be sent instead of the audio
		InputMessageContent *InputMessageContent `json:"input_message_content"` // optional
	}

	InputMessageContent interface{}

	// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
	InputTextMessageContent struct {
		// Text of the message to be sent, 1-4096 characters
		MessageText string `json:"message_text"`

		// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
		ParseMode string `json:"parse_mode"` // optional

		// Disables link previews for links in the sent message
		DisableWebPagePreview bool `json:"disable_web_page_preview"` // optional
	}

	// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
	InputLocationMessageContent struct {
		// Latitude of the location in degrees
		Latitude float32 `json:"latitude"`

		// Longitude of the location in degrees
		Longitude float32 `json:"longitude"`
	}

	// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
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
		FoursquareID string `json:"foursquare_id"` // optional
	}

	// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
	InputContactMessageContent struct {
		// Contact's phone number
		PhoneNumber string `json:"phone_number"`

		// Contact's first name
		FirstName string `json:"first_name"`

		// Contact's last name
		LastName string `json:"last_name"` // optional
	}

	// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
	ChosenInlineResult struct {
		// The unique identifier for the result that was chosen
		ResultID string `json:"result_id"`

		// The user that chose the result
		From *User `json:"from"`

		// Sender location, only for bots that require user location
		Location *Location `json:"location"` // optional

		// Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
		InlineMessageID string `json:"inline_message_id"` // optional

		// The query that was used to obtain the result
		Query string `json:"query"`
	}

	/* ==== ====== */

	// LabeledPrice represents a portion of the price for goods or services.
	LabeledPrice struct {
		// Portion label
		Label string `json:"label"`

		//	Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		Amount int `json:"amount"`
	}

	// Invoice contains basic information about an invoice.
	Invoice struct {
		// Product name
		Title string `json:"title"`

		// Product description
		Description string `json:"description"`

		// Unique bot deep-linking parameter that can be used to generate this invoice
		StartParameter string `json:"start_parameter"`

		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`
	}

	// ShippingAddress represents a shipping address.
	ShippingAddress struct {
		// ISO 3166-1 alpha-2 country code
		CountryCode string `json:"country_code"`

		// State, if applicable
		State string `json:"state"`

		// City
		City string `json:"city"`

		// First line for the address
		StreetLine1 string `json:"street_line1"`

		// Second line for the address
		StreetLine2 string `json:"street_line2"`

		// Address post code
		PostCode string `json:"post_code"`
	}

	// OrderInfo represents information about an order.
	OrderInfo struct {
		// User name
		Name string `json:"name"` // optional

		// User's phone number
		PhoneNumber string `json:"phone_number"` // optional

		// User email
		Email string `json:"email"` // optional

		// User shipping address
		ShippingAddress *ShippingAddress // optional
	}

	// ShippingOption represents one shipping option.
	ShippingOption struct {
		// Shipping option identifier
		ID string `json:"id"`

		// Option title
		Title string `json:"title"`

		// List of price portions
		Prices *[]LabeledPrice `json:"prices"`
	}

	// SuccessfulPayment contains basic information about a successful payment.
	SuccessfulPayment struct {
		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// Identifier of the shipping option chosen by the user
		ShippingOptionID string `json:"shipping_option_id"` // optional

		// Order info provided by the user
		OrderInfo *OrderInfo `json:"order_info"` // optional

		// Telegram payment identifier
		TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

		// Provider payment identifier
		ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
	}

	// ShippingQuery contains information about an incoming shipping query.
	ShippingQuery struct {
		// Unique query identifier
		ID string `json:"id"`

		// User who sent the query
		From *User `json:"from"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// User specified shipping address
		ShippingAddress *ShippingAddress `json:"shipping_address"`
	}

	// PreCheckoutQuery contains information about an incoming pre-checkout query.
	PreCheckoutQuery struct {
		// Unique query identifier
		ID string `json:"id"`

		// User who sent the query
		From *User `json:"from"`

		// Three-letter ISO 4217 currency code
		Currency string `json:"currency"`

		// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
		TotalAmount int `json:"total_amount"`

		// Bot specified invoice payload
		InvoicePayload string `json:"invoice_payload"`

		// Identifier of the shipping option chosen by the user
		ShippingOptionID string `json:"shipping_option_id"` // optional

		// Order info provided by the user
		OrderInfo *OrderInfo `json:"order_info"` // optional
	}

	// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
	Game struct {
		// Title of the game
		Title string `json:"title"`

		// Description of the game
		Description string `json:"description"`

		// Photo that will be displayed in the game message in chats.
		Photo *[]PhotoSize `json:"photo"`

		// Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
		Text string `json:"text"` // optional

		// Special entities that appear in text, such as usernames, URLs, bot commands, etc.
		TextEntities *[]MessageEntity `json:"text_entities"` // optional

		// Animation that will be displayed in the game message in chats. Upload via BotFather
		Animation *Animation `json:"animation"` // optional
	}

	// Animation provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
	Animation struct {
		// Unique file identifier
		FileID string `json:"file_id"`

		// Animation thumbnail as defined by sender
		Thumb *PhotoSize `json:"thumb"` // optional

		// Original animation filename as defined by sender
		FileName string `json:"file_name"` // optional

		// MIME type of the file as defined by sender
		MimeType string `json:"mime_type"` // optional

		// File size
		FileSize int `json:"file_size"` // optional
	}

	// CallbackGame a placeholder, currently holds no information. Use BotFather to set up your game.
	CallbackGame struct{}

	// GameHighScore represents one row of the high scores table for a game.
	GameHighScore struct {
		// Position in high score table for the game
		Position int `json:"position"`

		// User
		User *User `json:"user"`

		// Score
		Score int `json:"score"`
	}
)

func (msg *Message) Time() time.Time {
	return time.Unix(int64(msg.Date), 0)
}
