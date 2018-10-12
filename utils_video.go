package telegram

func (v *Video) HasThumb() bool {
	return v != nil && v.Thumb != nil
}

func (v *Video) File() *File {
	return &File{
		FileID:   v.FileID,
		FileSize: v.FileSize,
	}
}
