package telegram

func (vn *VideoNote) HasThumb() bool {
	return vn != nil && vn.Thumb != nil
}

func (vn *VideoNote) File() *File {
	return &File{
		FileID:   vn.FileID,
		FileSize: vn.FileSize,
	}
}
