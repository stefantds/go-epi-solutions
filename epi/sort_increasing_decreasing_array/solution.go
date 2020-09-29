package sort_increasing_decreasing_array

import m "github.com/stefantds/go-epi-judge/epi/sorted_arrays_merge"

func SortKIncreasingDecreasingArray(a []int) []int {
	if len(a) < 2 {
		return a
	}

	sortedSubarrays := make([][]int, 0)
	currentSubarray := make([]int, 1)
	currentSubarray[0] = a[0]
	currentIncreasing := a[0] < a[1]

	for i := 1; i < len(a); i++ {
		switch {
		case currentIncreasing && a[i-1] <= a[i]:
			currentSubarray = append(currentSubarray, a[i])
		case currentIncreasing && a[i-1] > a[i]:
			sortedSubarrays = append(sortedSubarrays, currentSubarray)
			currentSubarray = make([]int, 1)
			currentSubarray[0] = a[i]
			currentIncreasing = false
		case !currentIncreasing && a[i-1] >= a[i]:
			currentSubarray = append(currentSubarray, a[i])
		case !currentIncreasing && a[i-1] < a[i]:
			sortedSubarrays = append(sortedSubarrays, reverseArray(currentSubarray))
			currentSubarray = make([]int, 1)
			currentSubarray[0] = a[i]
			currentIncreasing = true
		}
	}

	switch currentIncreasing {
	case true:
		sortedSubarrays = append(sortedSubarrays, currentSubarray)
	case false:
		sortedSubarrays = append(sortedSubarrays, reverseArray(currentSubarray))
	}

	return m.MergeSortedArrays(sortedSubarrays)
}

func reverseArray(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
