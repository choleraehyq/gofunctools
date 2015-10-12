package functools

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	initial := 0
	in := []int{1, 2, 3}
	add := func(a, b int) int {
		return a + b
	}
	expect := 6
	out := Reduce(add, in, initial)
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("Reduce() failed: expect %v got %v", expect, out)
	}
}
