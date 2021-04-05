package primitive_divide

func Divide(x int, y int) int {
	yPower := 0
	longY := int64(y)
	longX := int64(x)

	// calculate the biggest power of 2 that multiplies by y is inferior to x
	for longY<<(yPower+1) <= longX {
		yPower += 1
	}

	result := 0
	for longX >= longY {
		// skip the powers of 2 that are missing from the decomposition
		for (longY << yPower) > longX {
			yPower -= 1
		}

		result += (1 << yPower)
		longX -= (longY << yPower)
		yPower -= 1
	}

	return result
}
