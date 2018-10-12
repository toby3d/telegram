package telegram

import (
	"strings"

	http "github.com/valyala/fasthttp"
)

// ParseURL selects URL from entered text of message and parse it as fasthttp.URI.
func (e *MessageEntity) ParseURL(messageText string) *http.URI {
	if e == nil || !e.IsURL() || messageText == "" {
		return nil
	}

	from := e.Offset
	to := from + e.Length
	text := []rune(messageText)
	if len(text) < to {
		return nil
	}

	link := http.AcquireURI()
	link.Update(string(text[from:to]))

	return link
}

// IsBold checks that the current entity is a bold tag.
func (e *MessageEntity) IsBold() bool {
	return e != nil && strings.EqualFold(e.Type, EntityBold)
}

// IsBotCommand checks that the current entity is a bot command.
func (e *MessageEntity) IsBotCommand() bool {
	return e != nil && strings.EqualFold(e.Type, EntityBotCommand)
}

// IsCode checks that the current entity is a code tag.
func (e *MessageEntity) IsCode() bool {
	return e != nil && strings.EqualFold(e.Type, EntityCode)
}

// IsEmail checks that the current entity is a email.
func (e *MessageEntity) IsEmail() bool {
	return e != nil && strings.EqualFold(e.Type, EntityEmail)
}

// IsHashtag checks that the current entity is a hashtag.
func (e *MessageEntity) IsHashtag() bool {
	return e != nil && strings.EqualFold(e.Type, EntityHashtag)
}

// IsItalic checks that the current entity is a italic tag.
func (e *MessageEntity) IsItalic() bool {
	return e != nil && strings.EqualFold(e.Type, EntityItalic)
}

// IsMention checks that the current entity is a username mention.
func (e *MessageEntity) IsMention() bool {
	return e != nil && strings.EqualFold(e.Type, EntityMention)
}

// IsPre checks that the current entity is a pre tag.
func (e *MessageEntity) IsPre() bool {
	return e != nil && strings.EqualFold(e.Type, EntityPre)
}

// IsTextLink checks that the current entity is a text link.
func (e *MessageEntity) IsTextLink() bool {
	return e != nil && strings.EqualFold(e.Type, EntityTextLink)
}

// IsTextMention checks that the current entity is a mention without username.
func (e *MessageEntity) IsTextMention() bool {
	return e != nil && strings.EqualFold(e.Type, EntityTextMention)
}

// IsURL checks that the current entity is a URL
func (e *MessageEntity) IsURL() bool {
	return e != nil && strings.EqualFold(e.Type, EntityURL)
}

// TextLink parse current text link entity as fasthttp.URI.
func (e *MessageEntity) TextLink() *http.URI {
	if e == nil {
		return nil
	}

	link := http.AcquireURI()
	link.Update(e.URL)

	return link
}
