package reverse_bits

func ReverseBits(x int64) int64 {
	var result int64
	var i int
	for i = 0; x != 0; i++ {
		result <<= 1
		result |= x & 1
		x >>= 1
	}

	// add leading 0s from x to the end of result
	if 64-i > 0 {
		result <<= uint(64 - i)
	}

	return result
}
