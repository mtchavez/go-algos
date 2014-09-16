package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	x := []int{2, 3232323, 2, 334, 6562, 778, 4, 5, 7, 1, 2, 3, 6, 22, 5, 6, 87, 23, 1111}
	MergeSort(x, 0, len(x)-1)
	expected := []int{1, 2, 2, 2, 3, 4, 5, 5, 6, 6, 7, 22, 23, 87, 334, 778, 1111, 6562, 3232323}
	if !reflect.DeepEqual(expected, x) {
		t.Errorf("Expected to be sorted but got %x", x)
	}
}
