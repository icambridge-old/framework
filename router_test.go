package framework

import (
	"testing"
)

func TestRouter_hasRoute(t *testing.T) {
	// Todo build setUp and tearDown
	router := NewRouter()
	router.RegisterController(TestController{})

	if router.hasRoute("TestController", "Test") == false {
		t.Error("Expected TestController to be registered it wasn't")
	}
}

func TestRouter_RegisterController(t *testing.T) {

	a := NewRouter()
	c := TestController{}
	a.RegisterController(c)

	_, ok := a.controllers["TestController"]

	if ok == false {
		t.Error("Expected TestController to be registered it wasn't")
	}

}

func TestRouter_getControllerAndAction_TwoPartsGiven(t *testing.T) {
	router := NewRouter()
	controller, action, params := router.getControllerAndAction("/home/index")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
		return
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
		return
	}

	if len(params) != 0 {
		t.Errorf("Expected empty params but got something back")
		return
	}
}


func TestRouter_getControllerAndAction_TwoPartsGivenWithParams(t *testing.T) {
	router := NewRouter()
	controller, action, params := router.getControllerAndAction("/home/index/one/two")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
		return
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
		return
	}

	if len(params) != 2 {
		t.Errorf("Expected empty params but got something back")
		return
	}

	if params[0] != "one" {

		t.Errorf("Expected one but got %v", params[0])
		return
	}

	if params[1] != "two" {

		t.Errorf("Expected two but got %v", params[1])
		return
	}
}

func TestRouter_getControllerAndAction_OnePartsGiven(t *testing.T) {
	router := NewRouter()
	controller, action, params := router.getControllerAndAction("/home")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}

	if len(params) != 0 {
		t.Errorf("Expected empty params but got something back")
	}
}

func TestRouter_getControllerAndAction_ZeroPartsGiven(t *testing.T) {

	router := NewRouter()
	controller, action, params := router.getControllerAndAction("/")
	expectedController, expectedAction := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}

	if len(params) != 0 {
		t.Errorf("Expected empty params but got something back")
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
