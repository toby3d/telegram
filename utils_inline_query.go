package telegram

// HasLocation checks what current InlineQuery has Location info.
func (iq *InlineQuery) HasLocation() bool {
	return iq != nil && iq.Location != nil
}

// HasOffset checks what current InlineQuery has Offset.
func (iq *InlineQuery) HasOffset() bool {
	return iq != nil && iq.Offset != ""
}

// HasQuery checks what current InlineQuery has Query string.
func (iq *InlineQuery) HasQuery() bool {
	return iq != nil && iq.Query != ""
}
