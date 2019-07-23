//go:generate ffjson $GOFILE
package telegram

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	// Bot represents a bot user with access token getted from @BotFather and
	// fasthttp.Client for requests.
	Bot struct {
		*User
		AccessToken string
		Client      *http.Client
	}

	// UpdatesChannel is a channel for reading updates of bot.
	UpdatesChannel <-chan Update
	ShutdownFunc   func() error
)

const (
	DefaultAudioSeparator = " – "
	DefaultAudioTitle     = "[untitled]"
)

var ErrNotEqual = errors.New("credentials hash and credentials data hash is not equal")

// NewForceReply calls the response interface to the message.
func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

// NewInlineMentionURL creates a url.URL for the mention user without username.
func NewInlineMentionURL(userID int) *http.URI {
	link := http.AcquireURI()
	link.Update("tg://user?id=" + strconv.Itoa(userID))
	return link
}

func NewMarkdownBold(text string) string {
	return "*" + text + "*"
}

func NewMarkdownItalic(text string) string {
	return "_" + text + "_"
}

func NewMarkdownURL(text string, link fmt.Stringer) string {
	return "[" + text + "](" + link.String() + ")"
}

func NewMarkdownMention(text string, id int) string {
	return NewMarkdownURL(text, NewInlineMentionURL(id))
}

func NewMarkdownCode(text string) string {
	return "`" + text + "`"
}

func NewMarkdownCodeBlock(text string) string {
	return "```" + text + "```"
}

func NewHTMLBold(text string) string {
	return "<b>" + text + "</b>"
}

func NewHTMLItalic(text string) string {
	return "<i>" + text + "</i>"
}

func NewHTMLURL(text string, link fmt.Stringer) string {
	return `<a href="` + link.String() + `">` + text + `</a>`
}

func NewHTMLMention(text string, id int) string {
	return NewHTMLURL(text, NewInlineMentionURL(id))
}

func NewHTMLCode(text string) string {
	return "<code>" + text + "</code>"
}

func NewHTMLCodeBlock(text string) string {
	return "<pre>" + text + "</pre>"
}

func (a *Animation) HasThumb() bool {
	return a != nil && a.Thumb != nil
}

func (a *Animation) File() *File {
	if a == nil {
		return nil
	}

	return &File{
		FileID:   a.FileID,
		FileSize: a.FileSize,
	}
}

func (a *Audio) FullName(separator string) (name string) {
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

func (a *Audio) HasPerformer() bool {
	return a != nil && a.Performer != ""
}

func (a *Audio) HasTitle() bool {
	return a != nil && a.Title != ""
}

func (a *Audio) HasThumb() bool {
	return a != nil && a.Thumb != nil
}

func (a *Audio) File() *File {
	if a == nil {
		return nil
	}

	return &File{
		FileID:   a.FileID,
		FileSize: a.FileSize,
	}
}

// SetClient allow set custom fasthttp.Client (for proxy traffic, for example).
func (b *Bot) SetClient(newClient *http.Client) {
	if b == nil {
		b = new(Bot)
	}

	b.Client = newClient
}

// New creates a new default Bot structure based on the input access token.
func New(accessToken string) (*Bot, error) {
	var err error
	b := new(Bot)
	b.SetClient(defaultClient)
	b.AccessToken = accessToken

	b.User, err = b.GetMe()
	return b, err
}

// IsMessageFromMe checks that the input message is a message from the current
// bot.
func (b *Bot) IsMessageFromMe(m *Message) bool {
	return b != nil && b.User != nil &&
		m != nil && m.From != nil && m.From.ID == b.ID
}

// IsForwardFromMe checks that the input message is a forwarded message from the
// current bot.
func (b *Bot) IsForwardFromMe(m *Message) bool {
	return b != nil && b.User != nil &&
		m.IsForward() && m.ForwardFrom.ID == b.ID
}

// IsReplyToMe checks that the input message is a reply to the current bot.
func (b *Bot) IsReplyToMe(m *Message) bool {
	return m.Chat.IsPrivate() ||
		(m.IsReply() && b.IsMessageFromMe(m.ReplyToMessage))
}

// IsCommandToMe checks that the input message is a command for the current bot.
func (b *Bot) IsCommandToMe(m *Message) bool {
	if !m.IsCommand() {
		return false
	}

	if m.Chat.IsPrivate() {
		return true
	}

	parts := strings.Split(m.RawCommand(), "@")
	if len(parts) <= 1 {
		return false
	}

	return strings.EqualFold(parts[1], b.User.Username)
}

// IsMessageMentionsMe checks that the input message mentions the current bot.
func (b *Bot) IsMessageMentionsMe(m *Message) bool {
	if b == nil || b.User == nil || m == nil {
		return false
	}

	if b.IsCommandToMe(m) {
		return true
	}

	var entities []MessageEntity
	switch {
	case m.HasMentions():
		entities = m.Entities
	case m.HasCaptionMentions():
		entities = m.CaptionEntities
	}

	for _, entity := range entities {
		if entity.IsMention() && entity.User.ID == b.ID {
			return true
		}
	}

	return false
}

// IsForwardMentionsMe checks that the input forwarded message mentions the
// current bot.
func (b *Bot) IsForwardMentionsMe(m *Message) bool {
	return m.IsForward() && b.IsMessageMentionsMe(m)
}

// IsReplyMentionsMe checks that the input message mentions the current bot.
func (b *Bot) IsReplyMentionsMe(m *Message) bool {
	return m.IsReply() && b.IsMessageMentionsMe(m.ReplyToMessage)
}

// IsMessageToMe checks that the input message is addressed to the current bot.
func (b *Bot) IsMessageToMe(m *Message) bool {
	if m == nil || m.Chat == nil {
		return false
	}

	if m.Chat.IsPrivate() || b.IsCommandToMe(m) || b.IsReplyToMe(m) || b.IsMessageMentionsMe(m) {
		return true
	}

	return false
}

// NewFileURL creates a url.URL to file with path getted from GetFile method.
func (b *Bot) NewFileURL(filePath string) *http.URI {
	if b == nil || b.AccessToken == "" ||
		filePath == "" {
		return nil
	}

	result := http.AcquireURI()
	result.SetScheme("https")
	result.SetHost("api.telegram.org")
	result.SetPath(path.Join("file", "bot"+b.AccessToken, filePath))

	return result
}

// NewRedirectURL creates new url.URL for redirecting from one chat to another.
func (b *Bot) NewRedirectURL(param string, group bool) *http.URI {
	if b == nil || b.User == nil || b.User.Username == "" {
		return nil
	}

	link := http.AcquireURI()
	link.SetScheme("https")
	link.SetHost("t.me")
	link.SetPath(b.User.Username)

	q := link.QueryArgs()
	key := "start"
	if group {
		key += "group"
	}
	q.Set(key, param)

	link.SetQueryStringBytes(q.QueryString())

	return link
}

func (b *Bot) DecryptFile(pf *PassportFile, fc *FileCredentials) (data []byte, err error) {
	secret, err := decodeField(fc.Secret)
	if err != nil {
		return nil, err
	}

	hash, err := decodeField(fc.FileHash)
	if err != nil {
		return nil, err
	}

	key, iv := decryptSecretHash(secret, hash)
	file, err := b.GetFile(pf.FileID)
	if err != nil {
		return nil, err
	}

	if _, data, err = b.Client.Get(nil, b.NewFileURL(file.FilePath).String()); err != nil {
		return nil, err
	}

	if data, err = decryptData(key, iv, data); err != nil {
		return nil, err
	}

	if !match(hash, data) {
		err = ErrNotEqual
		return nil, err
	}

	offset := int(data[0])
	data = data[offset:]
	return data, nil
}

// IsPrivate checks that the current chat is a private chat with single user.
func (c *Chat) IsPrivate() bool {
	return c != nil && strings.EqualFold(c.Type, ChatPrivate)
}

// IsGroup checks that the current chat is a group.
func (c *Chat) IsGroup() bool {
	return c != nil && strings.EqualFold(c.Type, ChatGroup)
}

// IsSuperGroup checks that the current chat is a supergroup.
func (c *Chat) IsSuperGroup() bool {
	return c != nil && strings.EqualFold(c.Type, ChatSuperGroup)
}

// IsChannel checks that the current chat is a channel.
func (c *Chat) IsChannel() bool {
	return c != nil && strings.EqualFold(c.Type, ChatChannel)
}

// HasPinnedMessage checks that the current chat has a pinned message.
func (c *Chat) HasPinnedMessage() bool {
	return c != nil && c.PinnedMessage != nil
}

// HasStickerSet checks that the current chat has a sticker set.
func (c *Chat) HasStickerSet() bool {
	return c != nil && c.StickerSetName != ""
}

// StickerSet return StickerSet structure if StickerSetName is available.
func (c *Chat) StickerSet(bot *Bot) *StickerSet {
	if !c.HasStickerSet() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(c.StickerSetName)
	if err != nil {
		return nil
	}

	return set
}

// FullName returns the full name of chat or FirstName if LastName is not available.
func (c *Chat) FullName() string {
	if c == nil {
		return ""
	}

	if c.HasLastName() {
		return c.FirstName + " " + c.LastName
	}

	return c.FirstName
}

// HaveLastName checks what the current user has a LastName.
func (c *Chat) HasLastName() bool {
	return c != nil && c.LastName != ""
}

// HaveUsername checks what the current user has a username.
func (c *Chat) HasUsername() bool {
	return c != nil && c.Username != ""
}

func (c *Chat) HasDescription() bool {
	return c != nil && c.Description != ""
}

func (c *Chat) HasInviteLink() bool {
	return c != nil && c.InviteLink != ""
}

func (cir *ChosenInlineResult) HasLocation() bool {
	return cir != nil && cir.Location != nil
}

// FullName returns the full name of contact or FirstName if LastName is not
// available.
func (c *Contact) FullName() string {
	if c == nil {
		return ""
	}

	if c.HasLastName() {
		return c.FirstName + " " + c.LastName
	}

	return c.FirstName
}

// HaveLastName checks what the current contact has a LastName.
func (c *Contact) HasLastName() bool {
	return c != nil && c.LastName != ""
}

func (c *Contact) HasTelegram() bool {
	return c != nil && c.UserID != 0
}

func (c *Contact) HasVCard() bool {
	return c != nil && c.VCard != ""
}

func (dc *DataCredentials) decrypt(d string) (data []byte, err error) {
	secret, err := decodeField(dc.Secret)
	if err != nil {
		return
	}

	hash, err := decodeField(dc.DataHash)
	if err != nil {
		return
	}

	key, iv := decryptSecretHash(secret, hash)
	if data, err = decodeField(d); err != nil {
		return
	}

	if data, err = decryptData(key, iv, data); err != nil {
		return
	}

	if !match(hash, data) {
		err = ErrNotEqual
	}

	offset := int(data[0])
	data = data[offset:]

	return
}

func (d *Document) HasThumb() bool {
	return d != nil && d.Thumb != nil
}

func (d *Document) File() *File {
	return &File{
		FileID:   d.FileID,
		FileSize: d.FileSize,
	}
}

func (ec *EncryptedCredentials) Decrypt(pk *rsa.PrivateKey) (*Credentials, error) {
	if ec == nil || pk == nil {
		return nil, nil
	}

	data, err := decrypt(pk, ec.Secret, ec.Hash, ec.Data)
	if err != nil {
		return nil, err
	}

	var c Credentials
	err = json.UnmarshalFast(data, &c)
	return &c, err
}

func (epe *EncryptedPassportElement) DecryptPersonalDetails(sv *SecureValue) (*PersonalDetails, error) {
	if !epe.IsPersonalDetails() || !sv.HasData() {
		return nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, err
	}

	var pd PersonalDetails
	err = json.UnmarshalFast(body, &pd)
	return &pd, err
}

func (epe *EncryptedPassportElement) DecryptPassport(sv *SecureValue, b *Bot) (*IDDocumentData, []byte, []byte, [][]byte, error) {
	if !epe.IsPassport() || !sv.HasData() || !sv.HasFrontSide() {
		return nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var idd IDDocumentData
	if err = json.UnmarshalFast(body, &idd); err != nil {
		return nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, s, nil, err
			}
		}
	}

	return &idd, fs, s, t, nil
}

func (epe *EncryptedPassportElement) DecryptInternalPassport(sv *SecureValue, b *Bot) (*IDDocumentData, []byte, []byte, [][]byte, error) {
	if !epe.IsInternalPassport() || !sv.HasData() || !sv.HasFrontSide() {
		return nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var idd IDDocumentData
	if err = json.UnmarshalFast(body, &idd); err != nil {
		return nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, s, nil, err
			}
		}
	}

	return &idd, fs, s, t, nil
}

func (epe *EncryptedPassportElement) DecryptDriverLicense(sv *SecureValue, b *Bot) (*IDDocumentData, []byte, []byte, []byte, [][]byte, error) {
	if !epe.IsDriverLicense() || !sv.HasData() || !sv.HasFrontSide() || !sv.HasReverseSide() {
		return nil, nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	var idd IDDocumentData
	if err = json.UnmarshalFast(body, &idd); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, nil, err
	}

	rs, err := b.DecryptFile(epe.ReverseSide, sv.ReverseSide)
	if err != nil {
		return &idd, nil, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, rs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, rs, s, nil, err
			}
		}
	}

	return &idd, fs, rs, s, t, nil
}

func (epe *EncryptedPassportElement) IsAddress() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeAddress)
}

func (epe *EncryptedPassportElement) IsBankStatement() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeBankStatement)
}

func (epe *EncryptedPassportElement) IsDriverLicense() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeDriverLicense)
}

func (epe *EncryptedPassportElement) IsEmail() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeEmail)
}

func (epe *EncryptedPassportElement) IsIdentityCard() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeIdentityCard)
}

func (epe *EncryptedPassportElement) IsInternalPassport() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeInternalPassport)
}

func (epe *EncryptedPassportElement) IsPassport() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePassport)
}

func (epe *EncryptedPassportElement) IsPassportRegistration() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePassportRegistration)
}

func (epe *EncryptedPassportElement) IsPersonalDetails() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePersonalDetails)
}

func (epe *EncryptedPassportElement) IsPhoneNumber() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePhoneNumber)
}

func (epe *EncryptedPassportElement) IsRentalAgreement() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeRentalAgreement)
}

func (epe *EncryptedPassportElement) IsTemporaryRegistration() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeTemporaryRegistration)
}

func (epe *EncryptedPassportElement) IsUtilityBill() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeUtilityBill)
}

// ParseURL selects URL from entered text of message and parse it as fasthttp.URI.
func (e *MessageEntity) ParseURL(messageText string) *http.URI {
	if e == nil || !e.IsURL() || messageText == "" {
		return nil
	}

	from := e.Offset
	to := from + e.Length
	text := []rune(messageText)
	if len(text) < to {
		return nil
	}

	link := http.AcquireURI()
	link.Update(string(text[from:to]))

	return link
}

// IsBold checks that the current entity is a bold tag.
func (e *MessageEntity) IsBold() bool {
	return e != nil && strings.EqualFold(e.Type, EntityBold)
}

// IsBotCommand checks that the current entity is a bot command.
func (e *MessageEntity) IsBotCommand() bool {
	return e != nil && strings.EqualFold(e.Type, EntityBotCommand)
}

// IsCode checks that the current entity is a code tag.
func (e *MessageEntity) IsCode() bool {
	return e != nil && strings.EqualFold(e.Type, EntityCode)
}

// IsEmail checks that the current entity is a email.
func (e *MessageEntity) IsEmail() bool {
	return e != nil && strings.EqualFold(e.Type, EntityEmail)
}

// IsHashtag checks that the current entity is a hashtag.
func (e *MessageEntity) IsHashtag() bool {
	return e != nil && strings.EqualFold(e.Type, EntityHashtag)
}

// IsItalic checks that the current entity is a italic tag.
func (e *MessageEntity) IsItalic() bool {
	return e != nil && strings.EqualFold(e.Type, EntityItalic)
}

// IsMention checks that the current entity is a username mention.
func (e *MessageEntity) IsMention() bool {
	return e != nil && strings.EqualFold(e.Type, EntityMention)
}

// IsPre checks that the current entity is a pre tag.
func (e *MessageEntity) IsPre() bool {
	return e != nil && strings.EqualFold(e.Type, EntityPre)
}

// IsTextLink checks that the current entity is a text link.
func (e *MessageEntity) IsTextLink() bool {
	return e != nil && strings.EqualFold(e.Type, EntityTextLink)
}

// IsTextMention checks that the current entity is a mention without username.
func (e *MessageEntity) IsTextMention() bool {
	return e != nil && strings.EqualFold(e.Type, EntityTextMention)
}

// IsURL checks that the current entity is a URL
func (e *MessageEntity) IsURL() bool {
	return e != nil && strings.EqualFold(e.Type, EntityURL)
}

// TextLink parse current text link entity as fasthttp.URI.
func (e *MessageEntity) TextLink() *http.URI {
	if e == nil {
		return nil
	}

	link := http.AcquireURI()
	link.Update(e.URL)

	return link
}

func (idd *IDDocumentData) ExpiryTime() *time.Time {
	if idd == nil || idd.ExpiryDate == "" {
		return nil
	}

	et, err := time.Parse("02.01.2006", idd.ExpiryDate)
	if err != nil {
		return nil
	}

	return &et
}

// NewInlineKeyboardMarkup creates a new inline keyboard markup for message.
func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton
	keyboard = append(keyboard, rows...)
	return &InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewInlineKeyboardRow creates a new inline keyboard row for buttons.
func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}

// NewInlineKeyboardButton creates a new inline keyboard callback button.
func NewInlineKeyboardButton(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

// NewInlineKeyboardButtonURL creates a new inline keyboard button with URL.
func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

// NewInlineKeyboardButtonSwitch creates a new inline keyboard button to make
// specific inline query in other chat.
func NewInlineKeyboardButtonSwitch(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

// NewInlineKeyboardButtonSwitchSelf creates a new inline keyboard button to make
// specific inline query in same chat.
func NewInlineKeyboardButtonSwitchSelf(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:                         text,
		SwitchInlineQueryCurrentChat: query,
	}
}

// NewInlineKeyboardButtonGame creates a new inline keyboard button with game
// callback.
func NewInlineKeyboardButtonGame(text string) InlineKeyboardButton {
	var game CallbackGame
	return InlineKeyboardButton{
		Text:         text,
		CallbackGame: &game,
	}
}

// NewInlineKeyboardButtonPay creates a new inline keyboard button with pay
// callback.
func NewInlineKeyboardButtonPay(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}

// HasLocation checks what current InlineQuery has Location info.
func (iq *InlineQuery) HasLocation() bool {
	return iq != nil && iq.Location != nil
}

// HasOffset checks what current InlineQuery has Offset.
func (iq *InlineQuery) HasOffset() bool {
	return iq != nil && iq.Offset != ""
}

// HasQuery checks what current InlineQuery has Query string.
func (iq *InlineQuery) HasQuery() bool {
	return iq != nil && iq.Query != ""
}

// NewInlineQueryResultCachedAudio creates a new inline query result with cached
// audio.
func NewInlineQueryResultCachedAudio(resultID, fileID string) *InlineQueryResultCachedAudio {
	return &InlineQueryResultCachedAudio{
		Type:        TypeAudio,
		ID:          resultID,
		AudioFileID: fileID,
	}
}

// NewInlineQueryResultCachedDocument creates a new inline query result with
// cached document.
func NewInlineQueryResultCachedDocument(resultID, fileID, title string) *InlineQueryResultCachedDocument {
	return &InlineQueryResultCachedDocument{
		Type:           TypeDocument,
		ID:             resultID,
		Title:          title,
		DocumentFileID: fileID,
	}
}

// NewInlineQueryResultCachedGif creates a new inline query result with cached
// GIF.
func NewInlineQueryResultCachedGif(resultID, fileID string) *InlineQueryResultCachedGif {
	return &InlineQueryResultCachedGif{
		Type:      TypeGIF,
		ID:        resultID,
		GifFileID: fileID,
	}
}

// NewInlineQueryResultCachedMpeg4Gif creates a new inline query result with
// cached MPEG GIF.
func NewInlineQueryResultCachedMpeg4Gif(resultID, fileID string) *InlineQueryResultCachedMpeg4Gif {
	return &InlineQueryResultCachedMpeg4Gif{
		Type:        TypeMpeg4Gif,
		ID:          resultID,
		Mpeg4FileID: fileID,
	}
}

// NewInlineQueryResultCachedPhoto creates a new inline query result with cached
// photo.
func NewInlineQueryResultCachedPhoto(resultID, fileID string) *InlineQueryResultCachedPhoto {
	return &InlineQueryResultCachedPhoto{
		Type:        TypePhoto,
		ID:          resultID,
		PhotoFileID: fileID,
	}
}

// NewInlineQueryResultCachedSticker creates a new inline query result with
// cached sticker.
func NewInlineQueryResultCachedSticker(resultID, fileID string) *InlineQueryResultCachedSticker {
	return &InlineQueryResultCachedSticker{
		Type:          TypeSticker,
		ID:            resultID,
		StickerFileID: fileID,
	}
}

// NewInlineQueryResultCachedVideo creates a new inline query result with cached
// video.
func NewInlineQueryResultCachedVideo(resultID, fileID, title string) *InlineQueryResultCachedVideo {
	return &InlineQueryResultCachedVideo{
		Type:        TypeVideo,
		ID:          resultID,
		VideoFileID: fileID,
		Title:       title,
	}
}

// NewInlineQueryResultCachedVoice creates a new inline query result with cached
// voice.
func NewInlineQueryResultCachedVoice(resultID, fileID, title string) *InlineQueryResultCachedVoice {
	return &InlineQueryResultCachedVoice{
		Type:        TypeVoice,
		ID:          resultID,
		VoiceFileID: fileID,
		Title:       title,
	}
}

// NewInlineQueryResultArticle creates a new inline query result with article.
func NewInlineQueryResultArticle(resultID, title string, content interface{}) *InlineQueryResultArticle {
	return &InlineQueryResultArticle{
		Type:                TypeArticle,
		ID:                  resultID,
		Title:               title,
		InputMessageContent: content,
	}
}

// NewInlineQueryResultAudio creates a new inline query result with audio.
func NewInlineQueryResultAudio(resultID, audioURL, title string) *InlineQueryResultAudio {
	return &InlineQueryResultAudio{
		Type:     TypeAudio,
		ID:       resultID,
		AudioURL: audioURL,
		Title:    title,
	}
}

// NewInlineQueryResultContact creates a new inline query result with contact.
func NewInlineQueryResultContact(resultID, phoneNumber, firstName string) *InlineQueryResultContact {
	return &InlineQueryResultContact{
		Type:        TypeContact,
		ID:          resultID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewInlineQueryResultGame creates a new inline query result with game.
func NewInlineQueryResultGame(resultID, gameShortName string) *InlineQueryResultGame {
	return &InlineQueryResultGame{
		Type:          TypeGame,
		ID:            resultID,
		GameShortName: gameShortName,
	}
}

// NewInlineQueryResultDocument creates a new inline query result with document.
func NewInlineQueryResultDocument(resultID, title, documentURL, mimeType string) *InlineQueryResultDocument {
	return &InlineQueryResultDocument{
		Type:        TypeDocument,
		ID:          resultID,
		Title:       title,
		DocumentURL: documentURL,
		MimeType:    mimeType,
	}
}

// NewInlineQueryResultGif creates a new inline query result with GIF.
func NewInlineQueryResultGif(resultID, gifURL, thumbURL string) *InlineQueryResultGif {
	return &InlineQueryResultGif{
		Type:     TypeGIF,
		ID:       resultID,
		GifURL:   gifURL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultLocation creates a new inline query result with location.
func NewInlineQueryResultLocation(resultID, title string, latitude, longitude float32) *InlineQueryResultLocation {
	return &InlineQueryResultLocation{
		Type:      TypeLocation,
		ID:        resultID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
	}
}

// NewInlineQueryResultMpeg4Gif creates a new inline query result with MPEG GIF.
func NewInlineQueryResultMpeg4Gif(resultID, mpeg4URL, thumbURL string) *InlineQueryResultMpeg4Gif {
	return &InlineQueryResultMpeg4Gif{
		Type:     TypeMpeg4Gif,
		ID:       resultID,
		Mpeg4URL: mpeg4URL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultPhoto creates a new inline query result with photo.
func NewInlineQueryResultPhoto(resultID, photoURL, thumbURL string) *InlineQueryResultPhoto {
	return &InlineQueryResultPhoto{
		Type:     TypePhoto,
		ID:       resultID,
		PhotoURL: photoURL,
		ThumbURL: thumbURL,
	}
}

// NewInlineQueryResultVenue creates a new inline query result with venue.
func NewInlineQueryResultVenue(resultID, title, address string, latitude, longitude float32) *InlineQueryResultVenue {
	return &InlineQueryResultVenue{
		Type:      TypeVenue,
		ID:        resultID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// NewInlineQueryResultVideo creates a new inline query result with video.
func NewInlineQueryResultVideo(resultID, videoURL, mimeType, thumbURL, title string) *InlineQueryResultVideo {
	return &InlineQueryResultVideo{
		Type:     TypeVideo,
		ID:       resultID,
		VideoURL: videoURL,
		MimeType: mimeType,
		ThumbURL: thumbURL,
		Title:    title,
	}
}

// NewInlineQueryResultVoice creates a new inline query result with voice.
func NewInlineQueryResultVoice(resultID, voiceURL, title string) *InlineQueryResultVoice {
	return &InlineQueryResultVoice{
		Type:     TypeVoice,
		ID:       resultID,
		VoiceURL: voiceURL,
		Title:    title,
	}
}

func (iqra *InlineQueryResultArticle) ResultID() string {
	return iqra.ID
}

func (iqra *InlineQueryResultArticle) ResultType() string {
	return iqra.Type
}

func (iqra *InlineQueryResultArticle) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqra.ReplyMarkup
}

func (iqrp *InlineQueryResultPhoto) ResultID() string {
	return iqrp.ID
}

func (iqrp *InlineQueryResultPhoto) ResultType() string {
	return iqrp.Type
}

func (iqrp *InlineQueryResultPhoto) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrp.ReplyMarkup
}

func (iqrg *InlineQueryResultGif) ResultID() string {
	return iqrg.ID
}

func (iqrg *InlineQueryResultGif) ResultType() string {
	return iqrg.Type
}

func (iqrg *InlineQueryResultGif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrg.ReplyMarkup
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultID() string {
	return iqrm4g.ID
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultType() string {
	return iqrm4g.Type
}

func (iqrm4g *InlineQueryResultMpeg4Gif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrm4g.ReplyMarkup
}

func (iqrv *InlineQueryResultVideo) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVideo) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVideo) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqra *InlineQueryResultAudio) ResultID() string {
	return iqra.ID
}

func (iqra *InlineQueryResultAudio) ResultType() string {
	return iqra.Type
}

func (iqra *InlineQueryResultAudio) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqra.ReplyMarkup
}

func (iqrv *InlineQueryResultVoice) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVoice) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVoice) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqrd *InlineQueryResultDocument) ResultID() string {
	return iqrd.ID
}

func (iqrd *InlineQueryResultDocument) ResultType() string {
	return iqrd.Type
}

func (iqrd *InlineQueryResultDocument) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrd.ReplyMarkup
}

func (iqrl *InlineQueryResultLocation) ResultID() string {
	return iqrl.ID
}

func (iqrl *InlineQueryResultLocation) ResultType() string {
	return iqrl.Type
}

func (iqrl *InlineQueryResultLocation) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrl.ReplyMarkup
}

func (iqrv *InlineQueryResultVenue) ResultID() string {
	return iqrv.ID
}

func (iqrv *InlineQueryResultVenue) ResultType() string {
	return iqrv.Type
}

func (iqrv *InlineQueryResultVenue) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrv.ReplyMarkup
}

func (iqrc *InlineQueryResultContact) ResultID() string {
	return iqrc.ID
}

func (iqrc *InlineQueryResultContact) ResultType() string {
	return iqrc.Type
}

func (iqrc *InlineQueryResultContact) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrc.ReplyMarkup
}

func (iqrg *InlineQueryResultGame) ResultID() string {
	return iqrg.ID
}

func (iqrg *InlineQueryResultGame) ResultType() string {
	return iqrg.Type
}

func (iqrg *InlineQueryResultGame) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrg.ReplyMarkup
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultID() string {
	return iqrcp.ID
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultType() string {
	return iqrcp.Type
}

func (iqrcp *InlineQueryResultCachedPhoto) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcp.ReplyMarkup
}

func (iqrcg *InlineQueryResultCachedGif) ResultID() string {
	return iqrcg.ID
}

func (iqrcg *InlineQueryResultCachedGif) ResultType() string {
	return iqrcg.Type
}

func (iqrcg *InlineQueryResultCachedGif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcg.ReplyMarkup
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultID() string {
	return iqrcm4g.ID
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultType() string {
	return iqrcm4g.Type
}

func (iqrcm4g *InlineQueryResultCachedMpeg4Gif) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcm4g.ReplyMarkup
}

func (iqrcs *InlineQueryResultCachedSticker) ResultID() string {
	return iqrcs.ID
}

func (iqrcs *InlineQueryResultCachedSticker) ResultType() string {
	return iqrcs.Type
}

func (iqrcs *InlineQueryResultCachedSticker) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcs.ReplyMarkup
}

func (iqrcd *InlineQueryResultCachedDocument) ResultID() string {
	return iqrcd.ID
}

func (iqrcd *InlineQueryResultCachedDocument) ResultType() string {
	return iqrcd.Type
}

func (iqrcd *InlineQueryResultCachedDocument) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcd.ReplyMarkup
}

func (iqrcv *InlineQueryResultCachedVideo) ResultID() string {
	return iqrcv.ID
}

func (iqrcv *InlineQueryResultCachedVideo) ResultType() string {
	return iqrcv.Type
}

func (iqrcv *InlineQueryResultCachedVideo) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcv.ReplyMarkup
}

func (iqrcv *InlineQueryResultCachedVoice) ResultID() string {
	return iqrcv.ID
}

func (iqrcv *InlineQueryResultCachedVoice) ResultType() string {
	return iqrcv.Type
}

func (iqrcv *InlineQueryResultCachedVoice) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrcv.ReplyMarkup
}

func (iqrca *InlineQueryResultCachedAudio) ResultID() string {
	return iqrca.ID
}

func (iqrca *InlineQueryResultCachedAudio) ResultType() string {
	return iqrca.Type
}

func (iqrca *InlineQueryResultCachedAudio) ResultReplyMarkup() *InlineKeyboardMarkup {
	return iqrca.ReplyMarkup
}

// NewInputTextMessageContent creates a new text of message.
func NewInputTextMessageContent(messageText string) *InputTextMessageContent {
	return &InputTextMessageContent{
		MessageText: messageText,
	}
}

// NewInputLocationMessageContent creates a new location.
func NewInputLocationMessageContent(latitude, longitude float32) *InputLocationMessageContent {
	return &InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewInputVenueMessageContent creates a new venue.
func NewInputVenueMessageContent(latitude, longitude float32, title, address string) *InputVenueMessageContent {
	return &InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// NewInputContactMessageContent creates a new contact.
func NewInputContactMessageContent(phoneNumber, firstName string) *InputContactMessageContent {
	return &InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewInputMediaPhoto creates a new photo in media album.
func NewInputMediaPhoto(media string) *InputMediaPhoto {
	return &InputMediaPhoto{
		Type:  TypePhoto,
		Media: media,
	}
}

// NewInputMediaVideo creates a new video in media album.
func NewInputMediaVideo(media string) *InputMediaVideo {
	return &InputMediaVideo{
		Type:  TypeVideo,
		Media: media,
	}
}

func (ima *InputMediaAnimation) File() string {
	if ima == nil {
		return ""
	}

	return ima.Media
}

func (ima *InputMediaAnimation) InputMediaCaption() string {
	if ima == nil {
		return ""
	}

	return ima.Caption
}

func (ima *InputMediaAnimation) InputMediaParseMode() string {
	if ima == nil {
		return ""
	}

	return ima.ParseMode
}

func (ima *InputMediaAnimation) InputMediaType() string {
	if ima == nil {
		return ""
	}

	return ima.Type
}

func (imd *InputMediaDocument) File() string {
	if imd == nil {
		return ""
	}

	return imd.Media
}

func (imd *InputMediaDocument) InputMediaCaption() string {
	if imd == nil {
		return ""
	}

	return imd.Caption
}

func (imd *InputMediaDocument) InputMediaParseMode() string {
	if imd == nil {
		return ""
	}

	return imd.ParseMode
}

func (imd *InputMediaDocument) InputMediaType() string {
	if imd == nil {
		return ""
	}

	return imd.Type
}

func (ima *InputMediaAudio) File() string {
	if ima == nil {
		return ""
	}

	return ima.Media
}

func (ima *InputMediaAudio) InputMediaCaption() string {
	if ima == nil {
		return ""
	}

	return ima.Caption
}

func (ima *InputMediaAudio) InputMediaParseMode() string {
	if ima == nil {
		return ""
	}

	return ima.ParseMode
}

func (ima *InputMediaAudio) InputMediaType() string {
	if ima == nil {
		return ""
	}

	return ima.Type
}

func (imp *InputMediaPhoto) File() string {
	if imp == nil {
		return ""
	}

	return imp.Media
}

func (imp *InputMediaPhoto) InputMediaCaption() string {
	if imp == nil {
		return ""
	}

	return imp.Caption
}

func (imp *InputMediaPhoto) InputMediaParseMode() string {
	if imp == nil {
		return ""
	}

	return imp.ParseMode
}

func (imp *InputMediaPhoto) InputMediaType() string {
	if imp == nil {
		return ""
	}

	return imp.Type
}

func (imv *InputMediaVideo) File() string {
	if imv == nil {
		return ""
	}

	return imv.Media
}

func (imv *InputMediaVideo) InputMediaCaption() string {
	if imv == nil {
		return ""
	}

	return imv.Caption
}

func (imv *InputMediaVideo) InputMediaParseMode() string {
	if imv == nil {
		return ""
	}

	return imv.ParseMode
}

func (imv *InputMediaVideo) InputMediaType() string {
	if imv == nil {
		return ""
	}

	return imv.Type
}

func (itmc *InputTextMessageContent) IsInputMessageContent() bool { return true }

func (ilmc *InputLocationMessageContent) IsInputMessageContent() bool { return true }

func (ivmc *InputVenueMessageContent) IsInputMessageContent() bool { return true }

func (icmc *InputContactMessageContent) IsInputMessageContent() bool { return true }

// NewReplyKeyboardRemove just hides keyboard.
func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewReplyKeyboardMarkup creates new keyboard markup of simple buttons.
func NewReplyKeyboardMarkup(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton
	keyboard = append(keyboard, rows...)
	return &ReplyKeyboardMarkup{Keyboard: keyboard}
}

// NewReplyKeyboardRow creates new keyboard row for buttons.
func NewReplyKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}

// NewReplyKeyboardButton creates new button with custom text for sending it.
func NewReplyKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

// NewReplyKeyboardButtonContact creates new button with custom text for sending
// user contact.
func NewReplyKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// NewReplyKeyboardButtonLocation creates new button with custom text for sending
// user location.
func NewReplyKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

// IsCreator checks that current member is creator.
func (m *ChatMember) IsCreator() bool {
	return m != nil && strings.EqualFold(m.Status, StatusCreator)
}

// IsAdministrator checks that current member is administrator.
func (m *ChatMember) IsAdministrator() bool {
	return m != nil && strings.EqualFold(m.Status, StatusAdministrator)
}

// IsRestricted checks that current member has been restricted.
func (m *ChatMember) IsRestricted() bool {
	return m != nil && strings.EqualFold(m.Status, StatusRestricted)
}

// IsLeft checks that current member has left the chat.
func (m *ChatMember) IsLeft() bool {
	return m != nil && strings.EqualFold(m.Status, StatusLeft)
}

// IsKicked checks that current member has been kicked.
func (m *ChatMember) IsKicked() bool {
	return m != nil && strings.EqualFold(m.Status, StatusKicked)
}

// UntilTime parse UntilDate of restrictions and returns time.Time.
func (m *ChatMember) UntilTime() *time.Time {
	if m == nil {
		return nil
	}

	ut := time.Unix(m.UntilDate, 0)
	return &ut
}

// IsCommand checks that the current message is a bot command.
func (m *Message) IsCommand() bool {
	if !m.IsText() || !m.HasEntities() {
		return false
	}

	entity := m.Entities[0]
	return entity.IsBotCommand() && entity.Offset == 0
}

// IsCommandEqual checks that the current message is a specific bot command.
func (m *Message) IsCommandEqual(command string) bool {
	return m.IsCommand() && strings.EqualFold(m.Command(), command)
}

// Command returns identifier of the bot command without bot username, if it was
// available
func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	return strings.Split(m.RawCommand(), "@")[0]
}

// RawCommand returns identifier of the bot command with bot username, if it was
// available
func (m *Message) RawCommand() string {
	if !m.IsCommand() {
		return ""
	}

	return string([]rune(m.Text)[1:m.Entities[0].Length])
}

// HasCommandArgument checks that the current command message contains argument.
func (m *Message) HasCommandArgument() bool {
	if !m.IsCommand() {
		return false
	}

	entity := m.Entities[0]
	if !entity.IsBotCommand() {
		return false
	}

	return len([]rune(m.Text)) != entity.Length
}

// CommandArgument returns raw command argument.
func (m *Message) CommandArgument() string {
	if !m.HasCommandArgument() {
		return ""
	}

	return string([]rune(m.Text)[m.Entities[0].Length+1:])
}

// IsReply checks that the current message is a reply on other message.
func (m *Message) IsReply() bool {
	return m != nil && m.ReplyToMessage != nil
}

// IsForward checks that the current message is a forward of other message.
func (m *Message) IsForward() bool {
	return m != nil && m.ForwardFrom != nil
}

// Time parse current message Date and returns time.Time.
func (m *Message) Time() *time.Time {
	if m == nil {
		return nil
	}

	t := time.Unix(m.Date, 0)
	return &t
}

// ForwardTime parse current message ForwardDate and returns time.Time.
func (m *Message) ForwardTime() *time.Time {
	if m == nil {
		return nil
	}

	ft := time.Unix(m.ForwardDate, 0)
	return &ft
}

// EditTime parse current message EditDate and returns time.Time.
func (m *Message) EditTime() *time.Time {
	if m == nil || !m.HasBeenEdited() {
		return nil
	}

	et := time.Unix(m.EditDate, 0)
	return &et
}

// HasBeenEdited checks that the current message has been edited.
func (m *Message) HasBeenEdited() bool {
	return m != nil && m.EditDate > 0
}

// IsText checks that the current message is just a text message.
func (m *Message) IsText() bool {
	return m != nil && m.Text != ""
}

// IsAudio checks that the current message is a audio.
func (m *Message) IsAudio() bool {
	return m != nil && m.Audio != nil
}

// IsDocument checks that the current message is a document.
func (m *Message) IsDocument() bool {
	return m != nil && m.Document != nil
}

// IsGame checks that the current message is a game.
func (m *Message) IsGame() bool {
	return m != nil && m.Game != nil
}

// IsPhoto checks that the current message is a photo.
func (m *Message) IsPhoto() bool {
	return m != nil && len(m.Photo) > 0
}

// IsSticker checks that the current message is a sticker.
func (m *Message) IsSticker() bool {
	return m != nil && m.Sticker != nil
}

// IsVideo checks that the current message is a video.
func (m *Message) IsVideo() bool {
	return m != nil && m.Video != nil
}

// IsVoice checks that the current message is a voice.
func (m *Message) IsVoice() bool {
	return m != nil && m.Voice != nil
}

// IsVideoNote checks that the current message is a video note.
func (m *Message) IsVideoNote() bool {
	return m != nil && m.VideoNote != nil
}

// IsContact checks that the current message is a contact.
func (m *Message) IsContact() bool {
	return m != nil && m.Contact != nil
}

// IsLocation checks that the current message is a location.
func (m *Message) IsLocation() bool {
	return m != nil && m.Location != nil
}

// IsVenue checks that the current message is a venue.
func (m *Message) IsVenue() bool {
	return m != nil && m.Venue != nil
}

// IsAnimation checks that the current message is a animation.
func (m *Message) IsAnimation() bool {
	return m != nil && m.Animation != nil
}

// IsNewChatMembersEvent checks that the current message is a event of entry of
// new members.
func (m *Message) IsNewChatMembersEvent() bool {
	return m != nil && len(m.NewChatMembers) > 0
}

// IsLeftChatMemberEvent checks that the current message is a event of members
// exit.
func (m *Message) IsLeftChatMemberEvent() bool {
	return m != nil && m.LeftChatMember != nil
}

// IsNewChatTitleEvent checks that the current message is a event of setting a
// new chat title.
func (m *Message) IsNewChatTitleEvent() bool {
	return m != nil && !strings.EqualFold(m.NewChatTitle, "")
}

// IsNewChatPhotoEvent checks that the current message is a event of setting a
// new chat avatar.
func (m *Message) IsNewChatPhotoEvent() bool {
	return m != nil && len(m.NewChatPhoto) > 0
}

// IsDeleteChatPhotoEvent checks that the current message is a event of deleting
// a chat avatar.
func (m *Message) IsDeleteChatPhotoEvent() bool {
	return m != nil && m.DeleteChatPhoto
}

// IsGroupChatCreatedEvent checks that the current message is a event of creating
// a new group.
func (m *Message) IsGroupChatCreatedEvent() bool {
	return m != nil && m.GroupChatCreated
}

// IsSupergroupChatCreatedEvent checks that the current message is a event of
// creating a new supergroup.
func (m *Message) IsSupergroupChatCreatedEvent() bool {
	return m != nil && m.SupergroupChatCreated
}

// IsChannelChatCreatedEvent checks that the current message is a event of
// creating a new channel.
func (m *Message) IsChannelChatCreatedEvent() bool {
	return m != nil && m.ChannelChatCreated
}

// IsPinnedMessage checks that the current message is a event of pinning another
// message.
func (m *Message) IsPinnedMessage() bool {
	return m != nil && m.PinnedMessage != nil
}

// IsInvoice checks that the current message is a invoice.
func (m *Message) IsInvoice() bool {
	return m != nil && m.Invoice != nil
}

// IsSuccessfulPayment checks that the current message is a event of successful
// payment.
func (m *Message) IsSuccessfulPayment() bool {
	return m != nil && m.SuccessfulPayment != nil
}

// HasEntities checks that the current message contains entities.
func (m *Message) HasEntities() bool {
	return m != nil && len(m.Entities) > 0
}

// HasCaptionEntities checks that the current media contains entities in caption.
func (m *Message) HasCaptionEntities() bool {
	return m != nil && len(m.CaptionEntities) > 0
}

// HasMentions checks that the current message contains mentions.
func (m *Message) HasMentions() bool {
	if !m.HasEntities() {
		return false
	}

	for _, entity := range m.Entities {
		if entity.IsMention() || entity.IsTextMention() {
			return true
		}
	}

	return false
}

// HasCaptionMentions checks that the current media contains mentions in caption.
func (m *Message) HasCaptionMentions() bool {
	if !m.HasCaptionEntities() {
		return false
	}

	for _, entity := range m.CaptionEntities {
		if entity.IsMention() || entity.IsTextMention() {
			return true
		}
	}

	return false
}

// HasCaption checks that the current media has caption.
func (m *Message) HasCaption() bool {
	return m != nil && m.Caption != ""
}

// HasAuthorSignature checks that the current channel post has author signature.
func (m *Message) HasAuthorSignature() bool {
	return m != nil && m.AuthorSignature != ""
}

// IsEvent checks what current message is a any chat event.
func (m *Message) IsEvent() bool {
	return m.IsChannelChatCreatedEvent() ||
		m.IsDeleteChatPhotoEvent() ||
		m.IsGroupChatCreatedEvent() ||
		m.IsLeftChatMemberEvent() ||
		m.IsNewChatMembersEvent() ||
		m.IsNewChatTitleEvent() ||
		m.IsSupergroupChatCreatedEvent() ||
		m.IsNewChatPhotoEvent()
}

func sortPhotos(ps []PhotoSize, reverse bool) []PhotoSize {
	buf := make([]PhotoSize, len(ps))
	copy(buf, ps)

	sort.Slice(buf, func(i, j int) bool {
		if reverse {
			return buf[i].Width > buf[j].Width &&
				buf[i].Height > buf[j].Height
		}

		return buf[i].Width < buf[j].Width &&
			buf[i].Height < buf[j].Height
	})

	return buf
}

func (m *Message) BigPhoto() *PhotoSize {
	if m == nil || !m.IsPhoto() {
		return nil
	}

	if len(m.Photo) == 1 {
		return &m.Photo[0]
	}

	sp := sortPhotos(m.Photo, true)
	return &sp[0]
}

func (m *Message) SmallPhoto() *PhotoSize {
	if m == nil || !m.IsPhoto() {
		return nil
	}

	if len(m.Photo) == 1 {
		return &m.Photo[0]
	}

	sp := sortPhotos(m.Photo, false)
	return &sp[0]
}

func (m *Message) BigChatPhoto() *PhotoSize {
	if m == nil || !m.IsNewChatPhotoEvent() {
		return nil
	}

	if len(m.NewChatPhoto) == 1 {
		return &m.NewChatPhoto[0]
	}

	sp := sortPhotos(m.NewChatPhoto, true)
	return &sp[0]
}

func (m *Message) SmallChatPhoto() *PhotoSize {
	if m == nil || !m.IsNewChatPhotoEvent() {
		return nil
	}

	if len(m.NewChatPhoto) == 1 {
		return &m.NewChatPhoto[0]
	}

	sp := sortPhotos(m.NewChatPhoto, false)
	return &sp[0]
}

func (m *Message) HasPoll() bool {
	return m != nil && m.Poll != nil
}

func decrypt(pk *rsa.PrivateKey, s, h, d string) (obj []byte, err error) {
	// Note that all base64-encoded fields should be decoded before use.
	secret, err := decodeField(s)
	if err != nil {
		return nil, err
	}

	hash, err := decodeField(h)
	if err != nil {
		return nil, err
	}

	data, err := decodeField(d)
	if err != nil {
		return nil, err
	}

	if pk != nil {
		// Decrypt the credentials secret (secret field in EncryptedCredentials)
		// using your private key
		secret, err = decryptSecret(pk, secret)
		if err != nil {
			return nil, err
		}
	}

	// Use this secret and the credentials hash (hash field in
	// EncryptedCredentials) to calculate credentials_key and credentials_iv
	key, iv := decryptSecretHash(secret, hash)
	if err != nil {
		return nil, err
	}

	// Decrypt the credentials data (data field in EncryptedCredentials) by
	// AES256-CBC using these credentials_key and credentials_iv.
	data, err = decryptData(key, iv, data)
	if err != nil {
		return nil, err
	}

	// IMPORTANT: At this step, make sure that the credentials hash is equal
	// to SHA256(credentials_data)
	if !match(hash, data) {
		return nil, ErrNotEqual
	}

	// Credentials data is padded with 32 to 255 random padding bytes to make
	// its length divisible by 16 bytes. The first byte contains the length
	// of this padding (including this byte). Remove the padding to get the
	// data.
	offset := int(data[0])
	return data[offset:], nil
}

func decodeField(rawField string) (field []byte, err error) {
	return base64.StdEncoding.DecodeString(rawField)
}

func decryptSecret(pk *rsa.PrivateKey, s []byte) (secret []byte, err error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, pk, s, nil)
}

func decryptSecretHash(s, h []byte) (key, iv []byte) {
	hash := sha512.New()
	var err error
	if _, err = hash.Write(s); err != nil {
		return
	}
	if _, err = hash.Write(h); err != nil {
		return
	}
	sh := hash.Sum(nil)

	return sh[0:32], sh[32 : 32+16]
}

func match(h, d []byte) bool {
	dh := sha256.New()
	if _, err := dh.Write(d); err != nil {
		return false
	}

	return bytes.EqualFold(h, dh.Sum(nil))
}

func decryptData(key, iv, data []byte) (buf []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	buf = make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(buf, data)

	return
}

func (peedf *PassportElementErrorDataField) PassportElementErrorMessage() string {
	if peedf == nil {
		return ""
	}

	return peedf.Message
}

func (peedf *PassportElementErrorDataField) PassportElementErrorSource() string {
	if peedf == nil {
		return ""
	}

	return peedf.Source
}

func (peedf *PassportElementErrorDataField) PassportElementErrorType() string {
	if peedf == nil {
		return ""
	}

	return peedf.Type
}

func (peeff *PassportElementErrorFrontSide) PassportElementErrorMessage() string {
	if peeff == nil {
		return ""
	}

	return peeff.Message
}

func (peeff *PassportElementErrorFrontSide) PassportElementErrorSource() string {
	if peeff == nil {
		return ""
	}

	return peeff.Source
}

func (peeff *PassportElementErrorFrontSide) PassportElementErrorType() string {
	if peeff == nil {
		return ""
	}

	return peeff.Type
}

func (peerf *PassportElementErrorReverseSide) PassportElementErrorMessage() string {
	if peerf == nil {
		return ""
	}

	return peerf.Message
}

func (peerf *PassportElementErrorReverseSide) PassportElementErrorSource() string {
	if peerf == nil {
		return ""
	}

	return peerf.Source
}

func (peerf *PassportElementErrorReverseSide) PassportElementErrorType() string {
	if peerf == nil {
		return ""
	}

	return peerf.Type
}

func (pees *PassportElementErrorSelfie) PassportElementErrorMessage() string {
	if pees == nil {
		return ""
	}

	return pees.Message
}

func (pees *PassportElementErrorSelfie) PassportElementErrorSource() string {
	if pees == nil {
		return ""
	}

	return pees.Source
}

func (pees *PassportElementErrorSelfie) PassportElementErrorType() string {
	if pees == nil {
		return ""
	}

	return pees.Type
}

func (peef *PassportElementErrorFile) PassportElementErrorMessage() string {
	if peef == nil {
		return ""
	}

	return peef.Message
}

func (peef *PassportElementErrorFile) PassportElementErrorSource() string {
	if peef == nil {
		return ""
	}

	return peef.Source
}

func (peef *PassportElementErrorFile) PassportElementErrorType() string {
	if peef == nil {
		return ""
	}

	return peef.Type
}

func (peef *PassportElementErrorFiles) PassportElementErrorMessage() string {
	if peef == nil {
		return ""
	}

	return peef.Message
}

func (peef *PassportElementErrorFiles) PassportElementErrorSource() string {
	if peef == nil {
		return ""
	}

	return peef.Source
}

func (peef *PassportElementErrorFiles) PassportElementErrorType() string {
	if peef == nil {
		return ""
	}

	return peef.Type
}

func (peetf *PassportElementErrorTranslationFile) PassportElementErrorMessage() string {
	if peetf == nil {
		return ""
	}

	return peetf.Message
}

func (peetf *PassportElementErrorTranslationFile) PassportElementErrorSource() string {
	if peetf == nil {
		return ""
	}

	return peetf.Source
}

func (peetf *PassportElementErrorTranslationFile) PassportElementErrorType() string {
	if peetf == nil {
		return ""
	}

	return peetf.Type
}

func (peetf *PassportElementErrorTranslationFiles) PassportElementErrorMessage() string {
	if peetf == nil {
		return ""
	}

	return peetf.Message
}

func (peetf *PassportElementErrorTranslationFiles) PassportElementErrorSource() string {
	if peetf == nil {
		return ""
	}

	return peetf.Source
}

func (peetf *PassportElementErrorTranslationFiles) PassportElementErrorType() string {
	if peetf == nil {
		return ""
	}

	return peetf.Type
}

func (peeu *PassportElementErrorUnspecified) PassportElementErrorMessage() string {
	if peeu == nil {
		return ""
	}

	return peeu.Message
}

func (peeu *PassportElementErrorUnspecified) PassportElementErrorSource() string {
	if peeu == nil {
		return ""
	}

	return peeu.Source
}

func (peeu *PassportElementErrorUnspecified) PassportElementErrorType() string {
	if peeu == nil {
		return ""
	}

	return peeu.Type
}

func (pseoos *PassportScopeElementOneOfSeveral) PassportScopeElementTranslation() bool {
	if pseoos == nil {
		return false
	}

	return pseoos.Translation
}

func (pseoos *PassportScopeElementOneOfSeveral) PassportScopeElementSelfie() bool {
	if pseoos == nil {
		return false
	}

	return pseoos.Selfie
}

func (pseo *PassportScopeElementOne) PassportScopeElementTranslation() bool {
	if pseo == nil {
		return false
	}

	return pseo.Translation
}

func (pseo *PassportScopeElementOne) PassportScopeElementSelfie() bool {
	if pseo == nil {
		return false
	}

	return pseo.Selfie
}

func (pd *PersonalDetails) BirthTime() *time.Time {
	if pd == nil || pd.BirthDate == "" {
		return nil
	}

	bt, err := time.Parse("02.01.2006", pd.BirthDate)
	if err != nil {
		return nil
	}

	return &bt
}

func (pd *PersonalDetails) FullName() string {
	if pd == nil {
		return ""
	}

	return pd.FirstName + " " + pd.LastName
}

func (pd *PersonalDetails) FullNameNative() string {
	if pd == nil {
		return ""
	}

	return pd.FirstNameNative + " " + pd.LastNameNative
}

func (sv *SecureValue) HasData() bool {
	return sv != nil && sv.Data != nil
}

func (sv *SecureValue) HasFiles() bool {
	return sv != nil && len(sv.Files) > 0
}

func (sv *SecureValue) HasFrontSide() bool {
	return sv != nil && sv.FrontSide != nil
}

func (sv *SecureValue) HasReverseSide() bool {
	return sv != nil && sv.ReverseSide != nil
}

func (sv *SecureValue) HasSelfie() bool {
	return sv != nil && sv.Selfie != nil
}

func (sv *SecureValue) HasTranslation() bool {
	return sv != nil && len(sv.Translation) > 0
}

// InSet checks that the current sticker in the stickers set.
//
// For uploaded WebP files this return false.
func (s *Sticker) InSet() bool {
	return s != nil && s.SetName != ""
}

// IsWebP check that the current sticker is a WebP file uploaded by user.
func (s *Sticker) IsWebP() bool {
	return s != nil && s.SetName == ""
}

// Set use bot for getting parent StickerSet if SetName is present.
//
// Return nil if current sticker has been uploaded by user as WebP file.
func (s *Sticker) Set(bot *Bot) *StickerSet {
	if s.IsWebP() || bot == nil {
		return nil
	}

	set, err := bot.GetStickerSet(s.SetName)
	if err != nil {
		return nil
	}

	return set
}

func (s *Sticker) HasThumb() bool {
	return s != nil && s.Thumb != nil
}

func (s *Sticker) IsMask() bool {
	return s != nil && s.MaskPosition != nil
}

func (s *Sticker) File() *File {
	if s == nil {
		return nil
	}

	return &File{
		FileID:   s.FileID,
		FileSize: s.FileSize,
	}
}

// NewLongPollingChannel creates channel for receive incoming updates using long
// polling.
func (b *Bot) NewLongPollingChannel(params *GetUpdatesParameters) UpdatesChannel {
	if params == nil {
		params = &GetUpdatesParameters{
			Offset:  0,
			Limit:   100,
			Timeout: 60,
		}
	}

	channel := make(chan Update, params.Limit)
	go func() {
		for {
			updates, err := b.GetUpdates(params)
			if err != nil {
				dlog.Ln(err.Error())
				dlog.Ln("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.ID >= params.Offset {
					params.Offset = update.ID + 1
					channel <- update
				}
			}
		}
	}()

	return channel
}

// NewWebhookChannel creates channel for receive incoming updates via an outgoing webhook.
//
// If cert argument is provided by two strings (["path/to/cert.file", "path/to/cert.key"]), then TLS server will
// be created by this filepaths.
func (b *Bot) NewWebhookChannel(setURL *http.URI, params *SetWebhookParameters, ln net.Listener, cert ...string) (updates UpdatesChannel, shutdown ShutdownFunc) {
	defer http.ReleaseURI(setURL)
	if params == nil {
		params = &SetWebhookParameters{
			URL:            setURL.String(),
			MaxConnections: 40,
		}
	}

	var err error
	channel := make(chan Update, 100)
	handleFunc := func(ctx *http.RequestCtx) {
		dlog.Ln("Request path:", string(ctx.Path()))
		if !bytes.HasPrefix(ctx.Path(), setURL.Path()) {
			dlog.Ln("Unsupported request path:", string(ctx.Path()))
			return
		}
		dlog.Ln("Catched supported request path:", string(ctx.Path()))

		var update Update
		if err = json.UnmarshalFast(ctx.Request.Body(), &update); err != nil {
			return
		}

		channel <- update
	}

	srv := http.Server{
		Name:              b.Username,
		Concurrency:       params.MaxConnections,
		Handler:           handleFunc,
		ReduceMemoryUsage: true,
	}

	go func() {
		switch {
		case len(cert) == 2:
			dlog.Ln("Creating TLS router...")
			err = srv.ServeTLS(ln, cert[0], cert[1])
		default:
			dlog.Ln("Creating simple router...")
			err = srv.Serve(ln)
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	if _, err = b.SetWebhook(params); err != nil {
		log.Fatalln(err.Error())
	}

	return channel, srv.Shutdown
}

// IsMessage checks that the current update is a message creation event.
func (u *Update) IsMessage() bool {
	return u != nil && u.Message != nil
}

// IsEditedMessage checks that the current update is a editing message event.
func (u *Update) IsEditedMessage() bool {
	return u != nil && u.EditedMessage != nil
}

// IsChannelPost checks that the current update is a post channel creation event.
func (u *Update) IsChannelPost() bool {
	return u != nil && u.ChannelPost != nil
}

// IsEditedChannelPost checks that the current update is a editing post channel
// event.
func (u *Update) IsEditedChannelPost() bool {
	return u != nil && u.EditedChannelPost != nil
}

// IsInlineQuery checks that the current update is a inline query update.
func (u *Update) IsInlineQuery() bool {
	return u != nil && u.InlineQuery != nil
}

// IsChosenInlineResult checks that the current update is a chosen inline result
// update.
func (u *Update) IsChosenInlineResult() bool {
	return u != nil && u.ChosenInlineResult != nil
}

// IsCallbackQuery checks that the current update is a callback query update.
func (u *Update) IsCallbackQuery() bool {
	return u != nil && u.CallbackQuery != nil
}

// IsShippingQuery checks that the current update is a shipping query update.
func (u *Update) IsShippingQuery() bool {
	return u != nil && u.ShippingQuery != nil
}

// IsPreCheckoutQuery checks that the current update is a pre checkout query
// update.
func (u *Update) IsPreCheckoutQuery() bool {
	return u != nil && u.PreCheckoutQuery != nil
}

// Type return update type for current update.
func (u *Update) Type() string {
	switch {
	case u.IsCallbackQuery():
		return UpdateCallbackQuery
	case u.IsChannelPost():
		return UpdateChannelPost
	case u.IsChosenInlineResult():
		return UpdateChosenInlineResult
	case u.IsEditedChannelPost():
		return UpdateEditedChannelPost
	case u.IsEditedMessage():
		return UpdateEditedMessage
	case u.IsInlineQuery():
		return UpdateInlineQuery
	case u.IsMessage():
		return UpdateMessage
	case u.IsPreCheckoutQuery():
		return UpdatePreCheckoutQuery
	case u.IsShippingQuery():
		return UpdateShippingQuery
	default:
		return ""
	}
}

// Language parse LanguageCode of current user and returns language.Tag.
func (u *User) Language() language.Tag {
	if u == nil {
		return language.Und
	}

	tag, err := language.Parse(u.LanguageCode)
	if err != nil {
		return language.Und
	}

	return tag
}

// NewPrinter create simple message.Printer with User.Language() by default.
func (u *User) NewPrinter() *message.Printer {
	return message.NewPrinter(u.Language())
}

// FullName returns the full name of user or FirstName if LastName is not
// available.
func (u *User) FullName() string {
	if u == nil {
		return ""
	}

	if u.HasLastName() {
		return u.FirstName + " " + u.LastName
	}

	return u.FirstName
}

// HaveLastName checks what the current user has a LastName.
func (u *User) HasLastName() bool {
	return u != nil && u.LastName != ""
}

// HaveUsername checks what the current user has a username.
func (u *User) HasUsername() bool {
	return u != nil && u.Username != ""
}

func (v *Video) HasThumb() bool {
	return v != nil && v.Thumb != nil
}

func (v *Video) File() *File {
	if v == nil {
		return nil
	}

	return &File{
		FileID:   v.FileID,
		FileSize: v.FileSize,
	}
}

func (vn *VideoNote) HasThumb() bool {
	return vn != nil && vn.Thumb != nil
}

func (vn *VideoNote) File() *File {
	if vn == nil {
		return nil
	}

	return &File{
		FileID:   vn.FileID,
		FileSize: vn.FileSize,
	}
}

func (wi *WebhookInfo) HasURL() bool {
	return wi != nil && wi.URL != ""
}

func (wi *WebhookInfo) LastErrorTime() *time.Time {
	if wi == nil {
		return nil
	}

	led := time.Unix(wi.LastErrorDate, 0)
	return &led
}
