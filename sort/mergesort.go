package sort

func Mergesort(input []int) {
	buffer := make([]int, len(input))
	sort(input, buffer)
}

func sort(input, buffer []int) {
	if len(input) == 1 {
		buffer[0] = input[0]
		return
	}
	half := len(input) / 2
	l := input[:half]
	r := input[half:]
	sort(l, buffer[:half])
	sort(r, buffer[half:])
	merge(l, r, buffer)
	copy(input, buffer)
}

func merge(l, r, buffer []int) {
	i := 0
	ir := 0
	il := 0
	for ir < len(r) && il < len(l) {
		if l[il] < r[ir] {
			buffer[i] = l[il]
			il++
		} else {
			buffer[i] = r[ir]
			ir++
		}
		i++
	}

	for ; ir < len(r); ir++ {
		buffer[i] = r[ir]
		i++
	}

	for ; il < len(l); il++ {
		buffer[i] = l[il]
		i++
	}
}
