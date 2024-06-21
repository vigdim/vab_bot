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
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.Error("AccountPost не вернула результат", 0)
		return
	}
	//sName := string(ctx.FormValue("name"))
	sName := form.Value["name"][0]
	sPhone := form.Value["phone"][0]
	sEmail := form.Value["email"][0]
	sUserId := form.Value["us_id"][0]
	println(sName, sPhone, sEmail, sUserId)

	err = methods.AddUser(sUserId, sEmail, sName, sPhone, 0, "user")
	if err != nil {
		ctx.Error(err.Error(), 0)
	}
}

func Assets(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/assets/js/script.min.js")
}

func Init() {
	r := router.New()

	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/account", AccountGet)
	r.GET("/assets/js/script.min.js", Assets)
	r.POST("/account_post", AccountPost)

	// ngrok http --domain=workable-grouse-clean.ngrok-free.app 80
	// https://workable-grouse-clean.ngrok-free.app/

	log.Fatal(fasthttp.ListenAndServe(":80", r.Handler))
}
