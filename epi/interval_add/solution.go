package interval_add

import "github.com/stefantds/go-epi-judge/test_utils"

func AddInterval(disjointIntervals []Interval, newInterval Interval) []Interval {
	result := make([]Interval, 0)
	joinedInterval := newInterval
	i := 0
	for i < len(disjointIntervals) && disjointIntervals[i].Right < newInterval.Left {
		result = append(result, disjointIntervals[i])
		i++
	}
	for i < len(disjointIntervals) && disjointIntervals[i].Left <= newInterval.Right {
		joinedInterval.Left = test_utils.Min(joinedInterval.Left, disjointIntervals[i].Left)
		joinedInterval.Right = test_utils.Max(joinedInterval.Right, disjointIntervals[i].Right)
		i++
	}
	result = append(result, joinedInterval)

	for i < len(disjointIntervals) {
		result = append(result, disjointIntervals[i])
		i++
	}

	return result
}
