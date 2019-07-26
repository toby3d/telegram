package login_test

/*
import (
	"log"

	httprouter "github.com/buaazp/fasthttprouter"
	http "github.com/valyala/fasthttp"
	"gitlab.com/toby3d/telegram/login"
)

func Example_fastStart() {
	// We use bot AccessToken from @BotFather or telegram.Bot structure
	botAccessToken := "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"

	// Create example server with example callback handler
	r := httprouter.New()
	r.GET("/callback", func(ctx *http.RequestCtx) {
		defer ctx.SetConnectionClose()

		// You not need decode data to User structure if you want only
		// validate it
		u, err := login.ParseUser(ctx.QueryArgs())
		if err != nil {
			ctx.Error("bad request", http.StatusBadRequest)
			return
		}

		// Check User structure
		ok, err := login.CheckAuthorization(u, botAccessToken)
		if err != nil || !ok { // NANI!? It's a invalid data!
			ctx.Error("bad request", http.StatusBadRequest)
			return
		}

		// All is ok! Hello, human!
		ctx.Success("text/html", []byte("hello, "+u.FullName()+"!"))
	})

	if err := http.ListenAndServe(":8000", r.Handler); err != nil {
		log.Fatalln(err.Error())
	}
}
*/
