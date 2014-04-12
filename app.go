package framework

import (
	"fmt"
	"net/http"
)

type App struct {
	// The port the application is to run on
	port int
	// Maybe remove? Or move code from here to it?
	router Router
	//
	mux *http.ServeMux
}

func (a *App) RegisterRouter(r Router) {
	a.router = r
}

// TODO figure out how to test this...
func (a *App) Start() {

	dsn := fmt.Sprintf(":%d", a.port)

	http.ListenAndServe(dsn, a.mux)
}

func NewApp(port int) *App {
	return &App{port: port}
}
