package functools

import (
	"errors"
	"reflect"
)

// All applys a function(the first parameter) returning a boolean value to each element of a slice(second parameter), if all elements make that function return true then All will return true, otherwise false.
// Notice that All return a boolean value NOT an interface{}
func All(function, slice interface{}) (ret bool, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
	ret = all(function, slice)
	return
}

func all(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("all: The first param is not a slice")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyAllFuncType(fn, inType) {
		panic("all: Function must be of type func(" + inType.String() + ") bool")
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
