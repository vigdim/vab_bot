package http

import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

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
