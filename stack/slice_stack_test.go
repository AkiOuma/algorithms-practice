package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	data := map[string]float64{
		"a": 5,
		"b": 21,
		"c": 3,
		"d": 4,
		"e": 17,
		"f": 2,
	}
	result := RPN("a+(b-c*d+e)/f", data)
	if result != 18 {
		t.Error("error: RPN did not pass test")
	}
}
