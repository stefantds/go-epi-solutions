package longest_nondecreasing_subsequence

func LongestNondecreasingSubsequenceLength(a []int) int {
	longestEndingAtPos := make([]int, len(a))
	for i := range longestEndingAtPos {
		longestEndingAtPos[i] = 1
	}

	longestSeq := 0
	for i, v := range a {
		for j := i - 1; j >= 0; j-- {
			if v >= a[j] && longestEndingAtPos[j]+1 > longestEndingAtPos[i] {
				longestEndingAtPos[i] = longestEndingAtPos[j] + 1
			}
		}

		if longestEndingAtPos[i] > longestSeq {
			longestSeq = longestEndingAtPos[i]
		}
	}

	return longestSeq
}
