package search_first_key

func SearchFirstOfK(a []int, k int) int {
	right := len(a) - 1
	left := 0
	mid := (right + left) / 2
	found := -1

	for left <= right {
		switch {
		case a[mid] == k:
			found = mid
			right = mid - 1
			mid = (right + left) / 2
		case a[mid] < k:
			left = mid + 1
			mid = (right + left) / 2
		case a[mid] > k:
			right = mid - 1
			mid = (right + left) / 2
		}
	}

	return found
}
