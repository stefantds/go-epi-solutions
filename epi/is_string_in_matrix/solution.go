package is_string_in_matrix

func IsPatternContainedInGrid(grid [][]int, pattern []int) bool {
	if len(grid) == 0 {
		return false
	}

	// cache will remember for every position if a match starting from every
	// index in the pattern was found. The cache will initially contain
	// nil values.
	cache := make([][][]*bool, len(grid))
	for i := range grid {
		cache[i] = make([][]*bool, len(grid[0]))
		for j := range grid[0] {
			cache[i][j] = make([]*bool, len(pattern))
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if searchPattern(grid, pattern, i, j, 0, &cache) {
				return true
			}
		}
	}
	return false
}

func searchPattern(grid [][]int, pattern []int, x, y, patternIdx int, cache *[][][]*bool) bool {
	var result bool
	if len(grid) == 0 {
		return false
	}
	if patternIdx == len(pattern) {
		return true
	}
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	if (*cache)[x][y][patternIdx] != nil {
		return *(*cache)[x][y][patternIdx]
	}
	if grid[x][y] == pattern[patternIdx] {
		result = searchPattern(grid, pattern, x+1, y, patternIdx+1, cache) ||
			searchPattern(grid, pattern, x-1, y, patternIdx+1, cache) ||
			searchPattern(grid, pattern, x, y+1, patternIdx+1, cache) ||
			searchPattern(grid, pattern, x, y-1, patternIdx+1, cache)
	}
	(*cache)[x][y][patternIdx] = &result
	return result
}
