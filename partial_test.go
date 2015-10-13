package functools

import (
	"testing"
)

func TestPartial(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}
	rawAddTwo, funcType, err := Partial(add, 2)
	if err != nil {
		t.Fatalf("Partial() failed: %v", err)
	}
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
