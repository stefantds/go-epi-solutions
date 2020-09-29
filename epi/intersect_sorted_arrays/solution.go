package intersect_sorted_arrays

func IntersectTwoSortedArrays(a []int, b []int) []int {
	result := make([]int, 0)

	if len(a) == 0 || len(b) == 0 {
		return result
	}

	iterA, iterB := 0, 0

	for iterA < len(a) && iterB < len(b) {
		switch {
		case a[iterA] == b[iterB]:
			// check if the element is already present in the intersection
			if iterA == 0 || a[iterA-1] != a[iterA] {
				result = append(result, a[iterA])
			}
			iterA += 1
			iterB += 1
		case a[iterA] < b[iterB]:
			iterA += 1
		default: // a[iterA] > b[iterB]
			iterB += 1
		}
	}

	return result
}
