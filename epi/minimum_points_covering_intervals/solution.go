package minimum_points_covering_intervals

import (
	"math"
	"sort"
)

func FindMinimumVisits(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Right <= intervals[j].Right {
			return true
		}
		return false
	})

	lastVisit := math.MinInt64
	visits := 0
	for _, e := range intervals {
		if lastVisit < e.Left {
			lastVisit = e.Right
			visits++
		}
	}

	return visits
}
