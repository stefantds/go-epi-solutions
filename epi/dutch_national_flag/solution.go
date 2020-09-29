package dutch_national_flag

func DutchFlagPartition(pivotIndex int, a []Color) {
	pivot := a[pivotIndex]
	smallerIdx := 0
	equalIdx := 0
	largerIdx := len(a)

	for equalIdx < largerIdx {
		if a[equalIdx] < pivot {
			a[smallerIdx], a[equalIdx] = a[equalIdx], a[smallerIdx]
			smallerIdx++
			equalIdx++
		} else if a[equalIdx] == pivot {
			equalIdx++
		} else { // a[equalIdx] > pivot
			largerIdx--
			a[largerIdx], a[equalIdx] = a[equalIdx], a[largerIdx]
		}
	}
}
