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

func GetAccount(ctx *fasthttp.RequestCtx) {
	//fasthttp.
}

func Init() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/getaccount", GetAccount)

	// ngrok http --domain=workable-grouse-clean.ngrok-free.app 80
	// https://workable-grouse-clean.ngrok-free.app/
	log.Fatal(fasthttp.ListenAndServe(":80", r.Handler))
}

//func sayhello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Привет!")
//}

//func NewRouter() http.Handler {
//	router := chi.NewRouter()
//	r.Get("/{name}", HelloName)
//
//	// Настройка раздачи статических файлов
//	staticPath, _ := filepath.Abs("views/pages/")
//	fs := http.FileServer(http.Dir(staticPath))
//	router.Handle("/*", fs)
//
//	return r
//}
