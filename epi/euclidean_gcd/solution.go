package euclidean_gcd

func EuclideanGCD(x int, y int) int {
	if y > x {
		x, y = y, x
	}
	if y == 0 {
		return x
	}
	return EuclideanGCD(y, x%y)
}
