package max_trapped_water

import (
	"math"

	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func GetMaxTrappedWater(heights []int) int {
	maxVolume := math.MinInt32
	for i, j := 0, len(heights)-1; i < j; /* no increment */ {
		volume := waterVolume(heights, i, j)
		if volume > maxVolume {
			maxVolume = volume
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return maxVolume
}

// waterVolume calculates the water volume that would be trapped between two indices
func waterVolume(heights []int, i int, j int) int {
	return utils.Min(heights[i], heights[j]) * utils.Abs(i-j)
}
