package functools

import (
	"errors"
	"reflect"
)

// Any applys a function(the first parameter) returning a boolean value to each element of a slice(second parameter), if there exist at least one element make that function return true then Any will return true, otherwise false.
// Notice that Any return a boolean value NOT an interface{}
func Any(function, slice interface{}) (ret bool, err error) {
	defer getErr(&err)
	ret = any(function, slice)
	return
}

func any(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		newErr(errors.New("The first param is not a slice"), "Any")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyAnyFuncType(fn, inType) {
		newErr(errors.New("Function must be of type func("+inType.String()+") bool"), "Any")
	}
	var param [1]reflect.Value
	out := false
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		if fn.Call(param[:])[0].Bool() {
			out = true
			break
		}
	}
	return out
}

func verifyAnyFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	if fn.Type().In(0) != elemType || fn.Type().Out(0).Kind() != reflect.Bool {
		return false
	}
	return true
}
