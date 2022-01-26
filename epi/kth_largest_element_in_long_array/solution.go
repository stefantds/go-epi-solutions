package kth_largest_element_in_long_array

import "github.com/stefantds/go-epi-judge/epi/kth_largest_in_array"

func FindKthLargestUnknownLength(stream <-chan int, k int) int {
	candidates := make([]int, 0, 2*k)
	for v := range stream {
		candidates = append(candidates, v)
		if len(candidates) == 2*k {
			// FindKthLargest reorders the candidates such that the k largest element will be
			// in candidates[:k]
			_ = kth_largest_in_array.FindKthLargest(k, candidates)
			// discard the other elements
			candidates = candidates[:k]
		}
	}

	// return the k-th largest element from the remaining candidates
	return kth_largest_in_array.FindKthLargest(k, candidates)
}
