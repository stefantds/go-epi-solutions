package longest_contained_interval

func LongestContainedRange(a []int) int {
	unprocessed := make(map[int]bool, len(a))
	for _, v := range a {
		unprocessed[v] = true
	}

	maxLength := 0
	for _, v := range a {
		if unprocessed[v] {
			length := 0
			for i := v; unprocessed[i]; i++ {
				delete(unprocessed, i)
				length++
			}
			for i := v - 1; unprocessed[i]; i-- {
				delete(unprocessed, i)
				length++
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}
