package max_product_all_but_one

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func FindBiggestProductNMinusOneProduct(a []int) int {
	// suffixProducts[i] is the product of a[i+1:]
	suffixProducts := make([]int, len(a))
	suffixProducts[len(a)-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		suffixProducts[i] = suffixProducts[i+1] * a[i+1]
	}

	// prefixProducts[i] is the product of a[:i]
	prefixProducts := make([]int, len(a))
	prefixProducts[0] = 1
	for i := 1; i < len(a); i++ {
		prefixProducts[i] = prefixProducts[i-1] * a[i-1]
	}

	max := math.MinInt32
	for i := range a {
		product := prefixProducts[i] * suffixProducts[i]
		max = test_utils.Max(max, product)
	}

	return max
}

func FindBiggestProductNMinusOneProduct_NoExtraSpace(a []int) int {
	productWithout := func(excludeIndex int) int {
		product := 1
		for i, v := range a {
			if i != excludeIndex {
				product *= v
			}
		}

		return product
	}

	var countNegativeVals, countPositiveVals int
	var minPositiveIdx, minNegativeIdx, maxNegativeIdx int
	minPositive := math.MaxInt64
	minNegative := 0
	maxNegative := math.MinInt64
	for i, v := range a {
		if v < 0 {
			countNegativeVals++
			if v < minNegative {
				minNegative = v
				minNegativeIdx = i
			}
			if v > maxNegative {
				maxNegative = v
				maxNegativeIdx = i
			}
		} else {
			countPositiveVals++
			if v < minPositive {
				minPositive = v
				minPositiveIdx = i
			}
		}
	}

	if countNegativeVals == 0 {
		return productWithout(minPositiveIdx)
	}
	if countNegativeVals%2 == 0 {
		if countPositiveVals > 0 {
			return productWithout(minPositiveIdx)
		}
		return productWithout(minNegativeIdx)
	}
	// countNegativeVals is odd
	return productWithout(maxNegativeIdx)
}
