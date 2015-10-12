package functools

import (
	"reflect"
)

func Any(function, slice interface{}) (ret bool, err error) {
	err = nil
	defer func() {
		err = recover()
	}()
	ret = any(function, slice)
	return
}

func any(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("any: The first param is not a slice")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyAnyFuncType(fn, inType) {
		panic("any: Function must be of type func(" + inType.String() + ") bool")
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
	if fn.Type().NumIn() != 1 || fn.Type.NumOut() != 1 {
		return false
	}
	if fn.Type().In(0) != elemType || fn.Type().Out(0) != reflect.Bool {
		return false
	}
	return true
}
