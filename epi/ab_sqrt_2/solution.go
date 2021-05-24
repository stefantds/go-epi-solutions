package ab_sqrt_2

import (
	"math"
)

type abSqrt2 struct {
	a, b int
}

func (s abSqrt2) value() float64 {
	return float64(s.a) + float64(s.b)*math.Sqrt(2)
}

func GenerateFirstKABSqrt2(k int) []float64 {
	candidates := make([]abSqrt2, 0)
	i, j := 0, 0
	candidates = append(candidates, abSqrt2{a: 0, b: 0})
	for len(candidates) < k {
		candidateI := abSqrt2{a: candidates[i].a + 1, b: candidates[i].b}
		candidateJ := abSqrt2{a: candidates[j].a, b: candidates[j].b + 1}
		comp := compare(candidateI, candidateJ)
		if comp < 0 {
			candidates = append(candidates, candidateI)
			i++
		}
		if comp > 0 {
			candidates = append(candidates, candidateJ)
			j++
		}
		if comp == 0 {
			candidates = append(candidates, candidateI)
			i, j = i+1, j+1
		}
	}

	result := make([]float64, k)
	for i, c := range candidates {
		result[i] = c.value()
	}
	return result
}

func compare(v1, v2 abSqrt2) int {
	if v1.a == v2.a && v1.b == v2.b {
		return 0
	}
	if v1.value() < v2.value() {
		return -1
	}

	return 1
}
