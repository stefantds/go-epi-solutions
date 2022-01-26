package gcd

func GCD(x int64, y int64) int64 {
	if x > y {
		x, y = y, x
	}
	if x == 0 {
		return y
	}
	switch {
	case x&1 == 0 && y&1 == 0:
		return GCD(x>>1, y>>1) << 1 // 2 * GCD(x/2, y/2)
	case x&1 == 0 && y&1 != 0:
		return GCD(x>>1, y)
	case x&1 != 0 && y&1 == 0:
		return GCD(x, y>>1)
	default:
		// both are odd: use difference
		return GCD(x, y-x)
	}
}
