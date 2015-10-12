package functools

import (
    "reflect"
)

func All(function, slice interface{}) (ret bool, err error) {
    err = nil
    defer func() {
        err = recover()
    }()
    ret = all(function, slice)
    return
}

func all(function, slice interface{}) bool {
    in := reflect.ValueOf(slice)
    if in.Kind() != reflect.Slice {
        panic("all: The first param is not a slice")
    }
    fn := reflect.ValueOf(function)
    inType := in.Type().Elem()
    if !verifyAllFuncType(fn, inType) {
        panic("all: Function must be of type func(" + inType.String() + ") bool")
    }
    var param [1]reflect.Value
    out := true
    for i := 0; i < in.Len(); i++ {
        param[0] := in.Index(i)
        if !fn.Call(param[:])[0].Bool() {
            out = false
            break
        }
    }
    return out
}

func verifyAllFuncType(fn reflect.Value, elemType reflect.Type) bool {
    if fn.Kind() != reflect.Func {
        return false
    }
    if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
        return false
    }
    if fn.Type().In(0) != elemType || fn.Type().Out(0) != reflect.Bool {
        return false
    }
    return true
}