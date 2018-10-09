package telegram

import (
	http "github.com/valyala/fasthttp"
)

// NewForceReply calls the response interface to the message.
func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

// NewInlineMentionURL creates a url.URL for the mention user without username.
func NewInlineMentionURL(userID int) *http.URI {
	link := http.AcquireURI()
	link.SetScheme(SchemeTelegram)
	link.SetPath("user")

	q := link.QueryArgs()
	q.SetUint("id", userID)
	link.SetQueryStringBytes(q.QueryString())

	return link
}

func NewMarkdownBold(text string) string {
	return "*" + text + "*"
}

func NewMarkdownItalic(text string) string {
	return "_" + text + "_"
}

func NewMarkdownURL(text string, link *http.URI) string {
	return "[" + text + "](" + link.String() + ")"
}

func NewMarkdownMention(text string, id int) string {
	return NewMarkdownURL(text, NewInlineMentionURL(id))
}

func NewMarkdownCode(text string) string {
	return "`" + text + "`"
}

func NewMarkdownCodeBlock(text string) string {
	return "```" + text + "```"
}

func NewHtmlBold(text string) string {
	return "<b>" + text + "</b>"
}

func NewHtmlItalic(text string) string {
	return "<i>" + text + "</i>"
}

func NewHtmlURL(text string, link *http.URI) string {
	return `<a href="` + link.String() + `">` + text + `</a>`
}

func NewHtmlMention(text string, id int) string {
	return NewHtmlURL(text, NewInlineMentionURL(id))
}

func NewHtmlCode(text string) string {
	return "<code>" + text + "</code>"
}

func NewHtmlCodeBlock(text string) string {
	return "<pre>" + text + "</pre>"
}
