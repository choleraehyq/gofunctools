package functools

import (
	"reflect"
)

func Apply(function, slice interface{}) (ret interface{}, err error) {
	err = nil
	defer func() {
		err = recover()
	}()
	ret = apply(function, slice)
	return
}

func apply(function, slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("apply: The first param is not a slice")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyApplyFuncType(fn, inType) {
		panic("apply: Function must be of type func(" + inType.String() + ") outputElemType")
	}
    var param [1]reflect.Value
    out := in
    for i := 0; i < in.Len(); i++ {
        param[0] = in.Index(i)
        out.Index(i).Set(fn.Call(param[:])[0])
    }
    return out.interface()
}

func verifyApplyFuncType(fn reflect.Value, elemType reflect.Type) bool {
    if fn.Kind() != reflect.Func {
        return false
    }
    if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
        return false
    }
    if fn.Type().In(0) != elemType {
        return false
    }
    return true
}