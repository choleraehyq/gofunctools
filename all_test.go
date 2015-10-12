package functools

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	in := []int{2, 4, 6}
	expect := true
	isEven := func(a int) bool {
		if a%2 == 0 {
			return true
		}
		return false
	}
	out := All(isEven, in)
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("All() failed: expected %v got %v", expect, out)
	}
}
