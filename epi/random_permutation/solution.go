package random_permutation

import "math/rand"

func ComputeRandomPermutation(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}

	for i := 0; i < n; i++ {
		randPos := i + rand.Intn(n-i)
		result[i], result[randPos] = result[randPos], result[i]
	}

	return result
}
