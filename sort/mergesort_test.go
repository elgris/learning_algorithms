package sort

import (
	"reflect"
	"testing"
)

func TestMergesort(t *testing.T) {
	testData := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			input:    []int{0, 0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0, 0},
		},
		{
			input:    []int{5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			input:    []int{0, 1, 2, 0, 4, 1},
			expected: []int{0, 0, 1, 1, 2, 4},
		},
	}

	for i, testItem := range testData {
		Mergesort(testItem.input)

		if !reflect.DeepEqual(testItem.expected, testItem.input) {
			t.Fatalf("#%d: Expected: %v\nGot: %v", i, testItem.expected, testItem.input)
		}
	}
}
