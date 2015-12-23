package functools

import (
	"errors"
	"reflect"
)

// Partial will make a partial function. The first argument is the function being partialed, the rest arguments is the parameters sending to that function.
func Partial(function interface{}, params ...interface{}) (ret func(...interface{}) interface{}, err error) {
	defer getErr(&err)
	ret = partial(function, params...)
	return
}

func partial(function interface{}, params ...interface{}) (ret func(...interface{}) interface{}) {
	fn := reflect.ValueOf(function)
	if fn.Kind() != reflect.Func {
		newErr(errors.New("The first param is not a function"), "Partial")
	}
	inElem := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		inElem = append(inElem, reflect.ValueOf(param))
	}
	if !verifyPartialFuncType(fn, inElem) {
		newErr(errors.New("The type of function and params are not matched"), "Partial")
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
	for i := 0; i < len(in); i++ {
		if fn.Type().In(i) != in[i].Type() {
			return false
		}
	}
	return true
}
