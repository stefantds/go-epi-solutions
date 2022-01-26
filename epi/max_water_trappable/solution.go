package max_water_trappable

import "github.com/stefantds/go-epi-judge/test_utils"

func CalculateTrappingWater(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxIdx := 0
	maxHeight := heights[0]

	for i, h := range heights {
		if h > maxHeight {
			maxHeight = h
			maxIdx = i
		}
	}

	return maxWaterInRange(heights[:maxIdx]) +
		maxWaterInRange(reverse(heights[maxIdx+1:]))
}

// maxWaterInRange calculates the max trapped water, assuming that
// at the end there is a "wall" higher than any value in the slice
func maxWaterInRange(heights []int) int {
	maxSeenSoFar := 0
	trappedWater := 0
	for _, h := range heights {
		if h > maxSeenSoFar {
			maxSeenSoFar = h
		} else {
			trappedWater += maxSeenSoFar - h
		}
	}

	return trappedWater
}

// reverse reverses a slice and returns the resulting (new) slice
func reverse(s []int) []int {
	result := make([]int, len(s))
	n := len(result)
	for i, v := range s {
		result[n-1-i] = v
	}
	return result
}

type wall struct {
	height int
	length int
}

func CalculateTrappingWaterUsingStack(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]wall, 1)
	stack[0] = wall{
		height: heights[0],
		length: 1,
	}
	maxSeenSoFar := heights[0]
	trappedWater := 0

	for i := 0; i < len(heights); i++ {
		top := stack[len(stack)-1]
		currentHeight := heights[i]
		if currentHeight <= top.height {
			stack = append(stack, wall{height: currentHeight, length: 1})
		} else {
			// we've found a valley
			currentWaterHeight := test_utils.Min(maxSeenSoFar, currentHeight)

			// pop from the stack until we find the beginning of the current valley
			length := 1
			for len(stack) > 0 && currentHeight >= stack[len(stack)-1].height {
				stack, top = stack[:len(stack)-1], stack[len(stack)-1]
				trappedWater += (currentWaterHeight - top.height) * top.length
				length += top.length
			}

			if currentHeight > maxSeenSoFar {
				maxSeenSoFar = currentHeight
			}

			// we insert the current height & length in the stack, as it might still
			// be trapped water above the current level
			stack = append(stack, wall{height: currentHeight, length: length})
		}
	}

	return trappedWater
}
