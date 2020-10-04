package number_of_traversals_matrix

func NumberOfWays(n int, m int) int {
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}
	return countTraversals(n-1, m-1, cache)
}

func countTraversals(row int, col int, cache [][]int) int {
	if row < 0 || col < 0 {
		return 0
	}
	if row == 0 && col == 0 {
		return 1
	}
	if cache[row][col] != 0 {
		return cache[row][col]
	}
	t := countTraversals(row-1, col, cache) + countTraversals(row, col-1, cache)
	cache[row][col] = t
	return t
}
