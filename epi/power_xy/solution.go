package power_xy

func Power(x float64, y int) float64 {
	result := float64(1)

	if y < 0 {
		y = -y
		x = 1 / x
	}

	for y != 0 {
		if y&1 != 0 {
			result *= x
		}
		x *= x
		y >>= 1
	}

	return result
}
