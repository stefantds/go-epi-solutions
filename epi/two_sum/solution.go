package two_sum

import (
	"sort"
)

func HasTwoSum(a []int, t int) bool {
	if len(a) == 0 {
		return false
	}

	sort.Ints(a)
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
