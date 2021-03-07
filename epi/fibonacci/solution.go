package fibonacci

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}

	return current
}

func FibonacciCacheRecursive(n int) int {
	cache := make(map[int]int)
	return fibCacheHelper(n, cache)
}

func fibCacheHelper(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}
	_, ok := cache[n]
	if !ok {
		cache[n] = fibCacheHelper(n-1, cache) + fibCacheHelper(n-2, cache)
	}
	return cache[n]
}
