package functools

import (
    "reflect"
)

func Compose(functions ...interface{}) (ret interface{}, funcType reflect.Type, err error) {
    err = nil
    defer func() {
        err = recover()
    }()
    ret, funcType = compose(functions... , slice)
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

func verifyComposeFuncType(functions []interface) {
    for _, function := range functions {
        fn := reflect.ValueOf(function)
        if fn.Kind() != reflect.Func {
            panic("compose: Param %i is not a function", i)
        }
    }
    for i := 0; i < len(functions)-2; i++ {
        thisFn := reflect.ValueOf(function[i])
        nextFn := reflect.ValueOf(functions[i+1])
        if !canPipe(thisFn, nextFn) {
            panic("compose: Function %i and %i cannot be piped.", i, i+1)
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