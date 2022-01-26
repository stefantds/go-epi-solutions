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

func LongestSubarrayWithDistinctEntriesUsingSet(a []int) int {
	// currentValues will hold the values in the current window
	currentValues := make(map[int]bool)
	maxSubarrayLength := 0
	windowStart := 0

	for _, v := range a {
		_, exists := currentValues[v]
		if exists {
			for a[windowStart] != v {
				delete(currentValues, a[windowStart])
				windowStart++
			}
			// a[windowStart] == v; we still need to advance the window start by 1
			windowStart++
		}
		currentValues[v] = true
		if len(currentValues) > maxSubarrayLength {
			maxSubarrayLength = len(currentValues)
		}
	}

	return maxSubarrayLength
}
