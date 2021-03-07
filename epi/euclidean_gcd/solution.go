package euclidean_gcd

func EuclideanGCD(x int64, y int64) int64 {
	if y > x {
		x, y = y, x
	}
	if y == 0 {
		return x
	}
	return EuclideanGCD(y, x%y)
}
