package goin

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
		result   []int
	}{
		{[]int{1, 2, 3}, true, []int{1, 3, 2}},
		{[]int{3, 2, 1}, false, []int{3, 2, 1}},
		{[]int{1, 1, 5}, true, []int{1, 5, 1}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := NextPermutation(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}

			if !reflect.DeepEqual(test.input, test.result) {
				t.Errorf("Expected %v, but got %v", test.result, test.input)
			}
		})
	}
}

func TestLastPermutation(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
		result   []int
	}{
		{[]int{1, 2, 3}, false, []int{1, 2, 3}},
		{[]int{3, 2, 1}, true, []int{3, 1, 2}},
		{[]int{5, 1, 1}, true, []int{1, 5, 1}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := LastPermutation(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}

			if !reflect.DeepEqual(test.input, test.result) {
				t.Errorf("Expected %v, but got %v", test.result, test.input)
			}
		})
	}
}
