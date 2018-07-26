package telegram

// Version represents current version of Telegram API supported by this package
const Version = 4.0

// Action... represents available and supported status actions of bot
const (
	ActionFindLocation    = "find_location"
	ActionRecordAudio     = "record_audio"
	ActionRecordVideo     = "record_video"
	ActionRecordVideoNote = "record_video_note"
	ActionTyping          = "typing"
	ActionUploadAudio     = "upload_audio"
	ActionUploadDocument  = "upload_document"
	ActionUploadPhoto     = "upload_photo"
	ActionUploadVideo     = "upload_video"
	ActionUploadVideoNote = "upload_video_note"
)

// Chat... represents available and supported chat types
const (
	ChatChannel    = "channel"
	ChatGroup      = "group"
	ChatPrivate    = "private"
	ChatSuperGroup = "supergroup"
)

// Command... represents global commands which should be supported by any bot.
// You can user IsCommandEqual method of Message for checking.
//
// See: https://core.telegram.org/bots#global-commands
const (
	CommandStart    = "start"
	CommandHelp     = "help"
	CommandSettings = "settings"
)

// Entity... represents available and supported entity types
const (
	EntityBold        = "bold"
	EntityBotCommand  = "bot_command"
	EntityCashtag     = "cashtag"
	EntityCode        = "code"
	EntityEmail       = "email"
	EntityHashtag     = "hashtag"
	EntityItalic      = "italic"
	EntityMention     = "mention"
	EntityPhoneNumber = "phone_number"
	EntityPre         = "pre"
	EntityTextLink    = "text_link"
	EntityTextMention = "text_mention"
	EntityURL         = "url"
)

// Method... represents available and supported Telegram API methods
const (
	MethodAddStickerToSet         = "addStickerToSet"
	MethodAnswerCallbackQuery     = "answerCallbackQuery"
	MethodAnswerInlineQuery       = "answerInlineQuery"
	MethodAnswerPreCheckoutQuery  = "answerPreCheckoutQuery"
	MethodAnswerShippingQuery     = "answerShippingQuery"
	MethodCreateNewStickerSet     = "createNewStickerSet"
	MethodDeleteChatPhoto         = "deleteChatPhoto"
	MethodDeleteChatStickerSet    = "deleteChatStickerSet"
	MethodDeleteMessage           = "deleteMessage"
	MethodDeleteStickerFromSet    = "deleteStickerFromSet"
	MethodDeleteWebhook           = "deleteWebhook"
	MethodEditMessageCaption      = "editMessageCaption"
	MethodEditMessageLiveLocation = "editMessageLiveLocation"
	MethodEditMessageMedia        = "editMessageMedia"
	MethodEditMessageReplyMarkup  = "editMessageReplyMarkup"
	MethodEditMessageText         = "editMessageText"
	MethodExportChatInviteLink    = "exportChatInviteLink"
	MethodForwardMessage          = "forwardMessage"
	MethodGetChat                 = "getChat"
	MethodGetChatAdministrators   = "getChatAdministrators"
	MethodGetChatMember           = "getChatMember"
	MethodGetChatMembersCount     = "getChatMembersCount"
	MethodGetFile                 = "getFile"
	MethodGetGameHighScores       = "getGameHighScores"
	MethodGetMe                   = "getMe"
	MethodGetStickerSet           = "getStickerSet"
	MethodGetUpdates              = "getUpdates"
	MethodGetUserProfilePhotos    = "getUserProfilePhotos"
	MethodGetWebhookInfo          = "getWebhookInfo"
	MethodKickChatMember          = "kickChatMember"
	MethodLeaveChat               = "leaveChat"
	MethodPinChatMessage          = "pinChatMessage"
	MethodPromoteChatMember       = "promoteChatMember"
	MethodRestrictChatMember      = "restrictChatMember"
	MethodSendAnimation           = "sendAnimation"
	MethodSendAudio               = "sendAudio"
	MethodSendChatAction          = "sendChatAction"
	MethodSendContact             = "sendContact"
	MethodSendDocument            = "sendDocument"
	MethodSendGame                = "sendGame"
	MethodSendInvoice             = "sendInvoice"
	MethodSendLocation            = "sendLocation"
	MethodSendMediaGroup          = "sendMediaGroup"
	MethodSendMessage             = "sendMessage"
	MethodSendPhoto               = "sendPhoto"
	MethodSendSticker             = "sendSticker"
	MethodSendVenue               = "sendVenue"
	MethodSendVideo               = "sendVideo"
	MethodSendVideoNote           = "sendVideoNote"
	MethodSendVoice               = "sendVoice"
	MethodSetChatDescription      = "setChatDescription"
	MethodSetChatPhoto            = "setChatPhoto"
	MethodSetChatStickerSet       = "setChatStickerSet"
	MethodSetChatTitle            = "setChatTitle"
	MethodSetGameScore            = "setGameScore"
	MethodSetPassportDataErrors   = "setPassportDataErrors"
	MethodSetStickerPositionInSet = "setStickerPositionInSet"
	MethodSetWebhook              = "setWebhook"
	MethodStopMessageLiveLocation = "stopMessageLiveLocation"
	MethodUnbanChatMember         = "unbanChatMember"
	MethodUnpinChatMessage        = "unpinChatMessage"
	MethodUploadStickerFile       = "uploadStickerFile"
)

// Mode... represents available and supported parsing modes of messages
const (
	StyleHTML     = "html"
	StyleMarkdown = "markdown"
)

// Mime... represents available and supported MIME types of data
const (
	MimeHTML = "text/html"
	MimeMP4  = "video/mp4"
	MimePDF  = "application/pdf"
	MimeZIP  = "application/zip"
)

// Scheme... represents optional schemes for URLs
const (
	SchemeAttach   = "attach"
	SchemeTelegram = "tg"
)

// Status... represents available and supported statuses of ID
const (
	StatusAdministrator = "administrator"
	StatusCreator       = "creator"
	StatusKicked        = "kicked"
	StatusLeft          = "left"
	StatusMember        = "member"
	StatusRestricted    = "restricted"
)

// Type... represents available and supported types of data
const (
	TypeAddress               = "address"
	TypeArticle               = "article"
	TypeAudio                 = "audio"
	TypeBankStatement         = "bank_statement"
	TypeContact               = "contact"
	TypeDocument              = "document"
	TypeDriverLicense         = "driver_license"
	TypeEmail                 = "email"
	TypeGame                  = "game"
	TypeGIF                   = "gif"
	TypeIdentityCard          = "identity_card"
	TypeInternalPassport      = "internal_passport"
	TypeLocation              = "location"
	TypeMpeg4Gif              = "mpeg4_gif"
	TypePassport              = "passport"
	TypePassportRegistration  = "passport_registration"
	TypePersonalDetails       = "personal_details"
	TypePhoneNumber           = "phone_number"
	TypePhoto                 = "photo"
	TypeRentalAgreement       = "rental_agreement"
	TypeSticker               = "sticker"
	TypeTemporaryRegistration = "temporary_registration"
	TypeUtilityBill           = "utility_bill"
	TypeVenue                 = "venue"
	TypeVideo                 = "video"
	TypeVoice                 = "voice"
)

// Update... represents available and supported types of updates
const (
	UpdateCallbackQuery      = "callback_query"
	UpdateChannelPost        = "channel_post"
	UpdateChosenInlineResult = "chosen_inline_result"
	UpdateEditedChannelPost  = "edited_channel_post"
	UpdateEditedMessage      = "edited_message"
	UpdateInlineQuery        = "inline_query"
	UpdateMessage            = "message"
	UpdatePreCheckoutQuery   = "pre_checkout_query"
	UpdateShippingQuery      = "shipping_query"
)
