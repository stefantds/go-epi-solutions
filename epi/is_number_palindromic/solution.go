package is_number_palindromic

import "math"

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	numDigits := int(math.Log10(float64(x))) + 1
	powerTen := int(math.Pow10(numDigits - 1))

	for numDigits > 1 {
		if x%10 != x/powerTen {
			return false
		} else {
			x %= powerTen
			x /= 10
			powerTen /= 100
			numDigits -= 2
		}
	}
	return true
}
