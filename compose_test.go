package functools

import (
	"testing"
)

func TestCompose(t *testing.T) {
	add := func(a, b int) {
		return a + b
	}
	minusOne := func(a int) {
		return a - 1
	}
	isEven := func(a int) bool {
		if a%2 == 0 {
			return true
		}
		return false
	}
	rawFunc, funcType := Compose(add, minusOne, isEven)
	if _, ok := rawFunc.(funcType); !ok {
		t.Fatalf("Compose() failed: returned wrong funcType")
	}
	composedFunc := rawFunc.(funcType)
	expect := true
	out := composedFunc(1, 2)
	if expect != out {
		t.Fatalf("Compose() failed: expected %v got %v", expect, out)
	}
}
