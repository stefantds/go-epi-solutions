package spiral_ordering

func MatrixInSpiralOrder(squareMatrix [][]int) []int {
	if len(squareMatrix) == 0 {
		return []int{}
	}

	n := len(squareMatrix)
	result := make([]int, 0, n*n)
	for i := 0; i < (n+1)/2; i++ {
		spiralOneLayer(squareMatrix, i, n-1-i, &result)
	}

	return result
}

func spiralOneLayer(m [][]int, start int, end int, result *[]int) {
	// special case for the center of the matrix
	if start == end {
		*result = append(*result, m[start][end])
		return
	}

	for i := start; i < end; i++ {
		*result = append(*result, m[start][i])
	}
	for i := start; i < end; i++ {
		*result = append(*result, m[i][end])
	}
	for i := end; i > start; i-- {
		*result = append(*result, m[end][i])
	}
	for i := end; i > start; i-- {
		*result = append(*result, m[i][start])
	}
}
