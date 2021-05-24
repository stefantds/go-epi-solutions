package search_row_col_sorted_matrix

func MatrixSearch(a [][]int, x int) bool {
	if len(a) == 0 {
		return false
	}
	row := 0
	col := len(a[0]) - 1

	for row < len(a) && col >= 0 {
		if a[row][col] == x {
			return true
		}

		if a[row][col] < x {
			row++
		} else {
			col--
		}
	}

	return false
}
