// Package functools is a simple Golang library including some commonly used functional programming tools.
// There is no generic in golang, so most of the functions will return interface{}, be sure to type assert it to the type you want before you use it.
package functools

import (
	"errors"
	"reflect"
)

// Apply applys a function(the first parameter) to each element of a slice(second parameter). Just like Map in other language.
func Apply(function, slice interface{}) (ret interface{}, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
	ret = apply(function, slice)
	return
}

func apply(function, slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("apply: The first param is not a slice")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyApplyFuncType(fn, inType) {
		panic("apply: Function must be of type func(" + inType.String() + ") outputElemType")
	}
	var param [1]reflect.Value
	out := in
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		out.Index(i).Set(fn.Call(param[:])[0])
	}
	return out.Interface()
}

func verifyApplyFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	if fn.Type().In(0) != elemType {
		return false
	}
	return true
}
