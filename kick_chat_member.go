package telegram

import json "github.com/pquerna/ffjson/ffjson"

// KickChatMemberParameters represents data for KickChatMember method.
type KickChatMemberParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	UntilDate int64 `json:"until_date"`
	UserID    int   `json:"user_id"`
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
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodKickChatMember)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
