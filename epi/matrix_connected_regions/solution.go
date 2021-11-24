package matrix_connected_regions

func FlipColor(x int, y int, image [][]bool) {
	flip(x, y, &image, !image[x][y])
}

func flip(x int, y int, image *[][]bool, flipColor bool) {
	if x < 0 || y < 0 || x >= len(*image) || y >= len((*image)[x]) || (*image)[x][y] == flipColor {
		return
	}

	(*image)[x][y] = flipColor

	flip(x-1, y, image, flipColor)
	flip(x+1, y, image, flipColor)
	flip(x, y-1, image, flipColor)
	flip(x, y+1, image, flipColor)
}
