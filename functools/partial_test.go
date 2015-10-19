package functools

import (
	"testing"
)

func TestPartial(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}
	addTwo, err := Partial(add, 2)
	if err != nil {
		t.Fatalf("Partial() failed: %v", err)
	}
	expect := 3
	out := addTwo(1).(int)
	if expect != out {
		t.Fatalf("Partial() failed: expected %i got %i", expect, out)
	}
}
