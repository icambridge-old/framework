package framework

import (
	"reflect"
	"strings"
	"unicode"
)

type App struct {
	controllers map[string]StructInfo

	router Router
}

type StructInfo struct {
	Name    string
	Actions map[string]MethodInfo
	Type    reflect.Type
}

type MethodInfo struct {
	Name    string
	Type    reflect.Method
}

func (a *App) RegisterController(c interface{}) {
	t := reflect.TypeOf(c)

	ci := StructInfo{}
	ci.Name = t.Name()
	ci.Type = t
	ci.Actions = getMethods(t)

	a.controllers[t.Name()] = ci
}

func getMethods(reflectedType reflect.Type) map[string]MethodInfo {

	count := reflectedType.NumMethod()
	methods := map[string]MethodInfo{}

	for i := 0; i < count; i++ {
		method := reflectedType.Method(i)
		methods[method.Name] = MethodInfo{
			Name: method.Name,
			Type: method,
		}
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

	return &App{controllers: map[string]StructInfo{}}
}

func ucfirst(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}
