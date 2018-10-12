package telegram

func (d *Document) HasThumb() bool {
	return d != nil && d.Thumb != nil
}

func (d *Document) File() *File {
	return &File{
		FileID:   d.FileID,
		FileSize: d.FileSize,
	}
}
