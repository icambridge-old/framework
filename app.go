package framework

import (
	"reflect"
)

type App struct {
	controllers map[string]reflect.Type

	router Router
}

func (a *App) RegisterController(c interface{}) {
	t := reflect.TypeOf(c)
	a.controllers[t.Name()] = t
}

func (a *App) RegisterRouter(r Router) {
	a.router = r
}

func NewApp() *App {

	return &App{controllers: map[string]reflect.Type{}}
}
