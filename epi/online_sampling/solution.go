package online_sampling

import "math/rand"

func OnlineRandomSample(stream <-chan int, k int) []int {
	numSeen := 0
	vals := make([]int, k)
	for v := range stream {
		vals[numSeen] = v
		numSeen++
		if numSeen >= k {
			break
		}
	}

	for v := range stream {
		// generate a random position in [0, numSeen + 1]. If the position
		// is in [0, k], override the element at that position.
		nextItemPos := rand.Intn(numSeen + 1)
		if nextItemPos < k {
			vals[nextItemPos] = v
		}
		numSeen++
	}

	return vals
}
