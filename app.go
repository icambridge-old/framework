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
	mux.Handle("/js/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/css/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/font/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/", a.router)
	http.ListenAndServe(dsn, mux)
}

func NewApp(port int) *App {
	return &App{port: port}
}
