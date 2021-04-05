package reverse_digits

func Reverse(x int) int64 {
	result := int64(0)
	for x != 0 {
		result = result*10 + int64(x%10)
		x = x / 10
	}

	return result
}
