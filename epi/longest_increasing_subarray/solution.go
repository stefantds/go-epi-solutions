package longest_increasing_subarray

func FindLongestIncreasingSubarray(a []int) (start, end int) {
	maxLength := 1
	currentStart := 0
	currentMax := 1
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			currentMax++
			if currentMax > maxLength {
				start = currentStart
				end = i
				maxLength = currentMax
			}
		} else {
			currentMax = 1
			currentStart = i
		}
	}

	return start, end
}
