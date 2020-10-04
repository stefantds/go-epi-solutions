package minimum_waiting_time

import (
	"sort"
)

func MinimumTotalWaitingTime(serviceTimes []int) int {
	sort.Ints(serviceTimes)
	total := 0
	current := 0
	for i := 0; i < len(serviceTimes)-1; i++ {
		current = current + serviceTimes[i]
		total = total + current
	}
	return total
}
