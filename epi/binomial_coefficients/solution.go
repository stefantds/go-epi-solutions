package binomial_coefficients

func ComputeBinomialCoefficient(n int, k int) int {
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, k+1)
	}

	return countSubsets(n, k, cache)
}

func countSubsets(n, k int, cache [][]int) int {
	if k == 0 || n == k {
		return 1
	}
	if cache[n][k] == 0 {
		withoutCurrent := countSubsets(n-1, k, cache)
		withCurrent := countSubsets(n-1, k-1, cache)
		cache[n][k] = withoutCurrent + withCurrent
	}

	return cache[n][k]
}
