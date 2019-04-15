package telegram

import (
	"sort"
	"strings"
	"time"
)

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
