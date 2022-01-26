package search_unknown_length_array

func BinarySearchUnknownLength(a ArrayUnknownLength, k int) int {
	pos := 1

	for val, ok := a.Get(pos); ok && val <= k; val, ok = a.Get(pos) {
		if val == k {
			return pos
		}
		pos *= 2
	}

	// now pos is either in the array but a[pos] > k or pos > len(a)
	start, end := 0, pos
	for start <= end {
		mid := start + (end-start)/2
		val, ok := a.Get(mid)
		if ok && val == k {
			return mid
		}
		if !ok || val > k {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
