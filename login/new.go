package login

// App represents a widget which get and validate users authorizations.
type App struct{ SecretKey string }

// New create new app widget for validate authorizations with bot token as secret key.
func New(accessToken string) *App {
	return &App{SecretKey: accessToken}
}
