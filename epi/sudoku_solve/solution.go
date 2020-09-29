package sudoku_solve

import "math"

const emptyEntry = 0

func SolveSudoku(partialAssignment [][]int) bool {
	return solvePartialSudoku(0, 0, partialAssignment)
}

func solvePartialSudoku(i, j int, partialAssignment [][]int) bool {
	if i == len(partialAssignment) {
		i = 0
		j++
		if j == len(partialAssignment[i]) {
			return true
		}
	}

	if partialAssignment[i][j] != 0 {
		return solvePartialSudoku(i+1, j, partialAssignment)
	}

	for val := 1; val <= len(partialAssignment); val++ {
		if validToAddVal(partialAssignment, i, j, val) {
			partialAssignment[i][j] = val
			if solvePartialSudoku(i+1, j, partialAssignment) {
				return true
			}
		}
	}

	partialAssignment[i][j] = emptyEntry // undo unsuccessful attempt
	return false
}

func validToAddVal(partialAssignment [][]int, i, j, val int) bool {
	// check column constraints
	for _, r := range partialAssignment {
		if r[j] == val {
			return false
		}
	}

	// check row constraints
	for _, c := range partialAssignment[i] {
		if c == val {
			return false
		}
	}

	regionSize := int(math.Sqrt(float64(len(partialAssignment))))
	iReg := i / regionSize
	jReg := j / regionSize

	for a := iReg * regionSize; a < (iReg+1)*regionSize; a++ {
		for b := jReg * regionSize; b < (jReg+1)*regionSize; b++ {
			if val == partialAssignment[a][b] {
				return false
			}
		}
	}

	return true
}
