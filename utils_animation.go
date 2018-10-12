package telegram

func (a *Animation) HasThumb() bool {
	return a != nil && a.Thumb != nil
}

func (a *Animation) File() *File {
	if a == nil {
		return nil
	}

	return &File{
		FileID:   a.FileID,
		FileSize: a.FileSize,
	}
}
