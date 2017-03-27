package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testData := []struct {
		haystack      []int
		needle        int
		expected      int
		expectedFound bool
	}{
		{
			haystack:      []int{0, 0, 0, 0, 0, 0, 0},
			needle:        0,
			expected:      3,
			expectedFound: true,
		},
		{
			haystack:      []int{0, 0, 0, 0, 0, 0, 0},
			needle:        1,
			expectedFound: false,
		},
		{
			haystack:      []int{1, 2, 3, 4, 5, 6},
			needle:        1,
			expectedFound: true,
			expected:      0,
		},
		{
			haystack:      []int{1, 2, 3, 4, 5, 6},
			needle:        6,
			expectedFound: true,
			expected:      5,
		},
		{
			haystack:      []int{1, 2, 3, 4, 5, 6},
			needle:        4,
			expectedFound: true,
			expected:      3,
		},
	}

	for i, testItem := range testData {
		pos, ok := BinarySearch(testItem.haystack, testItem.needle)

		if ok != testItem.expectedFound {
			t.Fatalf("#%d: Expected: %v Got: %v", i, testItem.expectedFound, ok)
		} else if pos != testItem.expected {
			t.Fatalf("#%d: Expected position: %v Got: %v", i, testItem.expected, pos)
		}

	}

}
