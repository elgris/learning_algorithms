package search

func BinarySearch(haystack []int, needle int) (int, bool) {
	l := 0
	r := len(haystack) - 1

	for l <= r {
		middle := l + (r-l)>>1

		if haystack[middle] == needle {
			return middle, true
		}

		if haystack[middle] > needle {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}

	return 0, false

}
