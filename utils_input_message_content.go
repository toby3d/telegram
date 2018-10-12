package telegram

func (itmc *InputTextMessageContent) IsInputMessageContent() bool { return true }

func (ilmc *InputLocationMessageContent) IsInputMessageContent() bool { return true }

func (ivmc *InputVenueMessageContent) IsInputMessageContent() bool { return true }

func (icmc *InputContactMessageContent) IsInputMessageContent() bool { return true }
