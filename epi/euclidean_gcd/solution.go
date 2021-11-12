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

func EuclideanGCDNoRecursion(x int, y int) int {
	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}
	return x
}
