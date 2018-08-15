package telegram

func (cir *ChosenInlineResult) HasLocation() bool {
	return cir != nil && cir.Location != nil
}
