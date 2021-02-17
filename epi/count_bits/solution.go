package count_bits

func CountBits(x int) int {
	counter := 0
	for x != 0 {
		counter += x & 1
		x = x >> 1
	}

	return counter
}

// CountBitsV2 uses the lowest set bit to skip zero bits
func CountBitsV2(x int) int {
	counter := 0
	for x != 0 {
		counter += 1
		x = x & (x - 1) // drop lowest set bit of x
	}

	return counter
}
