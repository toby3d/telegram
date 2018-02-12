package telegram

type Bot struct {
	AccessToken string
	Self        *User
}

func NewBot(accessToken string) (*Bot, error) {
	var err error
	bot := &Bot{AccessToken: accessToken}

	bot.Self, err = bot.GetMe()
	return bot, err
}
