package framework

import (
	"testing"
)

func Test_SliceHasString_DoesHaveString(t *testing.T) {

	findString := "find me"
	s := []string{findString}

	if !SliceHasString(findString, s) {
		t.Errorf("Expectedt to find %v, but didn't.", findString)
	}
}


func Test_SliceHasString_DoesHaveString_WithOthers(t *testing.T) {

	findString := "find me"
	s := []string{"in here", findString, "others"}

	if !SliceHasString(findString, s) {
		t.Errorf("Expectedt to find %v, but didn't.", findString)
	}
}


func Test_SliceHasString_DoesntHaveString(t *testing.T) {

	findString := "find me"
	s := []string{}

	if SliceHasString(findString, s) {
		t.Errorf("Expected to not find %v, but did.", findString)
	}
}
