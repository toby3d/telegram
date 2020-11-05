package login_test

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	http "github.com/valyala/fasthttp"
	"gitlab.com/toby3d/telegram/v5/login"
)

const htmlTemplate string = `<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Telegram login</title>
  </head>
  <body>
    <script async src="https://telegram.org/js/telegram-widget.js?11" data-telegram-login="toby3dBot"
      data-size="large" data-auth-url="https://example.site/callback" data-request-access="write"></script>
  </body>
</html>`

func Example_fastStart() {
	// Use bot AccessToken from @BotFather as ClientSecret.
	c := login.Config{
		ClientSecret:       "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
		RedirectURL:        "https://example.site/callback",
		RequestWriteAccess: true,
	}

	// Create example server with authorization and token (callback) handlers.
	r := router.New()
	r.GET("/", func(ctx *http.RequestCtx) {
		// Render page with embeded Telegram Login button (until Telegram enable the possibility of login by
		// link.)
		ctx.SuccessString("text/html", htmlTemplate)

		// NOTE(toby3d): Telegram does not yet allow you to login without script via a link, as is common
		// in traditional OAuth2 applications, stopping at the last step with redirect to callback. The
		// 'embed=[0|1]' parameter has no effect now, which is very similar to a bug.
		//ctx.SuccessString("text/html", fmt.Sprintf(htmlTemplate, c.AuthCodeURL(language.English)))
	})
	r.GET("/callback", func(ctx *http.RequestCtx) {
		q := ctx.QueryArgs()
		u := login.User{
			AuthDate:  int64(q.GetUintOrZero(login.KeyAuthDate)),
			FirstName: string(q.Peek(login.KeyFirstName)),
			Hash:      string(q.Peek(login.KeyHash)),
			ID:        q.GetUintOrZero(login.KeyID),
			LastName:  string(q.Peek(login.KeyLastName)),
			PhotoURL:  string(q.Peek(login.KeyPhotoURL)),
			Username:  string(q.Peek(login.KeyUsername)),
		}

		if !c.Verify(&u) {
			ctx.Error("Unable to verify data", http.StatusUnauthorized)
			return
		}

		ctx.SuccessString("text/plain", fmt.Sprintf("Hello, %s!", u.FullName()))
	})

	if err := http.ListenAndServe(":80", r.Handler); err != nil {
		log.Fatalln(err.Error())
	}
}
