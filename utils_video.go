package telegram

func (v *Video) HasThumb() bool {
	return v != nil && v.Thumb != nil
}

func (v *Video) File() *File {
	if v == nil {
		return nil
	}

	return &File{
		FileID:   v.FileID,
		FileSize: v.FileSize,
	}
}
