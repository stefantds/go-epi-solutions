package matrix_rotation

func RotateMatrix(squareMatrix [][]int) {
	for row := 0; row < len(squareMatrix)/2; row++ {
		for col := row; col < len(squareMatrix)-row-1; col++ {
			oppRow := len(squareMatrix) - 1 - row
			oppCol := len(squareMatrix) - 1 - col

			tmp := squareMatrix[row][col]

			squareMatrix[row][col] = squareMatrix[oppCol][row]
			squareMatrix[oppCol][row] = squareMatrix[oppRow][oppCol]
			squareMatrix[oppRow][oppCol] = squareMatrix[col][oppRow]
			squareMatrix[col][oppRow] = tmp
		}
	}
}
