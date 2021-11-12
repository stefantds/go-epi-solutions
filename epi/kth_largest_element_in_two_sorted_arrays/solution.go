package kth_largest_element_in_two_sorted_arrays

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func FindKthNTwoSortedArrays(a []int, b []int, k int) int {
	left := test_utils.Max(0, k-len(b))
	right := test_utils.Min(k, len(a))

	for left < right {
		aPos := left + ((right - left) / 2)
		var aPrevVal, aPosVal int
		if aPos <= 0 {
			aPrevVal = math.MinInt
		} else {
			aPrevVal = a[aPos-1]
		}
		if aPos >= len(a) {
			aPosVal = math.MaxInt
		} else {
			aPosVal = a[aPos]
		}

		bPos := k - aPos
		var bPrevVal, bPosVal int
		if bPos <= 0 {
			bPrevVal = math.MinInt
		} else {
			bPrevVal = b[bPos-1]
		}
		if bPos >= len(b) {
			bPosVal = math.MaxInt
		} else {
			bPosVal = b[bPos]
		}

		if aPrevVal > bPosVal {
			right = aPos - 1
		} else if aPosVal < bPrevVal {
			left = aPos + 1
		} else {
			return test_utils.Max(aPrevVal, bPrevVal)
		}
	}

	// right == left
	var aPrevVal, bPrevVal int
	if left <= 0 {
		aPrevVal = math.MinInt
	} else {
		aPrevVal = a[left-1]
	}

	if k-left <= 0 {
		bPrevVal = math.MinInt
	} else {
		bPrevVal = b[k-left-1]
	}

	return test_utils.Max(aPrevVal, bPrevVal)
}
