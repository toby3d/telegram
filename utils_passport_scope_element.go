package telegram

func (pseoos *PassportScopeElementOneOfSeveral) PassportScopeElementTranslation() bool {
	if pseoos == nil {
		return false
	}

	return pseoos.Translation
}

func (pseoos *PassportScopeElementOneOfSeveral) PassportScopeElementSelfie() bool {
	if pseoos == nil {
		return false
	}

	return pseoos.Selfie
}

func (pseo *PassportScopeElementOne) PassportScopeElementTranslation() bool {
	if pseo == nil {
		return false
	}

	return pseo.Translation
}

func (pseo *PassportScopeElementOne) PassportScopeElementSelfie() bool {
	if pseo == nil {
		return false
	}

	return pseo.Selfie
}
