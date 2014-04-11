package framework

import (
	"reflect"
	"testing"
)

func TestApp_RegisterController(t *testing.T) {

	a := NewApp()
	c := TestController{}
	a.RegisterController(c)

	actual, ok := a.controllers["TestController"]

	if ok == false {
		t.Error("Expected TestController to be registered it wasn't")
	}

	expected := ControllerInfo{
		Name:    "TestController",
		Type:    reflect.TypeOf(c),
		Actions: []string{"Test"},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}



func TestApp_hasController(t *testing.T) {

	a := NewApp()
	a.RegisterController(TestController{})

	if a.hasController("TestController") == false {
		t.Error("Expected TestController to be registered it wasn't")
	}
}

func Test_getMethodNames_OneMethod(t *testing.T) {

	testType := reflect.TypeOf(TestController{})

	methods := getMethodNames(testType)

	expectedLen, actualLen := 1, len(methods)

	if expectedLen != actualLen {
		t.Errorf("Expected a count of %v but got %v", expectedLen, actualLen)
		return
	}

	expectedMethod, actualMethod := "Test", methods[0]

	if expectedMethod != actualMethod {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
	}
}



func Test_getMethodNames_TwoMethods(t *testing.T) {

	testType := reflect.TypeOf(SecondTest{})

	methods := getMethodNames(testType)

	expectedLen, actualLen := 2, len(methods)

	if expectedLen != actualLen {
		t.Errorf("Expected a count of %v but got %v", expectedLen, actualLen)
		return
	}

	expectedMethod, actualMethod := "TestOne", methods[0]

	if expectedMethod != actualMethod {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
	}
	expectedMethod, actualMethod = "TestTwo", methods[1]

	if expectedMethod != actualMethod {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
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
	expectedController, expectedAction := "Home", "Index"

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
	expectedController, expectedAction := "Home", "Index"

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
