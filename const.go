package telegram

// Version represents current version of Telegram API supported by this package
const Version string = "5.2.0"

// Action represents available and supported status actions of bot
const (
	ActionFindLocation    string = "find_location"
	ActionRecordVoice     string = "record_voice"
	ActionRecordVideo     string = "record_video"
	ActionRecordVideoNote string = "record_video_note"
	ActionTyping          string = "typing"
	ActionUploadVoice     string = "upload_voice"
	ActionUploadDocument  string = "upload_document"
	ActionUploadPhoto     string = "upload_photo"
	ActionUploadVideo     string = "upload_video"
	ActionUploadVideoNote string = "upload_video_note"
)

// Chat represents available and supported chat types
const (
	ChatChannel    string = "channel"
	ChatGroup      string = "group"
	ChatPrivate    string = "private"
	ChatSuperGroup string = "supergroup"
)

// Command represents global commands which should be supported by any bot. You can user IsCommandEqual method of
// Message for checking.
//
// See: https://core.telegram.org/bots#global-commands
const (
	CommandHelp     string = "help"
	CommandSettings string = "settings"
	CommandStart    string = "start"
)

// Entity represents available and supported entity types
const (
	EntityBold          string = "bold"
	EntityBotCommand    string = "bot_command"
	EntityCashtag       string = "cashtag"
	EntityCode          string = "code"
	EntityEmail         string = "email"
	EntityHashtag       string = "hashtag"
	EntityItalic        string = "italic"
	EntityMention       string = "mention"
	EntityPhoneNumber   string = "phone_number"
	EntityPre           string = "pre"
	EntityStrikethrough string = "strikethrough"
	EntityTextLink      string = "text_link"
	EntityTextMention   string = "text_mention"
	EntityUnderline     string = "underline"
	EntityURL           string = "url"
)

// Method represents available and supported Telegram API methods
const (
	MethodAddStickerToSet                 string = "addStickerToSet"
	MethodAnswerCallbackQuery             string = "answerCallbackQuery"
	MethodAnswerInlineQuery               string = "answerInlineQuery"
	MethodAnswerPreCheckoutQuery          string = "answerPreCheckoutQuery"
	MethodAnswerShippingQuery             string = "answerShippingQuery"
	MethodClose                           string = "close"
	MethodCopyMessage                     string = "copyMessage"
	MethodCreateChatInviteLink            string = "createChatInviteLink"
	MethodCreateNewStickerSet             string = "createNewStickerSet"
	MethodDeleteChatPhoto                 string = "deleteChatPhoto"
	MethodDeleteChatStickerSet            string = "deleteChatStickerSet"
	MethodDeleteMessage                   string = "deleteMessage"
	MethodDeleteStickerFromSet            string = "deleteStickerFromSet"
	MethodDeleteWebhook                   string = "deleteWebhook"
	MethodEditChatInviteLink              string = "editChatInviteLink"
	MethodEditMessageCaption              string = "editMessageCaption"
	MethodEditMessageLiveLocation         string = "editMessageLiveLocation"
	MethodEditMessageMedia                string = "editMessageMedia"
	MethodEditMessageReplyMarkup          string = "editMessageReplyMarkup"
	MethodEditMessageText                 string = "editMessageText"
	MethodExportChatInviteLink            string = "exportChatInviteLink"
	MethodForwardMessage                  string = "forwardMessage"
	MethodGetChat                         string = "getChat"
	MethodGetChatAdministrators           string = "getChatAdministrators"
	MethodGetChatMember                   string = "getChatMember"
	MethodGetChatMembersCount             string = "getChatMembersCount"
	MethodGetFile                         string = "getFile"
	MethodGetGameHighScores               string = "getGameHighScores"
	MethodGetMe                           string = "getMe"
	MethodGetMyCommands                   string = "getMyCommands"
	MethodGetStickerSet                   string = "getStickerSet"
	MethodGetUpdates                      string = "getUpdates"
	MethodGetUserProfilePhotos            string = "getUserProfilePhotos"
	MethodGetWebhookInfo                  string = "getWebhookInfo"
	MethodKickChatMember                  string = "kickChatMember"
	MethodLeaveChat                       string = "leaveChat"
	MethodLogOut                          string = "logOut"
	MethodPinChatMessage                  string = "pinChatMessage"
	MethodPromoteChatMember               string = "promoteChatMember"
	MethodRestrictChatMember              string = "restrictChatMember"
	MethodRevokeChatInviteLink            string = "revokeChatInviteLink"
	MethodSendAnimation                   string = "sendAnimation"
	MethodSendAudio                       string = "sendAudio"
	MethodSendChatAction                  string = "sendChatAction"
	MethodSendContact                     string = "sendContact"
	MethodSendDice                        string = "sendDice"
	MethodSendDocument                    string = "sendDocument"
	MethodSendGame                        string = "sendGame"
	MethodSendInvoice                     string = "sendInvoice"
	MethodSendLocation                    string = "sendLocation"
	MethodSendMediaGroup                  string = "sendMediaGroup"
	MethodSendMessage                     string = "sendMessage"
	MethodSendPhoto                       string = "sendPhoto"
	MethodSendPoll                        string = "sendPoll"
	MethodSendSticker                     string = "sendSticker"
	MethodSendVenue                       string = "sendVenue"
	MethodSendVideo                       string = "sendVideo"
	MethodSendVideoNote                   string = "sendVideoNote"
	MethodSendVoice                       string = "sendVoice"
	MethodSetChatAdministratorCustomTitle string = "setChatAdministratorCustomTitle"
	MethodSetChatDescription              string = "setChatDescription"
	MethodSetChatPermissions              string = "setChatPermissions"
	MethodSetChatPhoto                    string = "setChatPhoto"
	MethodSetChatStickerSet               string = "setChatStickerSet"
	MethodSetChatTitle                    string = "setChatTitle"
	MethodSetGameScore                    string = "setGameScore"
	MethodSetMyCommands                   string = "setMyCommands"
	MethodSetPassportDataErrors           string = "setPassportDataErrors"
	MethodSetStickerPositionInSet         string = "setStickerPositionInSet"
	MethodSetStickerSetThumb              string = "setStickerSetThumb"
	MethodSetWebhook                      string = "setWebhook"
	MethodStopMessageLiveLocation         string = "stopMessageLiveLocation"
	MethodStopPoll                        string = "stopPoll"
	MethodUnbanChatMember                 string = "unbanChatMember"
	MethodUnpinAllChatMessages            string = "unpinAllChatMessages"
	MethodUnpinChatMessage                string = "unpinChatMessage"
	MethodUploadStickerFile               string = "uploadStickerFile"
)

// Mode represents available and supported parsing modes of messages
const (
	ParseModeHTML       string = "HTML"
	ParseModeMarkdown   string = "Markdown"
	ParseModeMarkdownV2 string = "MarkdownV2"
)

// Point represent a type of point on face
const (
	PointForehead string = "forehead"
	PointEyes     string = "eyes"
	PointMouth    string = "mouth"
	PointChin     string = "chin"
)

// Mime represents available and supported MIME types of data
const (
	MimeGIF  string = "image/gif"
	MimeHTML string = "text/html"
	MimeJPEG string = "image/jpeg"
	MimeMP4  string = "video/mp4"
	MimePDF  string = "application/pdf"
	MimeZIP  string = "application/zip"
)

// Scheme represents optional schemes for URLs
const (
	SchemeAttach   string = "attach"
	SchemeTelegram string = "tg"
)

// Status represents available and supported statuses of ID
const (
	StatusAdministrator string = "administrator"
	StatusCreator       string = "creator"
	StatusKicked        string = "kicked"
	StatusLeft          string = "left"
	StatusMember        string = "member"
	StatusRestricted    string = "restricted"
)

// Type represents available and supported types of data
const (
	TypeAddress               string = "address"
	TypeArticle               string = "article"
	TypeAudio                 string = "audio"
	TypeBankStatement         string = "bank_statement"
	TypeContact               string = "contact"
	TypeDocument              string = "document"
	TypeDriverLicense         string = "driver_license"
	TypeEmail                 string = "email"
	TypeGame                  string = "game"
	TypeGIF                   string = "gif"
	TypeIdentityCard          string = "identity_card"
	TypeInternalPassport      string = "internal_passport"
	TypeLocation              string = "location"
	TypeMpeg4Gif              string = "mpeg4_gif"
	TypePassport              string = "passport"
	TypePassportRegistration  string = "passport_registration"
	TypePersonalDetails       string = "personal_details"
	TypePhoneNumber           string = "phone_number"
	TypePhoto                 string = "photo"
	TypeRentalAgreement       string = "rental_agreement"
	TypeSticker               string = "sticker"
	TypeTemporaryRegistration string = "temporary_registration"
	TypeUtilityBill           string = "utility_bill"
	TypeVenue                 string = "venue"
	TypeVideo                 string = "video"
	TypeVoice                 string = "voice"
)

// Update represents available and supported types of updates
const (
	UpdateCallbackQuery      string = "callback_query"
	UpdateChannelPost        string = "channel_post"
	UpdateChosenInlineResult string = "chosen_inline_result"
	UpdateEditedChannelPost  string = "edited_channel_post"
	UpdateEditedMessage      string = "edited_message"
	UpdateInlineQuery        string = "inline_query"
	UpdateMessage            string = "message"
	UpdatePoll               string = "poll"
	UpdatePreCheckoutQuery   string = "pre_checkout_query"
	UpdateShippingQuery      string = "shipping_query"
)

// Default represents a default values for some helpers
const (
	DefaultAudioSeparator string = " – "
	DefaultAudioTitle     string = "[untitled]"
)

// Poll represents a poll types
const (
	PollQuiz    string = "quiz"
	PollRegular string = "regular"
)

// Emoji represents emoji supported by SendDice method
const (
	EmojiBasketball  string = "🏀" // 1-5
	EmojiBowling     string = "🎳" // 1-6
	EmojiDart        string = "🎯" // 1-6
	EmojiGameDie     string = "🎲" // 1-6
	EmojiSlotMachine string = "🎰" // 1-64
	EmojiSoccer      string = "⚽" // 1-5
)

const (
	// FromAnonymous is a User ID for messages from anonymous group administrators.
	FromAnonymous int64 = 1087968824 // @GroupAnonymousBot
	// FromForwarder is a User ID for messages automatically forwarded to the discussion group.
	FromForwarder int64 = 777000
)
