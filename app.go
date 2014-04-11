package framework

import (
	"reflect"
	"strings"
	"unicode"
)

type App struct {
	controllers map[string]ControllerInfo

	router Router
}

type ControllerInfo struct {
	Name    string
	Actions []string
	Type    reflect.Type
}

func (a *App) RegisterController(c interface{}) {
	t := reflect.TypeOf(c)

	ci := ControllerInfo{}
	ci.Name = t.Name()
	ci.Type = t
	ci.Actions = getMethodNames(t)
	// Move to seperate function

	a.controllers[t.Name()] = ci
}

func getMethodNames(reflectedType reflect.Type) []string {

	count := reflectedType.NumMethod()
	methods := []string{}

	for i := 0; i < count; i++ {
		method := reflectedType.Method(i)
		methods = append(methods, method.Name)
	}

	return methods
}

func (a *App) getControllerAndAction(URI string) (string, string) {

	defaultController := "Home"
	defaultAction := "Index"

	URI = strings.Trim(URI, "/")
	parts := strings.Split(URI, "/")

	partsLen := len(parts)

	if partsLen >= 2 {
		return ucfirst(parts[0]), ucfirst(parts[1])
	} else if partsLen == 1 && parts[0] != "" {
		return ucfirst(parts[0]), defaultAction
	}

	return defaultController, defaultAction
}

func (a *App) hasController(controller string) bool {

	_, ok := a.controllers[controller]

	return ok
}

func (a *App) RegisterRouter(r Router) {
	a.router = r
}

func NewApp() *App {

	return &App{controllers: map[string]ControllerInfo{}}
}

func ucfirst(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}
