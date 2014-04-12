package framework

import (
	"reflect"
)

type StructInfo struct {
	Name    string
	Methods map[string]MethodInfo
	Type    reflect.Type
	Value   reflect.Value
}

type MethodInfo struct {
	Name string
	Type reflect.Method
}

func getStructInfo(c interface {}) StructInfo {

	reflected := reflect.TypeOf(c)
	structInfo := StructInfo{}
	structInfo.Name = reflected.Name()
	structInfo.Value = reflect.ValueOf(c)
	structInfo.Type = reflected
	structInfo.Methods = getMethods(reflected)

	return structInfo
}

func getMethods(reflectedType reflect.Type) map[string]MethodInfo {

	count := reflectedType.NumMethod()
	methods := map[string]MethodInfo{}

	for i := 0; i < count; i++ {
		method := reflectedType.Method(i)
		methods[method.Name] = MethodInfo{
			Name: method.Name,
			Type: method,
		}
	}

	return methods
}
