package telegram

func (iq *InlineQuery) HasLocation() bool {
	return iq != nil && iq.Location != nil
}
