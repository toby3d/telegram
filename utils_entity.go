package telegram

import (
	"fmt"
	"net/url"
)

func (entity *MessageEntity) ParseURL(messageText string) *url.URL {
	if entity == nil || !entity.IsURL() || messageText == "" {
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
	return entity != nil && entity.Type == EntityBold
}

func (entity *MessageEntity) IsBotCommand() bool {
	return entity != nil && entity.Type == EntityBotCommand
}

func (entity *MessageEntity) IsCode() bool {
	return entity != nil && entity.Type == EntityCode
}

func (entity *MessageEntity) IsEmail() bool {
	return entity != nil && entity.Type == EntityEmail
}

func (entity *MessageEntity) IsHashtag() bool {
	return entity != nil && entity.Type == EntityHashtag
}

func (entity *MessageEntity) IsItalic() bool {
	return entity != nil && entity.Type == EntityItalic
}

func (entity *MessageEntity) IsMention() bool {
	return entity != nil && entity.Type == EntityMention
}

func (entity *MessageEntity) IsPre() bool {
	return entity != nil && entity.Type == EntityPre
}

func (entity *MessageEntity) IsTextLink() bool {
	return entity != nil && entity.Type == EntityTextLink
}

func (entity *MessageEntity) IsTextMention() bool {
	return entity != nil && entity.Type == EntityTextMention
}

func (entity *MessageEntity) IsURL() bool {
	return entity != nil && entity.Type == EntityURL
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
