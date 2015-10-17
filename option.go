package functools

import (
	"reflect"
)

type Option struct {
	val interface{}
}

var None Option = Option{nil}

func Some(i interface{}) Option {
	return Option{i}
}

func (this *Option) Is_some() bool {
	if this.val != nil {
		return true
	}
	return false
}

func (this *Option) Is_none() bool {
	if this.val == nil {
		return true
	}
	return false
}

func (this *Option) Unwrap() interface{} {
	if this.Is_none() {
		panic("Option: Unwrapped option is none")
	}
	return this.val
}

func (this *Option) Bind(function interface{}) Option {
	fn := reflect.ValueOf(function)
	if !this.verifyBindFuncType(fn) {
		panic("Bind: Function must be of type func (valType) Option")
	}
	var param [1]reflect.Value
	param[0] = reflect.ValueOf(this.Unwrap())
	out := fn.Call(param[:])[0]
	return Option{out.Interface()}
}

func (this *Option) And(other Option) Option {
	if this.Is_none() {
		return None
	}
	return other
}

func (this *Option) And_then(function interface{}) Option {
	if this.Is_none() {
		return None
	}
	return this.Bind(function)
}

func (this *Option) verifyBindFuncType(fn reflect.Value) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	val := reflect.ValueOf(this.val)
	if fn.Type().In(0) != val.Type() {
		return false
	}
	return true
}
