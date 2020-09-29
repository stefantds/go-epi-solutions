package primitive_multiply

func PrimitiveMultiply(x int64, y int64) int64 {
	sum := int64(0)
	for x != 0 {
		if x&1 != 0 {
			sum = primitiveAdd(sum, y)
		}
		y <<= 1
		x >>= 1
	}
	return sum
}

func primitiveAdd(x int64, y int64) int64 {
	for y != 0 {
		carry := x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}
