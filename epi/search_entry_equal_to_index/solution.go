package search_entry_equal_to_index

func SearchEntryEqualToItsIndex(a []int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case a[mid] == mid:
			return mid
		case a[mid] > mid:
			right = mid - 1
		case a[mid] < mid:
			left = mid + 1
		}
	}

	return -1
}
