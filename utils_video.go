package telegram

func (v *Video) HasThumb() bool {
	return v != nil && v.Thumb != nil
}
