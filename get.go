//go:generate ffjson $GOFILE
package telegram

import json "github.com/pquerna/ffjson/ffjson"

type (
	// GetChatParameters represents data for GetChat method.
	GetChatParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetChatAdministratorsParameters represents data for GetChatAdministrators
	// method.
	GetChatAdministratorsParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetChatMemberParameters represents data for GetChatMember method.
	GetChatMemberParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Unique identifier of the target user
		UserID int `json:"user_id"`
	}

	// GetChatMembersCountParameters represents data for GetChatMembersCount method.
	GetChatMembersCountParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetFileParameters represents data for GetFile method.
	GetFileParameters struct {
		// File identifier to get info about
		FileID string `json:"file_id"`
	}

	// GetUpdatesParameters represents data for GetUpdates method.
	GetUpdatesParameters struct {
		// Identifier of the first update to be returned. Must be greater by one than the highest among the
		// identifiers of previously received updates. By default, updates starting with the earliest unconfirmed
		// update are returned. An update is considered confirmed as soon as getUpdates is called with an offset
		// higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset
		// update from the end of the updates queue. All previous updates will forgotten.
		Offset int `json:"offset,omitempty"`

		// Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`

		// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short
		// polling should be used for testing purposes only.
		Timeout int `json:"timeout,omitempty"`

		// List the types of updates you want your bot to receive. For example, specify ["message",
		// "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete
		// list of available update types. Specify an empty list to receive all updates regardless of type (default).
		// If not specified, the previous setting will be used.
		//
		// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so
		// unwanted updates may be received for a short period of time.
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

	// GetUserProfilePhotosParameters represents data for GetUserProfilePhotos method.
	GetUserProfilePhotosParameters struct {
		// Unique identifier of the target user
		UserID int `json:"user_id"`

		// Sequential number of the first photo to be returned. By default, all
		// photos are returned.
		Offset int `json:"offset,omitempty"`

		// Limits the number of photos to be retrieved. Values between 1—100 are
		// accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`
	}

	// GetGameHighScoresParameters represents data for GetGameHighScores method.
	GetGameHighScoresParameters struct {
		// Target user id
		UserID int `json:"user_id"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}

	// GetStickerSetParameters represents data for GetStickerSet method.
	GetStickerSetParameters struct {
		// Name of the sticker set
		Name string `json:"name"`
	}
)

// NewGameHighScores creates GetGameHighScoresParameters only with required parameters.
func NewGameHighScores(userID int) *GetGameHighScoresParameters {
	return &GetGameHighScoresParameters{
		UserID: userID,
	}
}

// GetChat get up to date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel,
// etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chatID int64) (chat *Chat, err error) {
	dst, err := json.MarshalFast(&GetChatParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChat)
	if err != nil {
		return
	}

	chat = new(Chat)
	err = json.UnmarshalFast(*resp.Result, chat)
	return
}

// GetChatAdministrators get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID int64) (members []ChatMember, err error) {
	dst, err := json.MarshalFast(&GetChatAdministratorsParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatAdministrators)
	if err != nil {
		return
	}

	err = json.UnmarshalFast(*resp.Result, &members)
	return
}

// GetChatMember get information about a member of a chat. Returns a ChatMember
// object on success.
func (bot *Bot) GetChatMember(chatID int64, userID int) (member *ChatMember, err error) {
	dst, err := json.MarshalFast(&GetChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatMember)
	if err != nil {
		return
	}

	member = new(ChatMember)
	err = json.UnmarshalFast(*resp.Result, member)
	return
}

// GetChatMembersCount get the number of members in a chat. Returns Int on
// success.
func (bot *Bot) GetChatMembersCount(chatID int64) (count int, err error) {
	dst, err := json.MarshalFast(&GetChatMembersCountParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatMembersCount)
	if err != nil {
		return
	}

	err = json.UnmarshalFast(*resp.Result, &count)
	return
}

// GetFile get basic info about a file and prepare it for downloading. For the
// moment, bots can download files of up to 20MB in size. On success, a File
// object is returned. The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is
// taken from the response. It is guaranteed that the link will be valid for at
// least 1 hour. When the link expires, a new one can be requested by calling
// getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You
// should save the file's MIME type and name (if available) when the File object
// is received.
func (bot *Bot) GetFile(fileID string) (file *File, err error) {
	dst, err := json.MarshalFast(&GetFileParameters{FileID: fileID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetFile)
	if err != nil {
		return
	}

	file = new(File)
	err = json.UnmarshalFast(*resp.Result, file)
	return
}

// GetMe testing your bot's auth token. Requires no parameters. Returns basic
// information about the bot in form of a User object.
func (bot *Bot) GetMe() (me *User, err error) {
	resp, err := bot.request(nil, MethodGetMe)
	if err != nil {
		return
	}

	me = new(User)
	err = json.UnmarshalFast(*resp.Result, me)
	return
}

// GetUpdates receive incoming updates using long polling. An Array of Update objects is returned.
func (bot *Bot) GetUpdates(params *GetUpdatesParameters) (updates []Update, err error) {
	if params == nil {
		params = &GetUpdatesParameters{Limit: 100}
	}

	src, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(src, MethodGetUpdates)
	if err != nil {
		return
	}

	updates = make([]Update, params.Limit)
	err = json.UnmarshalFast(*resp.Result, &updates)
	return
}

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (bot *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParameters) (photos *UserProfilePhotos, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetUserProfilePhotos)
	if err != nil {
		return
	}

	photos = new(UserProfilePhotos)
	err = json.UnmarshalFast(*resp.Result, photos)
	return
}

// GetWebhookInfo get current webhook status. Requires no parameters. On success,
// returns a WebhookInfo object. If the bot is using getUpdates, will return an
// object with the url field empty.
func (bot *Bot) GetWebhookInfo() (info *WebhookInfo, err error) {
	resp, err := bot.request(nil, MethodGetWebhookInfo)
	if err != nil {
		return
	}

	info = new(WebhookInfo)
	err = json.UnmarshalFast(*resp.Result, info)
	return
}

// GetGameHighScores get data for high score tables. Will return the score of the
// specified user and several of his neighbors in a game. On success, returns an
// Array of GameHighScore objects.
func (bot *Bot) GetGameHighScores(params *GetGameHighScoresParameters) (scores []GameHighScore, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetGameHighScores)
	if err != nil {
		return
	}

	err = json.UnmarshalFast(*resp.Result, &scores)
	return
}

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (set *StickerSet, err error) {
	dst, err := json.MarshalFast(&GetStickerSetParameters{Name: name})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetStickerSet)
	if err != nil {
		return
	}

	set = new(StickerSet)
	err = json.UnmarshalFast(*resp.Result, set)
	return
}
