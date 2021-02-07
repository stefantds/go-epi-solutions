package real_square_root

import "math"

const epsilon = float64(0.000001)

func SquareRootReal(x float64) float64 {
	var left, right float64
	if x < 1 {
		left, right = x, 1
	} else {
		left, right = 1, x
	}

	for !equal(left, right) {
		mid := left + (right-left)/2
		if mid*mid < x {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
