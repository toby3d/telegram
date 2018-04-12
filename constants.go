package telegram

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

const (
	ChatChannel    = "channel"
	ChatGroup      = "group"
	ChatPrivate    = "private"
	ChatSuperGroup = "supergroup"
)

const (
	EntityBold        = "bold"
	EntityBotCommand  = "bot_command"
	EntityCode        = "code"
	EntityEmail       = "email"
	EntityHashtag     = "hashtag"
	EntityItalic      = "italic"
	EntityMention     = "mention"
	EntityPre         = "pre"
	EntityTextLink    = "text_link"
	EntityTextMention = "text_mention"
	EntityURL         = "url"
)

const (
	MethodAnswerCallbackQuery     = "answerCallbackQuery"
	MethodAnswerInlineQuery       = "answerInlineQuery"
	MethodAnswerPreCheckoutQuery  = "answerPreCheckoutQuery"
	MethodAnswerShippingQuery     = "answerShippingQuery"
	MethodDeleteChatPhoto         = "deleteChatPhoto"
	MethodDeleteChatStickerSet    = "deleteChatStickerSet"
	MethodDeleteMessage           = "deleteMessage"
	MethodDeleteStickerFromSet    = "deleteStickerFromSet"
	MethodDeleteWebhook           = "deleteWebhook"
	MethodEditMessageCaption      = "editMessageCaption"
	MethodEditMessageLiveLocation = "editMessageLiveLocation"
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
	MethodSendChatAction          = "sendChatAction"
	MethodSendContact             = "sendContact"
	MethodSendDocument            = "sendDocument"
	MethodSendGame                = "sendGame"
	MethodSendInvoice             = "sendInvoice"
	MethodSendLocation            = "sendLocation"
	MethodSendMediaGroup          = "sendMediaGroup"
	MethodSendMessage             = "sendMessage"
	MethodSendPhoto               = "sendPhoto"
	MethodSendVenue               = "sendVenue"
	MethodSetChatDescription      = "setChatDescription"
	MethodSetChatPhoto            = "setChatPhoto"
	MethodSetChatStickerSet       = "setChatStickerSet"
	MethodSetChatTitle            = "setChatTitle"
	MethodSetGameScore            = "setGameScore"
	MethodSetStickerPositionInSet = "setStickerPositionInSet"
	MethodSetWebhook              = "setWebhook"
	MethodUnbanChatMember         = "unbanChatMember"
	MethodUnpinChatMessage        = "unpinChatMessage"
)

const (
	ModeHTML     = "html"
	ModeMarkdown = "markdown"
)

const (
	MimeHTML = "text/html"
	MimeMP4  = "video/mp4"
	MimePDF  = "application/pdf"
	MimeZIP  = "application/zip"
)

const PrefixAttach = "attach://"

const (
	StatusAdministrator = "administrator"
	StatusCreator       = "creator"
	StatusKicked        = "kicked"
	StatusLeft          = "left"
	StatusMember        = "member"
	StatusRestricted    = "restricted"
)

const (
	TypeArticle  = "article"
	TypeAudio    = "audio"
	TypeContact  = "contact"
	TypeDocument = "document"
	TypeGame     = "game"
	TypeGIF      = "gif"
	TypeLocation = "location"
	TypeMpeg4Gif = "mpeg4_gif"
	TypePhoto    = "photo"
	TypeSticker  = "sticker"
	TypeVenue    = "venue"
	TypeVideo    = "video"
	TypeVoice    = "voice"
)

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
