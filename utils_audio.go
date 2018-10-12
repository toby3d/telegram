package telegram

func (a *Audio) FullName(sep string) (name string) {
	if !a.HasTitle() {
		return
	}

	if a.HasPerformer() {
		name += a.Performer + sep
	}

	name += a.Title
	return
}

func (a *Audio) HasPerformer() bool {
	return a != nil && a.Performer != ""
}

func (a *Audio) HasTitle() bool {
	return a != nil && a.Title != ""
}

func (a *Audio) HasThumb() bool {
	return a != nil && a.Thumb != nil
}

func (a *Audio) File() *File {
	return &File{
		FileID:   a.FileID,
		FileSize: a.FileSize,
	}
}
