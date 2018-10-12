package telegram

func (a *Animation) HasThumb() bool {
	return a != nil && a.Thumb != nil
}

func (a *Animation) File() *File {
	return &File{
		FileID:   a.FileID,
		FileSize: a.FileSize,
	}
}
