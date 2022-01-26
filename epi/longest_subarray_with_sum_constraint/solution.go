package longest_subarray_with_sum_constraint

import "github.com/stefantds/go-epi-judge/test_utils"

func FindLongestSubarrayLessEqualK(a []int, k int) int {
	// build the prefix sum array
	prefixSums := make([]int, len(a))
	for i := range a {
		if i == 0 {
			prefixSums[i] = a[i]
		} else {
			prefixSums[i] = a[i] + prefixSums[i-1]
		}
	}

	if prefixSums[len(a)-1] <= k {
		return len(a)
	}

	// build the min prefix sum array
	minPrefixSums := make([]int, len(a))
	copy(minPrefixSums, prefixSums)
	for i := len(a) - 2; i >= 0; i-- {
		if minPrefixSums[i+1] < minPrefixSums[i] {
			minPrefixSums[i] = minPrefixSums[i+1]
		}
	}

	x, y, maxLength := 0, 0, 0
	for x < len(a) && y < len(a) {
		var minCurrSum int
		if x > 0 {
			minCurrSum = minPrefixSums[y] - prefixSums[x-1]
		} else {
			minCurrSum = minPrefixSums[y]
		}
		if minCurrSum <= k {
			currLength := y - x + 1
			maxLength = test_utils.Max(currLength, maxLength)
			y++
		} else {
			x++
		}
	}

	return maxLength
}
