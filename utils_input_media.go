package telegram

func (ima *InputMediaAnimation) File() string {
	if ima == nil {
		return ""
	}

	return ima.Media
}

func (ima *InputMediaAnimation) InputMediaCaption() string {
	if ima == nil {
		return ""
	}

	return ima.Caption
}

func (ima *InputMediaAnimation) InputMediaParseMode() string {
	if ima == nil {
		return ""
	}

	return ima.ParseMode
}

func (ima *InputMediaAnimation) InputMediaType() string {
	if ima == nil {
		return ""
	}

	return ima.Type
}

func (imd *InputMediaDocument) File() string {
	if imd == nil {
		return ""
	}

	return imd.Media
}

func (imd *InputMediaDocument) InputMediaCaption() string {
	if imd == nil {
		return ""
	}

	return imd.Caption
}

func (imd *InputMediaDocument) InputMediaParseMode() string {
	if imd == nil {
		return ""
	}

	return imd.ParseMode
}

func (imd *InputMediaDocument) InputMediaType() string {
	if imd == nil {
		return ""
	}

	return imd.Type
}

func (ima *InputMediaAudio) File() string {
	if ima == nil {
		return ""
	}

	return ima.Media
}

func (ima *InputMediaAudio) InputMediaCaption() string {
	if ima == nil {
		return ""
	}

	return ima.Caption
}

func (ima *InputMediaAudio) InputMediaParseMode() string {
	if ima == nil {
		return ""
	}

	return ima.ParseMode
}

func (ima *InputMediaAudio) InputMediaType() string {
	if ima == nil {
		return ""
	}

	return ima.Type
}

func (imp *InputMediaPhoto) File() string {
	if imp == nil {
		return ""
	}

	return imp.Media
}

func (imp *InputMediaPhoto) InputMediaCaption() string {
	if imp == nil {
		return ""
	}

	return imp.Caption
}

func (imp *InputMediaPhoto) InputMediaParseMode() string {
	if imp == nil {
		return ""
	}

	return imp.ParseMode
}

func (imp *InputMediaPhoto) InputMediaType() string {
	if imp == nil {
		return ""
	}

	return imp.Type
}

func (imv *InputMediaVideo) File() string {
	if imv == nil {
		return ""
	}

	return imv.Media
}

func (imv *InputMediaVideo) InputMediaCaption() string {
	if imv == nil {
		return ""
	}

	return imv.Caption
}

func (imv *InputMediaVideo) InputMediaParseMode() string {
	if imv == nil {
		return ""
	}

	return imv.ParseMode
}

func (imv *InputMediaVideo) InputMediaType() string {
	if imv == nil {
		return ""
	}

	return imv.Type
}
