package nonuniform_random_number

import "math/rand"

type interval struct {
	start, end float64
}

func NonuniformRandomNumberGeneration(values []int, probabilities []float64) int {
	probDistr := make([]interval, 0, len(probabilities))
	curr := float64(0)
	for _, v := range probabilities {
		probDistr = append(probDistr, interval{start: curr, end: curr + v})
		curr += v
	}

	randVal := rand.Float64()

	pos := findIntervalIndex(probDistr, 0, len(probDistr)-1, randVal)
	if pos == -1 {
		panic("internal error: value not found in aby interval")
	}

	return values[pos]
}

// findIntervalIndex returns the index of the interval containing
// the value.
// The intervals are assumed to be sorted and contiguous.
// The intervals are assumed closed to the left and open to the right.
func findIntervalIndex(intervals []interval, start, end int, value float64) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	i := intervals[mid]
	if value >= i.start && value < i.end {
		return mid
	}
	if value < i.start {
		return findIntervalIndex(intervals, start, mid-1, value)
	}
	return findIntervalIndex(intervals, mid+1, end, value)
}
