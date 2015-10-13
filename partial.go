package functools

import (
	"errors"
	"reflect"
)

func Partial(function interface{}, params ...interface{}) (ret interface{}, funcType reflect.Type, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
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
		inElem = append(inElem, reflect.ValueOf(param))
	}
	if !verifyPartialFuncType(fn, inElem) {
		panic("partial: The type of function and params are not matched")
	}

	inType := make([]reflect.Type, 0, fn.Type().NumIn()-len(params))
	for i := len(params); i < fn.Type().NumIn(); i++ {
		inType = append(inType, fn.Type().In(i))
	}
	outType := make([]reflect.Type, 0, fn.Type().NumOut())
	for i := 0; i < fn.Type().NumOut(); i++ {
		outType = append(outType, fn.Type().Out(i))
	}
	funcType = reflect.FuncOf(inType, outType, false)

	partialedFunc := func(in []reflect.Value) []reflect.Value {
		params := make([]reflect.Value, 0, len(in)+len(inElem))
		params = append(params, inElem...)
		params = append(params, in...)
		return fn.Call(params[:])
	}
	ret = reflect.MakeFunc(funcType, partialedFunc)
	return
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
