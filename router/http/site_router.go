package site_router

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/vigdim/vab_bot/database/methods"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func AccountGet(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/account.html")
}

func AccountPost(ctx *fasthttp.RequestCtx) {
	sName := string(ctx.FormValue("name"))
	sPhone := string(ctx.FormValue("phone"))
	sEmail := string(ctx.FormValue("email"))
	sUserId := string(ctx.FormValue("us_id"))

	// Создаем или обновляем пользователя. При обновлении роль не изменяется
	err := methods.AddUser(sUserId, sEmail, sName, sPhone, 0, "user")
	if err != nil {
		ctx.Error(err.Error(), 0)
	}

	//main_handlers.SendAnyMessage(_, _, "0000")
}

func Assets(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/assets/js/script.min.js")
}

func Init() {
	r := router.New()

	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/account", AccountGet)
	r.POST("/account_post", AccountPost)
	r.GET("/assets/js/script.min.js", Assets)

	// ngrok http --domain=workable-grouse-clean.ngrok-free.app 80
	// https://workable-grouse-clean.ngrok-free.app/
	log.Fatal(fasthttp.ListenAndServe(":80", r.Handler))
}
