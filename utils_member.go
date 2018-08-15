package telegram

import (
	"strings"
	"time"
)

// IsCreator checks that current member is creator.
func (m *ChatMember) IsCreator() bool {
	return m != nil && strings.EqualFold(m.Status, StatusCreator)
}

// IsAdministrator checks that current member is administrator.
func (m *ChatMember) IsAdministrator() bool {
	return m != nil && strings.EqualFold(m.Status, StatusAdministrator)
}

// IsMember checks that current member is a m.
func (m *ChatMember) IsMember() bool {
	return m != nil && strings.EqualFold(m.Status, StatusMember)
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
