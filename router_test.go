package framework

import (
	"reflect"
	"testing"
)

func TestRouter_hasController(t *testing.T) {
	// Todo build setUp and tearDown
	router := NewRouter()
	router.RegisterController(TestController{})

	if router.hasController("TestController") == false {
		t.Error("Expected TestController to be registered it wasn't")
	}
}

func TestRouter_RegisterController(t *testing.T) {

	a := NewRouter()
	c := TestController{}
	a.RegisterController(c)

	actual, ok := a.controllers["TestController"]

	if ok == false {
		t.Error("Expected TestController to be registered it wasn't")
	}

	rt := reflect.TypeOf(c)
	method := rt.Method(0)
	actions := map[string]MethodInfo{
		method.Name: MethodInfo{
			Name: method.Name,
			Type: method,
		},
	}

	expected := StructInfo{
		Name:    "TestController",
		Type:    rt,
		Methods: actions,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestRouter_getControllerAndAction_TwoPartsGiven(t *testing.T) {
	router := NewRouter()
	controller, action := router.getControllerAndAction("/home/index")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}

func TestRouter_getControllerAndAction_OnePartsGiven(t *testing.T) {
	router := NewRouter()
	controller, action := router.getControllerAndAction("/home")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}

func TestRouter_getControllerAndAction_ZeroPartsGiven(t *testing.T) {

	router := NewRouter()
	controller, action := router.getControllerAndAction("/")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}

type TestController struct {
	Controller
}

func (c TestController) Test() bool {
	return true
}

type SecondTest struct {
	Controller
}

func (c SecondTest) TestOne() bool {
	return true
}

func (c SecondTest) TestTwo() bool {
	return true
}
