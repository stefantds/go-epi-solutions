package n_queens

func NQueens(n int) [][]int {
	result := make([][]int, 0)
	nQueensHelper(n, 0, []int{}, &result)
	return result
}

func nQueensHelper(n int, row int, colPlacement []int, result *[][]int) {
	if row == n {
		placementCopy := append([]int{}, colPlacement...)
		*result = append(*result, placementCopy)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(colPlacement, col) {
			nQueensHelper(n, row+1, append(colPlacement, col), result)
		}
	}
}

// isValid returns true if newItem can be added to the existing values
// in currentPlacement without violating the non-attacking invariant.
func isValid(currentPlacement []int, newItem int) bool {
	newItemIdx := len(currentPlacement)
	for i, v := range currentPlacement {
		switch {
		case v == newItem:
			return false
		case newItemIdx-i == v-newItem:
			return false
		case newItemIdx-i == newItem-v:
			return false
		}
	}
	return true
}

/**
q _ _ _ _
_ q _ _ _
_ = x _ _
_ = _ q _
_ _ = = x

q x x x
x x _ _
x _ q _
x _ _ x

(n-2)*(3n-2) < n2, n >= 4
**/
