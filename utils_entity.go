package telegram

import (
	"fmt"
	"net/url"
)

// ParseURL selects URL from entered text of message and parse it as url.URL.
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

// IsBold checks that the current entity is a bold tag.
func (entity *MessageEntity) IsBold() bool {
	return entity != nil && entity.Type == EntityBold
}

// IsBotCommand checks that the current entity is a bot command.
func (entity *MessageEntity) IsBotCommand() bool {
	return entity != nil && entity.Type == EntityBotCommand
}

// IsCode checks that the current entity is a code tag.
func (entity *MessageEntity) IsCode() bool {
	return entity != nil && entity.Type == EntityCode
}

// IsEmail checks that the current entity is a email.
func (entity *MessageEntity) IsEmail() bool {
	return entity != nil && entity.Type == EntityEmail
}

// IsHashtag checks that the current entity is a hashtag.
func (entity *MessageEntity) IsHashtag() bool {
	return entity != nil && entity.Type == EntityHashtag
}

// IsItalic checks that the current entity is a italic tag.
func (entity *MessageEntity) IsItalic() bool {
	return entity != nil && entity.Type == EntityItalic
}

// IsMention checks that the current entity is a username mention.
func (entity *MessageEntity) IsMention() bool {
	return entity != nil && entity.Type == EntityMention
}

// IsPre checks that the current entity is a pre tag.
func (entity *MessageEntity) IsPre() bool {
	return entity != nil && entity.Type == EntityPre
}

// IsTextLink checks that the current entity is a text link.
func (entity *MessageEntity) IsTextLink() bool {
	return entity != nil && entity.Type == EntityTextLink
}

// IsTextMention checks that the current entity is a mention without username.
func (entity *MessageEntity) IsTextMention() bool {
	return entity != nil && entity.Type == EntityTextMention
}

// IsURL checks that the current entity is a URL
func (entity *MessageEntity) IsURL() bool {
	return entity != nil && entity.Type == EntityURL
}

// TextLink parse current text link entity as url.URL.
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
