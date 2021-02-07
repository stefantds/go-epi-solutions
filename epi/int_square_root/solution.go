package int_square_root

func SquareRoot(k int) int {
	return searchRoot(0, k, k)
}

func searchRoot(low, high, k int) int {
	if low > high {
		return low - 1
	}

	mid := low + (high-low)/2
	midSquared := mid * mid

	switch {
	case midSquared == k:
		return mid
	case midSquared < k:
		return searchRoot(mid+1, high, k)
	default:
		return searchRoot(low, mid-1, k)
	}
}
