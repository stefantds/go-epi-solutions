package kth_largest_in_array

import "math/rand"

func FindKthLargest(k int, a []int) int {
	left, right := 0, len(a)-1

	for left <= right {
		pivotIdx := left + rand.Intn(right-left+1) // take a random position in [right, left]
		newPivotIdx := partition(a, left, right, pivotIdx)

		switch {
		case newPivotIdx < k-1:
			left = newPivotIdx + 1
		case newPivotIdx > k-1:
			right = newPivotIdx - 1
		case newPivotIdx == k-1:
			return a[newPivotIdx]
		}
	}

	return -1
}

// partition partitions the slice a between position left and right (inclusive) as follows:
// elements a[left : pivotIdx] are greater than the pivot
// elements a[pivotIdx+1:right+1] are less than the pivot
// It returns the pivotIdx.
func partition(a []int, left, right, pivot int) int {
	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}

	pValue := a[pivot]
	swap(pivot, right)
	newPivotIdx := left
	for i := left; i < right; i++ {
		if a[i] > pValue {
			swap(i, newPivotIdx)
			newPivotIdx++
		}
	}

	swap(right, newPivotIdx)
	return newPivotIdx
}
