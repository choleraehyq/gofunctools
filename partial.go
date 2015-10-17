package functools

import (
	"errors"
	"reflect"
)

func Partial(function interface{}, params ...interface{}) (ret func(...interface{}) interface{}, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
	ret, funcType = partial(function, params...)
	return
}

func partial(function interface{}, params ...interface{}) (ret func(...interface{}) interface{}) {
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

	partialedFunc := func(in ...interface{}) interface{} {
		params := make([]reflect.Value, 0, len(in)+len(inElem))
		params = append(params, inElem...)
		for _, inParam := range in {
			params = append(params, reflect.ValueOf(inParam))
		}
		return fn.Call(params[:])[0].Interface()
	}
	return partialedFunc
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
