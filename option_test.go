package functools

import (
	"reflect"
	"testing"
)

func TestIs_none(t *testing.T) {
	opt := Some(1)
	expect := false
	if expect != opt.Is_none() {
		t.Fatalf("Option.Is_none(): expect %v got %v", expect, opt.Is_none())
	}
}

func TestIs_some(t *testing.T) {
	opt := None
	expect := false
	if expect != opt.Is_some() {
		t.Fatalf("Option.Is_some(): expect %v got %v", expect, opt.Is_none())
	}
}

func TestUnwrap(t *testing.T) {
	opt := Some(1)
	expect := 1
	ret := opt.Unwrap()
	if !reflect.DeepEqual(expect, ret) {
		t.Fatalf("Option.Unwrap() failed: expected %v got %v", expect, ret)
	}
}

func TestBind(t *testing.T) {
	opt := Some(1)
	expect := Some(2)
	double := func(in int) int {
		return in * 2
	}
	ret := opt.Bind(double)
	if !reflect.DeepEqual(expect, ret) {
		t.Fatalf("Option.Bind() failed: expected %v got %v", expect, ret)
	}
}

func TestAnd(t *testing.T) {
	opt := Some(nil)
	expect := None
	opt2 := Some(1)
	ret := opt.And(opt2)
	if !reflect.DeepEqual(ret, expect) {
		t.Fatalf("Option.And() failed: expected %v got %v", expect, ret)
	}
}

func TestAnd_then(t *testing.T) {
	opt1 := Some(1)
	double := func(in int) int {
		return in * 2
	}
	ret1 := opt1.And_then(double)
	expect1 := Some(2)
	if !reflect.DeepEqual(ret1, expect1) {
		t.Fatalf("Option.And_then() failed: expected %v got %v", expect1, ret1)
	}
	opt2 := None
	ret2 := opt2.And_then(double)
	expect2 := None
	if !reflect.DeepEqual(ret2, expect2) {
		t.Fatalf("Option.And_then() failed: expected %v got %v", expect2, ret2)
	}
}
