package knapsack

import utils "github.com/stefantds/go-epi-judge/test_utils"

// KnapsackVariant2 solves the problem using only O(k) extra space,
// where k is the number of values between 0 and capacity that can be achieved.
func KnapsackVariant2(items []Item, capacity int) int {
	// feasible is a dictionary holding all the values that have been previously
	// found as possible. The key of the map holds the value of the items and the
	// value of the map holds its best weight.
	feasible := make(map[int]int, capacity)

	// include the trivial 0 solution as a start
	feasible[0] = 0
	maxValue := 0

	for _, item := range items {
		newlyFeasible := make(map[int]int, capacity)
		for prevWeight, prevValue := range feasible {
			weightWithItem := prevWeight + item.Weight
			valueWithItem := prevValue + item.Value
			if weightWithItem <= capacity {
				// keep the value only if not found before or better than the current one
				if oldVal, exists := feasible[weightWithItem]; !exists || oldVal < valueWithItem {
					newlyFeasible[weightWithItem] = valueWithItem
					maxValue = utils.Max(maxValue, valueWithItem)
				}
			}
		}

		merge(&feasible, &newlyFeasible)
	}

	return maxValue
}

// merge copies all elements from m2 to m1
func merge(m1, m2 *map[int]int) *map[int]int {
	for k, v := range *m2 {
		(*m1)[k] = v
	}

	return m1
}
