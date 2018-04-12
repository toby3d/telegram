package telegram

func NewInputTextMessageContent(messageText string) *InputTextMessageContent {
	return &InputTextMessageContent{
		MessageText: messageText,
	}
}

func NewInputLocationMessageContent(latitude, longitude float32) *InputLocationMessageContent {
	return &InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func NewInputVenueMessageContent(latitude, longitude float32, title, address string) *InputVenueMessageContent {
	return &InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

func NewInputContactMessageContent(phoneNumber, firstName string) *InputContactMessageContent {
	return &InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

func NewInputMediaPhoto(media string) *InputMediaPhoto {
	return &InputMediaPhoto{
		Type:  TypePhoto,
		Media: media,
	}
}

func NewInputMediaVideo(media string) *InputMediaVideo {
	return &InputMediaVideo{
		Type:  TypeVideo,
		Media: media,
	}
}
