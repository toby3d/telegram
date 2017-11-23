package telegram

func (chat *Chat) IsPrivate() bool {
	return chat.Type == ChatPrivate
}

func (chat *Chat) IsGroup() bool {
	return chat.Type == ChatGroup
}

func (chat *Chat) IsSuperGroup() bool {
	return chat.Type == ChatSuperGroup
}

func (chat *Chat) IsChannel() bool {
	return chat.Type == ChatChannel
}
