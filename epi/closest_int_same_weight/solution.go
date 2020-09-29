package closest_int_same_weight

import "errors"

func ClosestIntSameBitCount(x int64) int64 {
	const numBits = 63

	for i := uint(0); i < numBits-1; i++ {
		// check if two consecutive bits are different
		if x>>i&1 != x>>(i+1)&1 {
			x ^= 1<<i | 1<<(i+1) // swap bits at position i and i+1
			return x
		}
	}

	panic(errors.New("all bits are 0 or 1"))
}
