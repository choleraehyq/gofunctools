package functools

import (
	"errors"
	"reflect"
)

// Partial will make a partial function. The first argument is the function being partialed, the rest arguments is the parameters sending to that function.
// Notice that the return value of the partial function is interface{}.
// Example:
// add := func(a, b int) int {
// 	return a + b
// }
// addTwo, err := Partial(add, 2)
// out := addTwo(1).(int)
func Partial(function interface{}, params ...interface{}) (ret func(...interface{}) interface{}, err error) {
	err = nil
	defer func() {
		if errString := recover(); errString != nil {
			err = errors.New(errString.(string))
		}
	}()
	ret = partial(function, params...)
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
	for i := 0; i < len(in); i++ {
		if fn.Type().In(i) != in[i].Type() {
			return false
		}
	}
	return true
}
