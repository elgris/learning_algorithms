package sort

type linkedNode struct {
	value    int
	next     *linkedNode
	previous *linkedNode
}

func Insertionsort(input []int) {
	for i := 1; i < len(input); i++ {
		item := input[i]

		pos := findPositionToInsertBinary(input[:i], item)

		for j := i; j > pos; j-- {
			input[j] = input[j-1]
		}
		input[pos] = item
	}
}

func findPositionToInsert(haystack []int, needle int) int {
	j := 0
	for ; (j < len(haystack)) && (needle > haystack[j]); j++ {
	}
	return j
}

func findPositionToInsertBinary(haystack []int, needle int) int {
	l := 0
	r := len(haystack) - 1

	for l < r {
		middle := l + (r-l)>>1

		if haystack[middle] > needle {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}

	if haystack[l] < needle {
		l++
	}
	return l
}
