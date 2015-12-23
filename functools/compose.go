package functools

import (
	"errors"
	"reflect"
	"strconv"
)

// Compose will compose the received functions, there is no limit on the number of functions.
// Notice that the return value of the composed function is interface{}.
func Compose(functions ...interface{}) (ret func(...interface{}) interface{}, err error) {
	defer getErr(&err)
	ret = compose(functions...)
	return
}

func compose(functions ...interface{}) (ret func(...interface{}) interface{}) {

	verifyComposeFuncType(functions)

	composedFunc := func(in ...interface{}) interface{} {
		param := make([]reflect.Value, 0, len(in))
		for _, inParam := range in {
			param = append(param, reflect.ValueOf(inParam))
		}
		for i := 0; i < len(functions); i++ {
			thisFn := reflect.ValueOf(functions[i])
			param = thisFn.Call(param[:])
		}
		return param[0].Interface()
	}
	return composedFunc
}

func verifyComposeFuncType(functions []interface{}) {
	for i, function := range functions {
		fn := reflect.ValueOf(function)
		if fn.Kind() != reflect.Func {
			newErr(errors.New("Param "+strconv.Itoa(i)+" is not a function"), "Compose")
		}
	}
	for i := 0; i < len(functions)-2; i++ {
		thisFn := reflect.ValueOf(functions[i])
		nextFn := reflect.ValueOf(functions[i+1])
		if !canPipe(thisFn, nextFn) {
			newErr(errors.New("Function "+strconv.Itoa(i)+" and "+strconv.Itoa(i+1)+" cannot be piped."), "Compose")
		}
	}
}

func canPipe(thisFn, nextFn reflect.Value) bool {
	if thisFn.Type().NumOut() != nextFn.Type().NumIn() {
		return false
	}
	for i := 0; i < thisFn.Type().NumOut(); i++ {
		if thisFn.Type().Out(i) != nextFn.Type().In(i) {
			return false
		}
	}
	return true
}
