package tree

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := sort.IntSlice{1, 4, 6, 9, 22, 34, 100}
	index, err := BinarySearch(array, 9)
	if index != 3 || err != nil {
		t.Error("BinarySearch test failed.")
		return
	}
	_, err = BinarySearch(array, 10000)
	if err == nil {
		t.Error("BinarySearch test failed.")
		return
	}
}
