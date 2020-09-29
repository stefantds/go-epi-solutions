package search_shifted_sorted_array

func SearchSmallest(a []int) int {
	return searchHelper(0, len(a)-1, a)
}

func searchHelper(left, right int, a []int) int {
	if a[left] < a[right] { // in case the array is not shifted
		return left
	}

	if left == right {
		return right
	}

	mid := left + (right-left)/2
	if a[mid] < a[left] {
		return searchHelper(left+1, mid, a)
	} else {
		return searchHelper(mid+1, right, a)
	}
}
