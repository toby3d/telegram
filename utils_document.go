package telegram

func (d *Document) HasThumb() bool {
	return d != nil && d.Thumb != nil
}
