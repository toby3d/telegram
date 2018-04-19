package telegram

import (
	"strings"
	"time"
)

func (msg *Message) IsCommand() bool {
	if !msg.IsText() || !msg.HasEntities() {
		return false
	}

	entity := msg.Entities[0]
	return entity.IsBotCommand() && entity.Offset == 0
}

func (msg *Message) IsCommandEqual(command string) bool {
	return msg.IsCommand() && strings.EqualFold(msg.Command(), command)
}

func (msg *Message) Command() string {
	if !msg.IsCommand() {
		return ""
	}

	return strings.Split(msg.RawCommand(), "@")[0]
}

func (msg *Message) RawCommand() string {
	if !msg.IsCommand() {
		return ""
	}

	return string([]rune(msg.Text)[1:msg.Entities[0].Length])
}

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

func (msg *Message) CommandArgument() string {
	if !msg.HasCommandArgument() {
		return ""
	}

	return string([]rune(msg.Text)[msg.Entities[0].Length+1:])
}

func (msg *Message) IsReply() bool {
	return msg != nil && msg.ReplyToMessage != nil
}

func (msg *Message) IsForward() bool {
	return msg != nil && msg.ForwardFrom != nil
}

func (msg *Message) Time() time.Time {
	if msg == nil {
		return time.Time{}
	}

	return time.Unix(msg.Date, 0)
}

func (msg *Message) ForwardTime() time.Time {
	if msg == nil {
		return time.Time{}
	}

	return time.Unix(msg.ForwardDate, 0)
}

func (msg *Message) EditTime() time.Time {
	var t time.Time
	if msg == nil || !msg.HasBeenEdited() {
		return t
	}

	return time.Unix(msg.EditDate, 0)
}

func (msg *Message) HasBeenEdited() bool {
	return msg != nil && msg.EditDate > 0
}

func (msg *Message) IsText() bool {
	return msg != nil && msg.Text != ""
}

func (msg *Message) IsAudio() bool {
	return !msg.IsText() && msg.Audio != nil
}

func (msg *Message) IsDocument() bool {
	return !msg.IsText() && msg.Document != nil
}

func (msg *Message) IsGame() bool {
	return !msg.IsText() && msg.Game != nil
}

func (msg *Message) IsPhoto() bool {
	return !msg.IsText() && len(msg.Photo) > 0
}

func (msg *Message) IsSticker() bool {
	return !msg.IsText() && msg.Sticker != nil
}

func (msg *Message) IsVideo() bool {
	return !msg.IsText() && msg.Video != nil
}

func (msg *Message) IsVoice() bool {
	return !msg.IsText() && msg.Voice != nil
}

func (msg *Message) IsVideoNote() bool {
	return !msg.IsText() && msg.VideoNote != nil
}

func (msg *Message) IsContact() bool {
	return !msg.IsText() && msg.Contact != nil
}

func (msg *Message) IsLocation() bool {
	return !msg.IsText() && msg.Location != nil
}

func (msg *Message) IsVenue() bool {
	return !msg.IsText() && msg.Venue != nil
}

func (msg *Message) IsNewChatMembersEvent() bool {
	return !msg.IsText() && len(msg.NewChatMembers) > 0
}

func (msg *Message) IsLeftChatMemberEvent() bool {
	return !msg.IsText() && msg.LeftChatMember != nil
}

func (msg *Message) IsNewChatTitleEvent() bool {
	return !msg.IsText() && msg.NewChatTitle != ""
}

func (msg *Message) IsNewChatPhotoEvent() bool {
	return !msg.IsText() && len(msg.NewChatPhoto) > 0
}

func (msg *Message) IsDeleteChatPhotoEvent() bool {
	return !msg.IsText() && msg.DeleteChatPhoto
}

func (msg *Message) IsGroupChatCreatedEvent() bool {
	return !msg.IsText() && msg.GroupChatCreated
}

func (msg *Message) IsSupergroupChatCreatedEvent() bool {
	return !msg.IsText() && msg.SupergroupChatCreated
}

func (msg *Message) IsChannelChatCreatedEvent() bool {
	return !msg.IsText() && msg.ChannelChatCreated
}

func (msg *Message) IsPinnedMessage() bool {
	return !msg.IsText() && msg.PinnedMessage != nil
}

func (msg *Message) IsInvoice() bool {
	return !msg.IsText() && msg.Invoice != nil
}

func (msg *Message) IsSuccessfulPayment() bool {
	return !msg.IsText() && msg.SuccessfulPayment != nil
}

func (msg *Message) HasEntities() bool {
	return msg.IsText() && len(msg.Entities) > 0
}

func (msg *Message) HasCaptionEntities() bool {
	return !msg.IsText() && len(msg.CaptionEntities) > 0
}

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

func (msg *Message) HasCaption() bool {
	return !msg.IsText() && msg.Caption != ""
}

func (msg *Message) HasAuthorSignature() bool {
	return msg != nil && msg.AuthorSignature != ""
}
