package functools

import (
	"errors"
	"reflect"
	"strconv"
)

func Compose(functions ...interface{}) (ret func(...interface{}) interface{}, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
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
			panic("compose: Param " + strconv.Itoa(i) + " is not a function")
		}
	}
	for i := 0; i < len(functions)-2; i++ {
		thisFn := reflect.ValueOf(functions[i])
		nextFn := reflect.ValueOf(functions[i+1])
		if !canPipe(thisFn, nextFn) {
			panic("compose: Function " + strconv.Itoa(i) + " and " + strconv.Itoa(i+1) + " cannot be piped.")
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
