package max_sum_subarray

import (
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func FindMaximumSubarray(a []int) int {
	max := 0 // 0 is the sum for a empty subarray
	currentSum := 0
	for _, v := range a {
		currentSum = utils.Max(currentSum+v, v)
		max = utils.Max(max, currentSum)
	}

	return max
}
