package pascal_triangle

func GeneratePascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)
	triangle[0] = []int{1}

	for i := 1; i < numRows; i++ {
		prevRow := triangle[i-1]
		newRow := make([]int, len(prevRow)+1)
		newRow[0] = 1
		for j := 1; j < len(prevRow); j++ {
			newRow[j] = prevRow[j-1] + prevRow[j]
		}
		newRow[len(newRow)-1] = 1
		triangle[i] = newRow
	}

	return triangle
}
