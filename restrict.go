package telegram

type RestrictChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Unique identifier of the target user
	UserID int `json:"user_id"`

	// New user permissions
	Permissions ChatPermissions `json:"permissions"`

	// Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days
	// or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have
// the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on
// success.
func (b *Bot) RestrictChatMember(params RestrictChatMemberParameters) (bool, error) {
	dst, err := parser.Marshal(&params)
	if err != nil {
		return false, err
	}

	resp, err := b.request(dst, MethodRestrictChatMember)
	if err != nil {
		return false, err
	}

	var ok bool
	err = parser.Unmarshal(resp.Result, &ok)
	return ok, err
}
