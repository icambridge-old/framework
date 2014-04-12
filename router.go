package framework

import (
	"reflect"
	"strings"
)

func NewRouter() *Router {
	return &Router{controllers: map[string]StructInfo{}}
}

type Router struct {
	// Contains the controllers that are registered for the application
	controllers map[string]StructInfo
}

func (r *Router) RegisterController(c interface{}) {
	reflected := reflect.TypeOf(c)
	structInfo := getStructInfo(reflected)
	r.controllers[reflected.Name()] = structInfo
}

func (a *Router) getControllerAndAction(URI string) (string, string) {

	defaultController := "Home"
	defaultAction := "Index"

	URI = strings.Trim(URI, "/")
	parts := strings.Split(URI, "/")

	partsLen := len(parts)

	if partsLen >= 2 {
		return UpperFirst(parts[0]), UpperFirst(parts[1])
	} else if partsLen == 1 && parts[0] != "" {
		return UpperFirst(parts[0]), defaultAction
	}

	return defaultController, defaultAction
}

func (a *Router) hasController(controller string) bool {

	_, ok := a.controllers[controller]

	return ok
}
