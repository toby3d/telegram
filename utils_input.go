package telegram

// NewInputTextMessageContent creates a new text of message.
func NewInputTextMessageContent(messageText string) *InputTextMessageContent {
	return &InputTextMessageContent{
		MessageText: messageText,
	}
}

// NewInputLocationMessageContent creates a new location.
func NewInputLocationMessageContent(latitude, longitude float32) *InputLocationMessageContent {
	return &InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewInputVenueMessageContent creates a new venue.
func NewInputVenueMessageContent(latitude, longitude float32, title, address string) *InputVenueMessageContent {
	return &InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// NewInputContactMessageContent creates a new contact.
func NewInputContactMessageContent(phoneNumber, firstName string) *InputContactMessageContent {
	return &InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewInputMediaPhoto creates a new photo in media album.
func NewInputMediaPhoto(media string) *InputMediaPhoto {
	return &InputMediaPhoto{
		Type:  TypePhoto,
		Media: media,
	}
}

// NewInputMediaVideo creates a new video in media album.
func NewInputMediaVideo(media string) *InputMediaVideo {
	return &InputMediaVideo{
		Type:  TypeVideo,
		Media: media,
	}
}
