package swap_bits

func SwapBits(x int64, i int, j int) int64 {
	if i >= 64 || j >= 64 || i < 0 || j < 0 {
		panic("the bit index must be between 0 and 63")
	}

	iValue := (x >> i) & 1
	jValue := (x >> j) & 1
	if iValue == jValue {
		return x // bits at index i and j have the same value, no need to swap
	}

	iMask := int64(1 << i)
	jMask := int64(1 << j)

	// bits at index i and j have different value. It's enough to flip their value
	// using XOR
	return x ^ iMask ^ jMask
}
