package max_submatrix

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

// maxRanges represents the ranges of "true" values for a given position
// going up or left from that position
type maxRanges struct {
	left, up int
}

func MaxRectangleSubmatrix(a [][]bool) int {
	if len(a) == 0 {
		return 0
	}
	maxRangesMatrix := make([][]maxRanges, len(a))
	for i := 0; i < len(a); i++ {
		maxRangesMatrix[i] = make([]maxRanges, len(a[0]))
	}

	maxSizeFound := 0
	for i := range a {
		for j := range a[i] {
			// calculate the max ranges towards up and left from the current position
			if !a[i][j] {
				maxRangesMatrix[i][j] = maxRanges{up: 0, left: 0}
				continue
			}
			if i == 0 {
				maxRangesMatrix[i][j].up = 1
			} else {
				maxRangesMatrix[i][j].up = maxRangesMatrix[i-1][j].up + 1
			}
			if j == 0 {
				maxRangesMatrix[i][j].left = 1
			} else {
				maxRangesMatrix[i][j].left = maxRangesMatrix[i][j-1].left + 1
			}

			// iterate back to find the max submatrix with the bottom right
			// corner at the current position
			minHeight := math.MaxInt32
			currentMaxFound := 0
			for k := j; k >= 0; k-- {
				minHeight = test_utils.Min(minHeight, maxRangesMatrix[i][k].up)
				currentSize := minHeight * (j - k + 1)
				currentMaxFound = test_utils.Max(currentMaxFound, currentSize)
			}

			if currentMaxFound > maxSizeFound {
				maxSizeFound = currentMaxFound
			}
		}
	}

	return maxSizeFound
}
