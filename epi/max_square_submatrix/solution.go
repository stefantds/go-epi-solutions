package max_square_submatrix

import "github.com/stefantds/go-epi-judge/test_utils"

// maxRanges represents the ranges of "true" values for a given position
// the maxSquare is the size of the square submatrix having the given position
// as the bottom right corner that contains only true values.
type maxRanges struct {
	left, up, maxSquare int
}

func MaxSquareSubmatrix(a [][]bool) int {
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
			if !a[i][j] {
				maxRangesMatrix[i][j] = maxRanges{up: 0, left: 0, maxSquare: 0}
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

			if i == 0 || j == 0 {
				maxRangesMatrix[i][j].maxSquare = 1
			} else {
				maxRangesMatrix[i][j].maxSquare = test_utils.Min(
					test_utils.Min(
						maxRangesMatrix[i-1][j-1].maxSquare+1,
						maxRangesMatrix[i][j].left,
					),
					maxRangesMatrix[i][j].up,
				)
			}

			if maxRangesMatrix[i][j].maxSquare > maxSizeFound {
				maxSizeFound = maxRangesMatrix[i][j].maxSquare
			}
		}
	}

	return maxSizeFound * maxSizeFound
}
