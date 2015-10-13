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
	opt := Some(nil)
	expect := false
	if expect != opt.Is_some() {
		t.Fatalf("Option.Is_some(): expect %v got %v", expect, opt.Is_none())
	}
}

func TestBind(t *testing.T) {

}

func TestAnd(t *testing.T) {

}

func TestAnd_then(t *testing.T) {

}

func TestUnwrap(t *testing.T) {

}