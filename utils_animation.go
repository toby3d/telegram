package telegram

func (a *Animation) HasThumb() bool {
	return a != nil && a.Thumb != nil
}
