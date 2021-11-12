package is_valid_sudoku

import "math"

func IsValidSudoku(partialAssignment [][]int) bool {
	if len(partialAssignment) == 0 || len(partialAssignment[0]) == 0 {
		panic("this is not a sudoku board!")
	}
	n := len(partialAssignment)
	squareSize := int(math.Sqrt(float64(n)))
	rowsSets := make([]map[int]bool, n)
	colsSets := make([]map[int]bool, n)
	squareSets := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		rowsSets[i] = make(map[int]bool)
		colsSets[i] = make(map[int]bool)
		squareSets[i] = make(map[int]bool)
	}

	for rIdx, row := range partialAssignment {
		for cIdx, val := range row {
			if val == 0 {
				continue
			}
			if rowsSets[rIdx][val] {
				return false
			}
			if colsSets[cIdx][val] {
				return false
			}
			sqIdx := squareIdx(rIdx, cIdx, squareSize)
			if squareSets[sqIdx][val] {
				return false
			}
			rowsSets[rIdx][val] = true
			colsSets[cIdx][val] = true
			squareSets[sqIdx][val] = true
		}
	}

	return true
}

func squareIdx(rowIdx, colIdx, squareSize int) int {
	return squareSize*(rowIdx/squareSize) + (colIdx / squareSize)
}
