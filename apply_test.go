package functools

import (
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {
	in := []int{1, 2, 3}
	expect := []int{2, 4, 6}
	double := func(in int) int {
		return in * 2
	}
	out := Apply(double, a)
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("Apply() failed: expected %v got %v", expect, out)
	}
}
