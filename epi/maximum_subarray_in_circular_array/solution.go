package maximum_subarray_in_circular_array

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func MaxSubarraySumInCircular(a []int) int {
	// calculate the max wraping sum first
	var maxWrapingSum int

	// maxStarts[i] is the max sum for subarrays a[:j], with j<=i
	maxStarts := make([]int, len(a))
	maxStarts[0] = a[0]
	currSum := a[0]
	for i := 1; i < len(a); i++ {
		maxStarts[i] = test_utils.Max(maxStarts[i-1], currSum+a[i])
		currSum += a[i]
	}

	// maxEnds[i] is the max sum for subarrays a[j:], with j>=i
	maxEnds := make([]int, len(a))
	maxEnds[len(a)-1] = a[len(a)-1]
	currSum = a[len(a)-1]
	for i := len(a) - 2; i >= 0; i-- {
		maxEnds[i] = test_utils.Max(maxEnds[i+1], currSum+a[i])
		currSum += a[i]
	}

	// find the position with the max wrapping subarray by
	// finding the max value of maxStarts[i]+maxEnds[i+1]
	for i := 0; i < len(a)-1; i++ {
		maxWrapingSum = test_utils.Max(maxWrapingSum, maxStarts[i]+maxEnds[i+1])
	}

	// calculate the max non-wraping subarray sum using Kadane's algorithm
	maxNonWrapingSum := maxSubarrayNonCircular(a)

	// the maximum in the circular array is the max of the two cases
	return test_utils.Max(maxWrapingSum, maxNonWrapingSum)
}

func MaxSubarraySumInCircular_MinSubstract(a []int) int {
	minSubarraySum := 0
	currentMin := 0
	for _, v := range a {
		currentMin = test_utils.Min(currentMin+v, v)
		minSubarraySum = test_utils.Min(currentMin, minSubarraySum)
	}

	totalSum := 0
	for _, v := range a {
		totalSum += v
	}

	// max wraping sum is the total sum - min subarray sum
	maxWrapingSum := totalSum - minSubarraySum

	// calculate the max non-wraping subarray sum using Kadane's algorithm
	maxNonWrapingSum := maxSubarrayNonCircular(a)

	// the maximum in the circular array is the max of the two cases
	return test_utils.Max(maxWrapingSum, maxNonWrapingSum)
}

func maxSubarrayNonCircular(a []int) int {
	maxSum := math.MinInt
	currentMax := 0
	for _, v := range a {
		currentMax = test_utils.Max(currentMax+v, v)
		maxSum = test_utils.Max(currentMax, maxSum)
	}

	return maxSum
}
