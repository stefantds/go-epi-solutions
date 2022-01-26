package max_safe_height

func GetHeight(cases int, drops int) int {
	cache := make([][]int, cases+1)
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, drops+1)
		cache[i][0] = 0
	}
	for i := 0; i < drops; i++ {
		cache[0][drops] = 0
	}
	for c := 1; c <= cases; c++ {
		for d := 1; d <= drops; d++ {
			cache[c][d] = 1 + cache[c-1][d-1] + cache[c][d-1]
		}
	}

	return cache[cases][drops]
}
