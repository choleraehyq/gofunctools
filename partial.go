package functools

import (
	"reflect"
)

func Partial(function interface{}, params ...interface{}) (ret interface{}, funcType reflect.Type, err error) {
	err = nil
	defer func() {
		err = recover()
	}()
	ret, funcType = partial(function, params...)
	return
}

func partial(function interface{}, params ...interface{}) (ret interface{}, funcType reflect.Type) {
	fn := reflect.ValueOf(function)
	if fn.Kind() != reflect.Func {
		panic("partial: The first param is not a function")
	}
	inElem := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		inElem = append(in, reflect.ValueOf(param))
	}
	if !verifyPartialFuncType(fn, inElem) {
		panic("partial: The type of function and params are not matched")
	}
	partialedFunc := func(in []reflect.Value) []reflect.Value {
		params := make([]reflect.Value, 0, len(in)+len(inElem))
		params = append(params, inElem...)
		params = append(params, in...)
		return fn.Call(params[:])
	}
}

func verifyPartialFuncType(fn reflect.Value, in []reflect.Value) bool {
	if fn.Type().NumIn() <= len(in) {
		return false
	}
	for i := 0; i < fn.Type().NumIn(); i++ {
		if fn.Type().In(i) != in[i].Type() {
			return false
		}
	}
	return true
}
