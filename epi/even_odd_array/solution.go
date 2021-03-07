package even_odd_array

func EvenOdd(a []int) {
	for i, j := 0, len(a)-1; i < j; {
		if a[i]%2 == 0 {
			i++
			continue
		}

		a[i], a[j] = a[j], a[i]
		j-- // we know that a[i] was odd, so j can be decreased
	}
}

// EvenOddVersion2 swaps only when finding two items that are misplaced
func EvenOddVersion2(a []int) {
	for i, j := 0, len(a)-1; i < j; {
		switch {
		case a[i]%2 == 0:
			i++
		case a[j]%2 == 1:
			j--
		default:
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
}
