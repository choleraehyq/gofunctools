package functools

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4}
	expect := []int{2, 4}
	isEven := func(a int) bool {
		if a%2 == 0 {
			return true
		}
		return false
	}
	out := Filter(isEven, in)
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("Filter() failed: expected %v got %v", expect, out)
	}
}
