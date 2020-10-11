package three_sum

import (
	"sort"
)

func HasThreeSum(a []int, t int) bool {
	if len(a) == 0 {
		return false
	}

	sort.Ints(a)

	for _, v := range a {
		if hasTwoSum(a, t-v) {
			return true
		}
	}

	return false
}

// hasTwoSum finds a two sum in a already sorted array a
func hasTwoSum(a []int, t int) bool {
	lowIdx, highIdx := 0, len(a)-1
	for lowIdx <= highIdx {
		switch {
		case a[lowIdx]+a[highIdx] == t:
			return true
		case a[lowIdx]+a[highIdx] < t:
			lowIdx++
		case a[lowIdx]+a[highIdx] > t:
			highIdx--
		}
	}
	return false
}
