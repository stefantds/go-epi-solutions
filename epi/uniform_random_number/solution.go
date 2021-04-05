package uniform_random_number

import (
	"math"
	"math/rand"
)

func zeroOneRandom() int {
	return rand.Intn(2)
}

func UniformRandom(lowerBound int, upperBound int) int {
	limit := upperBound - lowerBound
	result := math.MaxInt64
	for result > limit {
		result = 0
		for i := 0; (1 << i) <= limit; i++ {
			result = (result << 1) | zeroOneRandom()
		}
	}

	return result + lowerBound
}
