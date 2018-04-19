package telegram

import (
	"strings"
	"time"
)

// IsCommand checks that the current message is a bot command.
func (msg *Message) IsCommand() bool {
	if !msg.IsText() || !msg.HasEntities() {
		return false
	}

	entity := msg.Entities[0]
	return entity.IsBotCommand() && entity.Offset == 0
}

// IsCommandEqual checks that the current message is a specific bot command.
func (msg *Message) IsCommandEqual(command string) bool {
	return msg.IsCommand() && strings.EqualFold(msg.Command(), command)
}

// Command returns identifier of the bot command without bot username, if it was
// available
func (msg *Message) Command() string {
	if !msg.IsCommand() {
		return ""
	}

	return strings.Split(msg.RawCommand(), "@")[0]
}

// RawCommand returns identifier of the bot command with bot username, if it was
// available
func (msg *Message) RawCommand() string {
	if !msg.IsCommand() {
		return ""
	}

	return string([]rune(msg.Text)[1:msg.Entities[0].Length])
}

// HasCommandArgument checks that the current command message contains argument.
func (msg *Message) HasCommandArgument() bool {
	if !msg.IsCommand() {
		return false
	}

	entity := msg.Entities[0]
	if !entity.IsBotCommand() {
		return false
	}

	return len([]rune(msg.Text)) != entity.Length
}

// CommandArgument returns raw command argument.
func (msg *Message) CommandArgument() string {
	if !msg.HasCommandArgument() {
		return ""
	}

	return string([]rune(msg.Text)[msg.Entities[0].Length+1:])
}

// IsReply checks that the current message is a reply on other message.
func (msg *Message) IsReply() bool {
	return msg != nil && msg.ReplyToMessage != nil
}

// IsForward checks that the current message is a forward of other message.
func (msg *Message) IsForward() bool {
	return msg != nil && msg.ForwardFrom != nil
}

// Time parse current message Date and returns time.Time.
func (msg *Message) Time() time.Time {
	if msg == nil {
		return time.Time{}
	}

	return time.Unix(msg.Date, 0)
}

// ForwardTime parse current message ForwardDate and returns time.Time.
func (msg *Message) ForwardTime() time.Time {
	if msg == nil {
		return time.Time{}
	}

	return time.Unix(msg.ForwardDate, 0)
}

// EditTime parse current message EditDate and returns time.Time.
func (msg *Message) EditTime() time.Time {
	var t time.Time
	if msg == nil || !msg.HasBeenEdited() {
		return t
	}

	return time.Unix(msg.EditDate, 0)
}

// HasBeenEdited checks that the current message has been edited.
func (msg *Message) HasBeenEdited() bool {
	return msg != nil && msg.EditDate > 0
}

// IsText checks that the current message is just a text message.
func (msg *Message) IsText() bool {
	return msg != nil && msg.Text != ""
}

// IsAudio checks that the current message is a audio.
func (msg *Message) IsAudio() bool {
	return !msg.IsText() && msg.Audio != nil
}

// IsDocument checks that the current message is a document.
func (msg *Message) IsDocument() bool {
	return !msg.IsText() && msg.Document != nil
}

// IsGame checks that the current message is a game.
func (msg *Message) IsGame() bool {
	return !msg.IsText() && msg.Game != nil
}

// IsPhoto checks that the current message is a photo.
func (msg *Message) IsPhoto() bool {
	return !msg.IsText() && len(msg.Photo) > 0
}

// IsSticker checks that the current message is a sticker.
func (msg *Message) IsSticker() bool {
	return !msg.IsText() && msg.Sticker != nil
}

// IsVideo checks that the current message is a video.
func (msg *Message) IsVideo() bool {
	return !msg.IsText() && msg.Video != nil
}

// IsVoice checks that the current message is a voice.
func (msg *Message) IsVoice() bool {
	return !msg.IsText() && msg.Voice != nil
}

// IsVideoNote checks that the current message is a video note.
func (msg *Message) IsVideoNote() bool {
	return !msg.IsText() && msg.VideoNote != nil
}

// IsContact checks that the current message is a contact.
func (msg *Message) IsContact() bool {
	return !msg.IsText() && msg.Contact != nil
}

// IsLocation checks that the current message is a location.
func (msg *Message) IsLocation() bool {
	return !msg.IsText() && msg.Location != nil
}

// IsVenue checks that the current message is a venue.
func (msg *Message) IsVenue() bool {
	return !msg.IsText() && msg.Venue != nil
}

// IsNewChatMembersEvent checks that the current message is a event of entry of
// new members.
func (msg *Message) IsNewChatMembersEvent() bool {
	return !msg.IsText() && len(msg.NewChatMembers) > 0
}

// IsLeftChatMemberEvent checks that the current message is a event of members
// exit.
func (msg *Message) IsLeftChatMemberEvent() bool {
	return !msg.IsText() && msg.LeftChatMember != nil
}

// IsNewChatTitleEvent checks that the current message is a event of setting a
// new chat title.
func (msg *Message) IsNewChatTitleEvent() bool {
	return !msg.IsText() && msg.NewChatTitle != ""
}

// IsNewChatPhotoEvent checks that the current message is a event of setting a
// new chat avatar.
func (msg *Message) IsNewChatPhotoEvent() bool {
	return !msg.IsText() && len(msg.NewChatPhoto) > 0
}

// IsDeleteChatPhotoEvent checks that the current message is a event of deleting
// a chat avatar.
func (msg *Message) IsDeleteChatPhotoEvent() bool {
	return !msg.IsText() && msg.DeleteChatPhoto
}

// IsGroupChatCreatedEvent checks that the current message is a event of creating
// a new group.
func (msg *Message) IsGroupChatCreatedEvent() bool {
	return !msg.IsText() && msg.GroupChatCreated
}

// IsSupergroupChatCreatedEvent checks that the current message is a event of
// creating a new supergroup.
func (msg *Message) IsSupergroupChatCreatedEvent() bool {
	return !msg.IsText() && msg.SupergroupChatCreated
}

// IsChannelChatCreatedEvent checks that the current message is a event of
// creating a new channel.
func (msg *Message) IsChannelChatCreatedEvent() bool {
	return !msg.IsText() && msg.ChannelChatCreated
}

// IsPinnedMessage checks that the current message is a event of pinning another
// message.
func (msg *Message) IsPinnedMessage() bool {
	return !msg.IsText() && msg.PinnedMessage != nil
}

// IsInvoice checks that the current message is a invoice.
func (msg *Message) IsInvoice() bool {
	return !msg.IsText() && msg.Invoice != nil
}

// IsSuccessfulPayment checks that the current message is a event of successful
// payment.
func (msg *Message) IsSuccessfulPayment() bool {
	return !msg.IsText() && msg.SuccessfulPayment != nil
}

// HasEntities checks that the current message contains entities.
func (msg *Message) HasEntities() bool {
	return msg.IsText() && len(msg.Entities) > 0
}

// HasCaptionEntities checks that the current media contains entities in caption.
func (msg *Message) HasCaptionEntities() bool {
	return !msg.IsText() && len(msg.CaptionEntities) > 0
}

// HasMentions checks that the current message contains mentions.
func (msg *Message) HasMentions() bool {
	if !msg.HasEntities() {
		return false
	}

	for _, entity := range msg.Entities {
		if entity.IsMention() || entity.IsTextMention() {
			return true
		}
	}

	return false
}

// HasCaptionMentions checks that the current media contains mentions in caption.
func (msg *Message) HasCaptionMentions() bool {
	if !msg.HasCaptionEntities() {
		return false
	}

	for _, entity := range msg.CaptionEntities {
		if entity.IsMention() || entity.IsTextMention() {
			return true
		}
	}

	return false
}

// HasCaption checks that the current media has caption.
func (msg *Message) HasCaption() bool {
	return !msg.IsText() && msg.Caption != ""
}

// HasAuthorSignature checks that the current channel post has author signature.
func (msg *Message) HasAuthorSignature() bool {
	return msg != nil && msg.AuthorSignature != ""
}
