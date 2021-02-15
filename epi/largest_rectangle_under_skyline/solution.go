package largest_rectangle_under_skyline

import (
	"github.com/stefantds/go-epi-judge/utils"
)

func CalculateLargestRectangle(heights []int) int {
	maxRectangle := 0
	candidates := make([]int, 0)

	for i := 0; i <= len(heights); i++ {
		for len(candidates) > 0 &&
			(i == len(heights) /* for the last element all candidates must be considered */ ||
				heights[candidates[len(candidates)-1]] >= heights[i]) {
			var head int
			candidates, head = candidates[:len(candidates)-1], candidates[len(candidates)-1]

			var area int
			{
				height := heights[head]
				var width int
				if len(candidates) == 0 {
					width = i
				} else {
					width = i - candidates[len(candidates)-1] - 1
				}
				area = height * width
			}

			maxRectangle = utils.Max(maxRectangle, area)
		}

		candidates = append(candidates, i)
	}

	return maxRectangle
}
