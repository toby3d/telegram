package telegram

func (vn *VideoNote) HasThumb() bool {
	return vn != nil && vn.Thumb != nil
}
