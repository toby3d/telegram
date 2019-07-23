//go:generate ffjson $GOFILE
package telegram

import json "github.com/pquerna/ffjson/ffjson"

// KickChatMemberParameters represents data for KickChatMember method.
type KickChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Unique identifier of the target user
	UserID int `json:"user_id"`

	// Date when the user will be unbanned, unix time. If user is banned for
	// more than 366 days or less than 30 seconds from the current time they
	// are considered to be banned forever
	UntilDate int64 `json:"until_date"`
}

// KickChatMember kick a user from a group, a supergroup or a channel. In the case of supergroups and
// channels, the user will not be able to return to the group on their own using invite links, etc.,
// unless unbanned first. The bot must be an administrator in the chat for this to work and must have
// the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups), this method will only work if the 'All Members Are
// Admins' setting is off in the target group. Otherwise members may only be removed by the group's
// creator or by the member that added them.
func (bot *Bot) KickChatMember(params *KickChatMemberParameters) (ok bool, err error) {
	dst, err := json.MarshalFast(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodKickChatMember)
	if err != nil {
		return
	}

	err = json.UnmarshalFast(*resp.Result, &ok)
	return
}
