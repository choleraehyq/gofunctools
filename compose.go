package functools

import (
	"errors"
	"reflect"
	"strconv"
)

func Compose(functions ...interface{}) (ret interface{}, funcType reflect.Type, err error) {
	err = nil
	defer func() {
		if interfaceErr := recover(); interfaceErr != nil {
			err = errors.New(interfaceErr.(string))
		}
	}()
	ret, funcType = compose(functions...)
	return
}

func compose(functions ...interface{}) (ret interface{}, funcType reflect.Type) {
	firstFn := reflect.ValueOf(functions[0])
	inType := make([]reflect.Type, 0, firstFn.Type().NumIn())
	for i := 0; i < firstFn.Type().NumIn(); i++ {
		inType = append(inType, firstFn.Type().In(i))
	}
	lastFn := reflect.ValueOf(functions[len(functions)-1])
	outType := make([]reflect.Type, 0, lastFn.Type().NumOut())
	for i := 0; i < lastFn.Type().NumOut(); i++ {
		outType = append(outType, lastFn.Type().Out(i))
	}
	funcType = reflect.FuncOf(inType, outType, false)

	verifyComposeFuncType(functions)

	composedFunc := func(in []reflect.Value) []reflect.Value {
		param := in
		for i := 0; i < len(functions); i++ {
			thisFn := reflect.ValueOf(functions[i])
			param = thisFn.Call(param[:])
		}
		return param
	}

	realFunc := reflect.MakeFunc(funcType, composedFunc)
	ret = realFunc.Interface()
	return
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
