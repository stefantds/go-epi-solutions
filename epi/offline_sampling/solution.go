package offline_sampling

import "math/rand"

func RandomSampling(k int, a []int) {
	for i := 0; i < k; i++ {
		randPos := i + rand.Intn(len(a)-i)
		a[i], a[randPos] = a[randPos], a[i]
	}
}
