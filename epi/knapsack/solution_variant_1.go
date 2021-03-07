package knapsack

// KnapsackVariant1 solves the problem using only O(capacity) extra space
func KnapsackVariant1(items []Item, capacity int) int {
	// cache[i] will hold the total value and total weight
	// for capacity i
	cache := make([]Item, capacity+1)

	// initialize the cache with the values for the first item,
	// i.e. best choice subject to capacity if the first item is the only choice.
	for i := 0; i <= capacity; i++ {
		if items[0].Weight <= i {
			cache[i] = Item{
				Value:  items[0].Value,
				Weight: items[0].Weight,
			}
		} else {
			cache[i] = Item{
				Value:  0,
				Weight: 0,
			}
		}
	}

	for i := 1; i < len(items); i++ {
		newCache := make([]Item, capacity+1)

		for currentCap := 0; currentCap <= capacity; currentCap++ {
			withCurrentItem := add(cache[currentCap], items[i])
			if withCurrentItem.Weight <= capacity {
				if less(newCache[withCurrentItem.Weight], withCurrentItem) {
					// we've found a better option for withCurrentItem.Weight capacity
					newCache[withCurrentItem.Weight] = Item{
						Value:  cache[currentCap].Value + items[i].Value,
						Weight: cache[currentCap].Weight + items[i].Weight,
					}
				}
			}

			// check if previously we've found a better value for the same capacity
			if less(newCache[currentCap], cache[currentCap]) {
				newCache[currentCap] = Item{
					Value:  cache[currentCap].Value,
					Weight: cache[currentCap].Weight,
				}
			}

			// check if we've found a better value for a smaller capacity
			if currentCap > 0 && less(newCache[currentCap], newCache[currentCap-1]) {
				newCache[currentCap] = Item{
					Value:  newCache[currentCap-1].Value,
					Weight: newCache[currentCap-1].Weight,
				}
			}
		}

		cache = newCache
	}

	return cache[capacity].Value
}

// less compares the absolute value of two items: it compares the items by
// value first. In case of equal value, it compares the weight (more weight means less value).
func less(i1, i2 Item) bool {
	if i1.Value < i2.Value {
		return true
	}
	if i1.Value > i2.Value {
		return false
	}
	return i1.Weight > i2.Weight
}

// add adds two items by creating another item
// with the value and the weight equal to the sum
// of values and weights respectively.
func add(i1, i2 Item) Item {
	return Item{
		Value:  i1.Value + i2.Value,
		Weight: i1.Weight + i2.Weight,
	}
}
