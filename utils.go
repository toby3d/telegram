package telegram

import (
	"fmt"
	"net/url"
	"strconv"
)

// NewForceReply calls the response interface to the message.
func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

// NewInlineMentionURL creates a url.URL for the mention user without username.
func NewInlineMentionURL(userID int) *url.URL {
	link := &url.URL{
		Scheme: SchemeTelegram,
		Path:   "user",
	}

	q := link.Query()
	q.Add("id", strconv.Itoa(userID))
	link.RawQuery = q.Encode()

	return link
}

func NewMarkdownBold(text string) string {
	return fmt.Sprint("*", text, "*")
}

func NewMarkdownItalic(text string) string {
	return fmt.Sprint("_", text, "_")
}

func NewMarkdownURL(text string, link *url.URL) string {
	return fmt.Sprint("[", text, "](", link.String(), ")")
}

func NewMarkdownMention(text string, id int) string {
	link := NewInlineMentionURL(id)
	return NewMarkdownURL(text, link)
}

func NewMarkdownCode(text string) string {
	return fmt.Sprint("`", text, "`")
}

func NewMarkdownCodeBlock(text string) string {
	return fmt.Sprint("```", text, "```")
}

func NewHtmlBold(text string) string {
	return fmt.Sprint("<b>", text, "</b>")
}

func NewHtmlItalic(text string) string {
	return fmt.Sprint("<i>", text, "</i>")
}

func NewHtmlURL(text string, link *url.URL) string {
	return fmt.Sprint(`<a href="`, link.String(), `">`, text, `</a>`)
}

func NewHtmlMention(text string, id int) string {
	link := NewInlineMentionURL(id)
	return NewHtmlURL(text, link)
}

func NewHtmlCode(text string) string {
	return fmt.Sprint("<code>", text, "</code>")
}

func NewHtmlCodeBlock(text string) string {
	return fmt.Sprint("<pre>", text, "</pre>")
}
