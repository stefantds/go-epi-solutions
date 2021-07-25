package number_of_traversals_staircase

func NumberOfWaysToTop(top int, maximumStep int) int {
	cache := make([]int, top+1)
	cache[0] = 1
	for i := 0; i <= top; i++ {
		for j := 1; j <= maximumStep; j++ {
			if i+j <= top {
				cache[i+j] += cache[i]
			}
		}
	}

	return cache[top]
}
