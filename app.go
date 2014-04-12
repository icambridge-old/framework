package framework

import (
	"fmt"
	"net/http"
)

type App struct {
	// The port the application is to run on
	port int
	// Maybe remove? Or move code from here to it?
	router *Router
}

func (a *App) RegisterRouter(r *Router) {
	a.router = r
}

// TODO figure out how to test this...
func (a *App) Start() {

	dsn := fmt.Sprintf(":%d", a.port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) { a.router.Handle(res, req) })
	http.ListenAndServe(dsn, mux)
}

func NewApp(port int) *App {
	return &App{port: port}
}
