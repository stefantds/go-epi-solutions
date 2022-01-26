package count_inversions

func CountInversions(a []int) int {
	return countInversionsInRange(a, 0, len(a))
}

// countInversionsInRange returns the number of inversions in a[start:end]
func countInversionsInRange(a []int, start, end int) int {
	if end-start <= 1 {
		return 0
	}

	mid := (end-start)/2 + start
	countInvFirstHalf := countInversionsInRange(a, start, mid)
	countInvSecondHalf := countInversionsInRange(a, mid, end)

	// count inversions between the two halves and sort a between indices start and end
	sorted := make([]int, 0, end-start)
	countInvBetweenHalves := 0
	leftIter, rightIter := start, mid
	for leftIter < mid && rightIter < end {
		if a[leftIter] <= a[rightIter] {
			sorted = append(sorted, a[leftIter])
			leftIter += 1
		} else {
			// element a[right] is smaller than mid-rightIter elements from the left half
			countInvBetweenHalves += mid - leftIter
			sorted = append(sorted, a[rightIter])
			rightIter += 1
		}
	}

	sorted = append(sorted, a[leftIter:mid]...)
	sorted = append(sorted, a[rightIter:end]...)

	copy(a[start:end], sorted)
	return countInvFirstHalf + countInvSecondHalf + countInvBetweenHalves
}
