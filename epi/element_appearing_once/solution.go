package element_appearing_once

func FindElementAppearsOnce(a []int64) int64 {
	// bitCounts[i] will hold the count of bits at position i set to 1
	// in values from a. The value is modulo 3.
	bitCounts := make([]int8, 64)
	for _, v := range a {
		for pos := 0; pos < 64; pos++ {
			if v&1 != 0 {
				bitCounts[pos] += 1
				bitCounts[pos] %= 3
			}
			v >>= 1
		}
	}

	// build an int value from the counts array
	// the values appearing 3 times would have cancelled themselves,
	// so the remaining values are the bits of the searched value.
	result := int64(0)
	for i := 63; i >= 0; i-- {
		result <<= 1
		if bitCounts[i] == 1 {
			result |= 1
		}
	}

	return result
}
