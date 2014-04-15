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

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// TODO test!
	URI := req.URL.Path

	controller, action, params := r.getControllerAndActionAndParams(URI)

	if !r.hasRoute(controller, action) {
		// todo handle 404 nicely
		r.fourOhFour(res)
		return
	}

	paramsValues := []reflect.Value{}

	for _, param := range params {
		paramsValues = append(paramsValues, reflect.ValueOf(param))
	}

	controllerInfo := r.controllers[controller]
	var methodType reflect.Method
	methodValue := controllerInfo.Value.MethodByName(action)
	methodType, _  = controllerInfo.Type.MethodByName(action)

	numIn := methodType.Type.NumIn() - 1

	if numIn != len(paramsValues) {
		r.fourOhFour(res)
		return
	}

	value := methodValue.Call(paramsValues)

	html := value[0].String()
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)
}

func (r *Router) fourOhFour(res http.ResponseWriter) {
	// TODO improve.
	fmt.Fprintf(res, "No such controller")
}

func (r *Router) RegisterController(c interface{}) {

	structInfo := getStructInfo(c)
	r.controllers[structInfo.Name] = structInfo
}

func (a *Router) getControllerAndActionAndParams(URI string) (string, string, []string) {

	defaultController := "Home"
	defaultAction := "Index"
	params := []string{}

	URI = strings.Trim(URI, "/")
	parts := strings.Split(URI, "/")

	partsLen := len(parts)

	if partsLen >= 2 {
		params = parts[2:]
		return UpperFirst(parts[0]), UpperFirst(parts[1]), params
	} else if partsLen == 1 && parts[0] != "" {
		return UpperFirst(parts[0]), defaultAction, params
	}

	return defaultController, defaultAction, params
}

func (a *Router) hasRoute(controller string, action string) bool {


	controllerInfo, ok := a.controllers[controller]

	if ok == false {
		return false
	}

	_, ok = controllerInfo.Methods[action]

	return ok
}
