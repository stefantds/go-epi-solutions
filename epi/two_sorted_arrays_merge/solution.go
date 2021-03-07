package two_sorted_arrays_merge

func MergeTwoSortedArrays(a []int, m int, b []int, n int) {
	writeIdx := m + n - 1
	posA, posB := m-1, n-1
	for posA >= 0 && posB >= 0 {
		var nextElem int
		if a[posA] > b[posB] {
			nextElem = a[posA]
			posA--
		} else {
			nextElem = b[posB]
			posB--
		}
		a[writeIdx] = nextElem
		writeIdx--
	}

	for posB >= 0 {
		a[writeIdx] = b[posB]
		writeIdx, posB = writeIdx-1, posB-1
	}
}
