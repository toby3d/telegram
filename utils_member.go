package telegram

import "time"

func (member *ChatMember) IsCreator() bool {
	return member != nil && member.Status == StatusCreator
}

func (member *ChatMember) IsAdministrator() bool {
	return member != nil && member.Status == StatusAdministrator
}

func (member *ChatMember) IsMember() bool {
	return member != nil && member.Status == StatusMember
}

func (member *ChatMember) IsRestricted() bool {
	return member != nil && member.Status == StatusRestricted
}

func (member *ChatMember) IsLeft() bool {
	return member != nil && member.Status == StatusLeft
}

func (member *ChatMember) IsKicked() bool {
	return member != nil && member.Status == StatusKicked
}

func (member *ChatMember) UntilTime() time.Time {
	if member == nil {
		return time.Time{}
	}

	return time.Unix(member.UntilDate, 0)
}
