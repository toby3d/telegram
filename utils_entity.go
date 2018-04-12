package telegram

import (
	"fmt"
	"net/url"
)

func (entity *MessageEntity) ParseURL(messageText string) *url.URL {
	if entity == nil {
		return nil
	}

	if !entity.IsURL() {
		return nil
	}

	if messageText == "" {
		return nil
	}

	from := entity.Offset
	to := from + entity.Length
	text := []rune(messageText)
	if len(text) < to {
		return nil
	}

	link, err := url.Parse(string(text[from:to]))
	if err == nil && link.Scheme == "" {
		link, err = url.Parse(fmt.Sprint("http://", link))
	}
	if err != nil {
		return nil
	}

	return link
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

func (entity *MessageEntity) IsHashtag() bool {
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

func (entity *MessageEntity) TextLink() *url.URL {
	if entity == nil {
		return nil
	}

	link, err := url.Parse(entity.URL)
	if err != nil {
		return nil
	}

	return link
}
