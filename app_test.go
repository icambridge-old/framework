package framework

import (
	"reflect"
	"testing"
)

func TestApp_RegisterRouter(t *testing.T) {

	r := Router{}

	a := NewApp(1)
	a.RegisterRouter(r)

	if !reflect.DeepEqual(r, a.router) {
		t.Errorf("Router didn't match, expected %v got %v", r, a.router)
	}
}
