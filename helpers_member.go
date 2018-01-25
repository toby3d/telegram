package telegram

import "time"

func (member *ChatMember) IsCreator() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusCreator
}

func (member *ChatMember) IsAdministrator() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusAdministrator
}

func (member *ChatMember) IsMember() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusMember
}

func (member *ChatMember) IsRestricted() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusRestricted
}

func (member *ChatMember) IsLeft() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusLeft
}

func (member *ChatMember) IsKicked() bool {
	if member == nil {
		return false
	}

	return member.Status == StatusKicked
}

func (member *ChatMember) UntilTime() time.Time {
	if member == nil {
		return time.Time{}
	}

	return time.Unix(member.UntilDate, 0)
}
