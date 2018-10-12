package telegram

func (sv *SecureValue) HasData() bool {
	return sv != nil && sv.Data != nil
}

func (sv *SecureValue) HasFiles() bool {
	return sv != nil && len(sv.Files) > 0
}

func (sv *SecureValue) HasFrontSide() bool {
	return sv != nil && sv.FrontSide != nil
}

func (sv *SecureValue) HasReverseSide() bool {
	return sv != nil && sv.ReverseSide != nil
}

func (sv *SecureValue) HasSelfie() bool {
	return sv != nil && sv.Selfie != nil
}

func (sv *SecureValue) HasTranslation() bool {
	return sv != nil && len(sv.Translation) > 0
}
