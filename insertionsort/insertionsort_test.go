package insertionsort

import (
	"reflect"
	"testing"
)

func TestSortAlreadySorted(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	sorted := Sort(list)
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected sorted list to not be changed, got: %x", sorted)
	}
}

func TestSortUnsorted(t *testing.T) {
	list := []int{4, 1, 78, 2, 6, 232, 99, 43}
	expected := []int{1, 2, 4, 6, 43, 78, 99, 232}
	sorted := Sort(list)
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected sorted list to not be changed, got: %x", sorted)
	}
}

func TestSortEmpty(t *testing.T) {
	list := []int{}
	expected := []int{}
	sorted := Sort(list)
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected sorted list to not be changed, got: %x", sorted)
	}
}

func TestSortOneValue(t *testing.T) {
	list := []int{632}
	expected := []int{632}
	sorted := Sort(list)
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected sorted list to not be changed, got: %x", sorted)
	}
}
