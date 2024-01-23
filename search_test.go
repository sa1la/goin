package goin

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	expected := 4
	result := BinarySearch(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	target = 11
	expected = -1
	result = BinarySearch(nums, target)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
