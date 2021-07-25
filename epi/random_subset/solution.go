package random_subset

import "math/rand"

func RandomSubset(n int, k int) []int {
	modififed := make(map[int]int, k)
	result := make([]int, k)

	valueAt := func(pos int) int {
		if v, ok := modififed[pos]; ok {
			return v
		} else {
			return pos
		}
	}

	for i := 0; i < k; i++ {
		randPos := i + rand.Intn(n-i)
		val := valueAt(randPos)
		result[i] = val

		valAtI := valueAt(i)
		modififed[randPos] = valAtI
	}

	return result
}
