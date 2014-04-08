package framework

import (
	"reflect"
	"testing"
)

func TestApp_RegisterController(t *testing.T) {

	type TestController struct {
		Controller
	}

	a := NewApp()
	a.RegisterController(TestController{})

	if _, ok := a.controllers["TestController"]; ok == false {
		t.Error("Expected TestController to be registered it wasn't")
	}

}

func TestNewApp(t *testing.T) {

	a := NewApp()

	if len(a.controllers) != 0 {
		t.Errorf("Expected no controllers, %v found", len(a.controllers))
	}
}

func TestApp_RegisterRouter(t *testing.T) {

	r := Router{}

	a := NewApp()
	a.RegisterRouter(r)

	if !reflect.DeepEqual(r, a.router) {
		t.Errorf("Router didn't match, expected %v got %v", r, a.router)
	}
}


func TestApp_getControllerAndAction_TwoPartsGiven(t *testing.T) {
	a := NewApp()
	controller, action := a.getControllerAndAction("/home/index")
	expectedController,expectedAction  := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}

func TestApp_getControllerAndAction_OnePartsGiven(t *testing.T) {
	a := NewApp()
	controller, action := a.getControllerAndAction("/home")
	expectedController,expectedAction  := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}

func TestApp_getControllerAndAction_ZeroPartsGiven(t *testing.T) {
	a := NewApp()
	controller, action := a.getControllerAndAction("/")
	expectedController,expectedAction  := "Home", "Index"

	if controller != expectedController {
		t.Errorf("Expected %v got %v", expectedController, controller)
	}

	if action != expectedAction {
		t.Errorf("Expected %v got %v", expectedAction, action)
	}
}
