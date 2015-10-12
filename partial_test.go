package functools

import (
	"testing"
)

func TestPartial(t *testing.T) {
	add := func(a, b) {
		return a + b
	}
	rawAddTwo, funcType := Partial(add, 2)
	if _, ok := rawAddTwo.(funcType); !ok {
		t.Fatalf("Partial() failed: returned wrong funcType")
	}
	addTwo := rawAddTwo.(funcType)
	expect := 3
	out := addTwo(1)
	if expect != out {
		t.Fatalf("Partial() failed: expected %i got %i", expect, out)
	}
}
