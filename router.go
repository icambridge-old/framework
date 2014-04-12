package framework

import (
	"fmt"
	"strings"
	"net/http"

	"reflect"
)

func NewRouter() *Router {
	return &Router{controllers: map[string]StructInfo{}}
}

type Router struct {
	// Contains the controllers that are registered for the application
	controllers map[string]StructInfo
}

func (r *Router) Handle(res http.ResponseWriter, req *http.Request) {

	URI := req.URL.Path

	controller, action := r.getControllerAndAction(URI)

	if !r.hasRoute(controller, action) {
		// todo handle 404 nicely
		fmt.Fprintf(res, "No such controller")
		return
	}

	controllerInfo := r.controllers[controller]
	value := controllerInfo.Value.MethodByName(action).Call([]reflect.Value{})

	html := value[0].String()

	fmt.Fprintf(res, html)
}

func (r *Router) RegisterController(c interface{}) {

	structInfo := getStructInfo(c)
	r.controllers[structInfo.Name] = structInfo
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

func (a *Router) hasRoute(controller string, action string) bool {


	controllerInfo, ok := a.controllers[controller]

	if ok == false {
		return false
	}

	_, ok = controllerInfo.Methods[action]

	return ok
}
