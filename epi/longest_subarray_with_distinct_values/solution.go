package longest_subarray_with_distinct_values

func LongestSubarrayWithDistinctEntries(a []int) int {
	// lastOccurrence holds for a value its position in
	// the current subarray.
	// The current subarray only contains each value once (but only
	// a subset of the values)
	lastOccurrence := make(map[int]int)

	maxSubarrayLength := 0
	start := 0
	end := 0

	for i, v := range a {
		pos, exists := lastOccurrence[v]
		end = i
		if exists && pos >= start {
			// the value is already in the subarray
			start = pos + 1
		}
		if end-start+1 > maxSubarrayLength {
			maxSubarrayLength = end - start + 1
		}
		lastOccurrence[v] = i
	}

	return maxSubarrayLength
}
