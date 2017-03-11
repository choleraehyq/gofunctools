package functools

import (
	"reflect"
)

// Type Option represents an optional value: every Option is either Some and contains a value, or None, and contains a nil.
// They have a number of uses. For example, if a function call failes, it can return None on error.
type Option struct {
	val interface{}
}

// None is a special Option containing nil.
var None Option = Option{nil}

// Some will return an Option containing the given value.
// Some cannot use to generate a None. If the given value is nil, it will panic.
func Some(i interface{}) Option {
	if i == nil {
		panic("Some() cannot get nil as argument. Please use None.")
	}
	return Option{i}
}

// Is_some determines whether this Option is Some. (Contains a value)
func (this *Option) Is_some() bool {
	if this.val != nil {
		return true
	}
	return false
}

// Is_none determines whether this Option is None.
func (this *Option) Is_none() bool {
	if this.val == nil {
		return true
	}
	return false
}

// Unwrap returns the wrapped value of an Option.
// If the Option is None, it will panic.
func (this *Option) Unwrap() interface{} {
	if this.Is_none() {
		panic("Option: Unwrapped option is none")
	}
	return this.val
}

// Bind will apply a given function to the wrapped value of this Option and return a new Option containing the return value of the paramter function.
// If this Option is None, then it will return None.
func (this *Option) Bind(function interface{}) Option {
	fn := reflect.ValueOf(function)
	if !this.verifyBindFuncType(fn) {
		panic("Bind: Function must be of type func (valType) Option")
	}
	if this.Is_none() {
		return None
	}
	var param [1]reflect.Value
	param[0] = reflect.ValueOf(this.Unwrap())
	out := fn.Call(param[:])[0]
	return Option{out.Interface()}
}

// And returns None if this option is None, otherwise returns the option received.
func (this *Option) And(other Option) Option {
	if this.Is_none() {
		return None
	}
	return other
}

// And_then returns None if this option is None, otherwise calls the received function with the wrapped value and returns the result Option.
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
