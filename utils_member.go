package telegram

import "time"

// IsCreator checks that current member is creator.
func (member *ChatMember) IsCreator() bool {
	return member != nil && member.Status == StatusCreator
}

// IsAdministrator checks that current member is administrator.
func (member *ChatMember) IsAdministrator() bool {
	return member != nil && member.Status == StatusAdministrator
}

// IsMember checks that current member is a member.
func (member *ChatMember) IsMember() bool {
	return member != nil && member.Status == StatusMember
}

// IsRestricted checks that current member has been restricted.
func (member *ChatMember) IsRestricted() bool {
	return member != nil && member.Status == StatusRestricted
}

// IsLeft checks that current member has left the chat.
func (member *ChatMember) IsLeft() bool {
	return member != nil && member.Status == StatusLeft
}

// IsKicked checks that current member has been kicked.
func (member *ChatMember) IsKicked() bool {
	return member != nil && member.Status == StatusKicked
}

// UntilTime parse UntilDate of restrictions and returns time.Time.
func (member *ChatMember) UntilTime() time.Time {
	if member == nil {
		return time.Time{}
	}

	return time.Unix(member.UntilDate, 0)
}
