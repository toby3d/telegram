package telegram

func (chat *Chat) IsPrivate() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatPrivate
}

func (chat *Chat) IsGroup() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatGroup
}

func (chat *Chat) IsSuperGroup() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatSuperGroup
}

func (chat *Chat) IsChannel() bool {
	if chat == nil {
		return false
	}

	return chat.Type == ChatChannel
}

func (chat *Chat) HasPinnedMessage() bool {
	if chat == nil {
		return false
	}

	return chat.PinnedMessage != nil
}
