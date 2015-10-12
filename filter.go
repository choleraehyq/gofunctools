package functools

import (
    "reflect"
)

func Filter(function, slice interface{}) (ret interface{}, err error) {
    err = nil
    defer func() {
        err = recover()
    }()
    ret = filter(function, slice)
    return
}

func filter(function, slice interface{}) interface{} {
    in := reflect.ValueOf(slice)
    if in.Kind() != reflect.Slice {
        panic("filter: The first param is not a slice")
    }
    fn := reflect.ValueOf(function)
    inType := in.Type().Elem()
    if !verifyFilterFuncType(fn, inType) {
        panic("apply: Function must be of type func(" + inType.String() + ") bool")
    }
    var param [1]reflect.Value
    out := reflect.MakeSlice(inType, 0, in.Len())
    for i := 0; i < in.Len(); i++ {
        param[0] := in.Index(i)
        if fn.Call(param[:])[0].Bool() {
            out = reflect.Append(out, in.Index(i))
        }
    }
    return out.interface()
}

func verifyFilterFuncType(fn reflect.Value, elemType reflect.Type) bool {
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