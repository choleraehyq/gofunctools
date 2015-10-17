package functools

import (
	"errors"
	"reflect"
)

func Filter(function, slice interface{}) (ret interface{}, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
	ret = filter(function, slice)
	return
}

func filter(function, slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("filter: The first param is not a slice")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyFilterFuncType(fn, inType) {
		panic("apply: Function must be of type func(" + inType.String() + ") bool")
	}
	var param [1]reflect.Value
	out := reflect.MakeSlice(in.Type(), 0, in.Len())
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		if fn.Call(param[:])[0].Bool() {
			out = reflect.Append(out, in.Index(i))
		}
	}
	return out.Interface()
}

func verifyFilterFuncType(fn reflect.Value, elemType reflect.Type) bool {
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
