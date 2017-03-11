package functools

import (
	"errors"
	"reflect"
)

// All applies a function(the first parameter) returning a boolean value to each element of a slice(second parameter), if all elements make that function return true then All will return true, otherwise false.
// Notice that All returns a boolean value, not an interface{}.
func All(function, slice interface{}) (ret bool, err error) {
	defer getErr(&err)
	ret = all(function, slice)
	return
}

func all(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		newErr(errors.New("The first param is not a slice"), "All")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyAllFuncType(fn, inType) {
		newErr(errors.New("Function must be of type func("+inType.String()+") bool"), "All")
	}
	var param [1]reflect.Value
	out := true
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		if !fn.Call(param[:])[0].Bool() {
			out = false
			break
		}
	}
	return out
}

func verifyAllFuncType(fn reflect.Value, elemType reflect.Type) bool {
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
