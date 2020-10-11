package sorted_array_remove_dups

func DeleteDuplicates(a []int) int {
	if len(a) == 0 {
		return 0
	}

	writeIdx := 1
	for readIdx := 1; readIdx < len(a); readIdx++ {
		if a[writeIdx-1] != a[readIdx] {
			a[writeIdx] = a[readIdx]
			writeIdx += 1
		}
	}

	return writeIdx
}
