package framework

import (
	"reflect"
	"testing"
)

func Test_getMethods_OneMethod(t *testing.T) {

	testType := reflect.TypeOf(TestController{})

	methods := getMethods(testType)

	expectedLen, actualLen := 1, len(methods)

	if expectedLen != actualLen {
		t.Errorf("Expected a count of %v but got %v", expectedLen, actualLen)
		return
	}

	expectedMethod := "Test"
	actualMethod, ok := methods[expectedMethod]

	if ok != true {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
	}
}

func Test_getMethods_TwoMethods(t *testing.T) {

	testType := reflect.TypeOf(SecondTest{})

	methods := getMethods(testType)

	expectedLen, actualLen := 2, len(methods)

	if expectedLen != actualLen {
		t.Errorf("Expected a count of %v but got %v", expectedLen, actualLen)
		return
	}

	expectedMethod := "TestOne"
	actualMethod, ok := methods[expectedMethod]

	if ok != true {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
	}

	expectedMethod = "TestTwo"
	actualMethod, ok = methods[expectedMethod]

	if ok != true {
		t.Errorf("Expected a method of %v but got %v", expectedMethod, actualMethod)
		return
	}
}
