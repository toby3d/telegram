package telegram

func NewForceReply(selective bool) *ForceReply {
	return &ForceReply{
		ForceReply: true,
		Selective:  selective,
	}
}
