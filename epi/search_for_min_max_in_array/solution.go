package search_for_min_max_in_array

import "github.com/stefantds/go-epi-judge/test_utils"

func FindMinMax(a []int) (min, max int) {
	if len(a) == 0 {
		panic("invalid input: empty array")
	}
	if len(a) == 1 {
		return a[0], a[0]
	}

	min, max = minMax(a[0], a[1])

	for i := 2; i < len(a)-1; i += 2 {
		newMin, newMax := minMax(a[i], a[i+1])
		min = test_utils.Min(min, newMin)
		max = test_utils.Max(max, newMax)
	}

	if len(a)%2 == 1 {
		// the last element doesn't have a pair so it couldn't be compared yet
		min = test_utils.Min(min, a[len(a)-1])
		max = test_utils.Max(max, a[len(a)-1])
	}

	return min, max
}

// minMax compares two int values and returns the smaller one
// followed by the bigger one
func minMax(value1, value2 int) (min, max int) {
	if value1 < value2 {
		return value1, value2
	}
	return value2, value1
}
