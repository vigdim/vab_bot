package site_router

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func Account(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/account.html")

	//_ = ctx.Response.SendFile("views/pages/account.html")
	//UserName := ctx.UserValue("UserName")
}

func Assets(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/assets/js/script.min.js")
}

func Init() {
	r := router.New()
	//r.ServeFiles("/assets/{filepath:*}", "views/pages/assets/{filepath:*}")
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/account", Account)

	r.GET("/assets/{filepath:*}", Assets)

	// ngrok http --domain=workable-grouse-clean.ngrok-free.app 80
	// https://workable-grouse-clean.ngrok-free.app/
	log.Fatal(fasthttp.ListenAndServe(":80", r.Handler))
}
