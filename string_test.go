package framework

import (
	"testing"
)

func Test_UpperFirst_AllLowerCase(t *testing.T) {

	in, out := "test", "Test"

	actual := UpperFirst(in)

	if actual != out {
		t.Errorf("Expected %v but got %v", out, actual)
	}
}

func Test_UpperFirst_AllUpperCase(t *testing.T) {

	in, out := "TEST", "TEST"

	actual := UpperFirst(in)

	if actual != out {
		t.Errorf("Expected %v but got %v", out, actual)
	}
}

func Test_UpperFirst_CamelCase(t *testing.T) {

	in, out := "testWord", "TestWord"

	actual := UpperFirst(in)

	if actual != out {
		t.Errorf("Expected %v but got %v", out, actual)
	}
}
