package framework

import (
	"reflect"
	"testing"
	"strings"
)

func TestRouter_ParseXml(t *testing.T) {

	xml := `<routes>
		<route>
			<id>home_index</id>
			<controller>Home</controller>
			<action>Index</action>
			<path>/</path>
		</route>
	</routes>`

	r := Router{}
	err := r.ParseXml(strings.NewReader(xml))

	if err != nil {
		t.Errorf("Didn't expect an error but got %v", err)
	}

	routesLen := len(r.Routes)

	if  routesLen != 1 {
		t.Errorf("Expected len 1 instead got %v", routesLen)
	}

	path := "/"
	route, ok := r.Routes[path]
	if ok != true {
		t.Errorf("Expected true for key '%v' existing got %v", path, ok)
	}

	expectedRoute := Route{Id: "home_index", Controller: "Home", Action: "Index", Path: "/"}


	if !reflect.DeepEqual(expectedRoute, route) {
		t.Errorf("Expected %v, got %v", expectedRoute, route)
	}
}
