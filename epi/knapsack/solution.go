package knapsack

import (
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func OptimumSubjectToCapacity(items []Item, capacity int) int {
	cache := make([][]int, capacity+1)
	for i := 0; i < capacity+1; i++ {
		cache[i] = make([]int, len(items))
		for j := 0; j < len(items); j++ {
			cache[i][j] = -1
		}
	}
	return maximiseValueHelper(items, capacity, 0, cache)
}

func maximiseValueHelper(items []Item, capacity int, itemIdx int, cache [][]int) int {
	if itemIdx == len(items) {
		return 0
	}
	if cache[capacity][itemIdx] != -1 {
		return cache[capacity][itemIdx]
	}
	item := items[itemIdx]
	valueWithoutItem := maximiseValueHelper(items, capacity, itemIdx+1, cache)
	if capacity < item.Weight {
		cache[capacity][itemIdx] = valueWithoutItem
		return valueWithoutItem
	}

	valueWithItem := item.Value + maximiseValueHelper(items, capacity-item.Weight, itemIdx+1, cache)
	result := utils.Max(valueWithItem, valueWithoutItem)
	cache[capacity][itemIdx] = result
	return result
}
