package int_as_array_multiply

func Multiply(num1 []int, num2 []int) []int {
	negative := (num1[0] < 0) != (num2[0] < 0)

	if num1[0] < 0 {
		num1[0] = -num1[0]
	}

	if num2[0] < 0 {
		num2[0] = -num2[0]
	}

	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			res := result[i+j+1] + num1[i]*num2[j]
			result[i+j+1] = res % 10
			result[i+j] = result[i+j] + res/10
		}
	}

	// find the first non zero value to exclude the leading zeros
	var firstNonZero int
	for i, v := range result {
		if v != 0 {
			firstNonZero = i
			break
		}
	}

	if firstNonZero == 0 && result[0] == 0 {
		return []int{0} // expected zero result
	}

	result = result[firstNonZero:]
	if negative {
		result[0] = -result[0]
	}

	return result
}
