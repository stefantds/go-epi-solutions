package minimum_weight_path_in_a_triangle

import (
	"math"

	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func MinimumPathTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	var prevMins = triangle[0]
	for _, row := range triangle[1:] {
		newMins := make([]int, len(row))
		for idx, v := range row {
			switch idx {
			case 0:
				newMins[idx] = v + prevMins[0]
			case len(row) - 1:
				newMins[idx] = v + prevMins[len(prevMins)-1]
			default:
				newMins[idx] = v + utils.Min(prevMins[idx-1], prevMins[idx])
			}
		}

		prevMins = newMins
	}

	min := math.MaxInt64
	for _, v := range prevMins {
		min = utils.Min(v, min)
	}

	return min
}
