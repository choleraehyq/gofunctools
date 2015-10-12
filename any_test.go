package functools

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	in := []int{1, 3, 4}
	expect := true
	isEven := func(a int) bool {
		if a%2 == 0 {
			return true
		}
		return false
	}
	out := Any(isEven, in)
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("Any() failed: expected %v got %v", expect, out)
	}
}
