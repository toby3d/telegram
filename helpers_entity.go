package telegram

import "net/url"

func (entity *MessageEntity) ParseURL() (*url.URL, error) {
	if entity != nil {
		return nil, nil
	}

	if entity.IsTextLink() {
		return url.Parse(entity.URL)
	}

	return nil, nil
}

func (entity *MessageEntity) IsBold() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityBold
}

func (entity *MessageEntity) IsBotCommand() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityBotCommand
}

func (entity *MessageEntity) IsCode() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityCode
}

func (entity *MessageEntity) IsEmail() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityEmail
}

func (entity *MessageEntity) IsHashTag() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityHashtag
}

func (entity *MessageEntity) IsItalic() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityItalic
}

func (entity *MessageEntity) IsMention() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityMention
}

func (entity *MessageEntity) IsPre() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityPre
}

func (entity *MessageEntity) IsTextLink() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityTextLink
}

func (entity *MessageEntity) IsTextMention() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityTextMention
}

func (entity *MessageEntity) IsURL() bool {
	if entity == nil {
		return false
	}

	return entity.Type == EntityURL
}
