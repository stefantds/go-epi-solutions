package search_frequent_items

// ResetIterator has a metgod Iterator that returns a new channel
// containing the same stream of values every time.
// It effectively allows to reset the channel and read again from it multiple times.
type ResetIterator interface {
	Iterator() <-chan string
}

func SearchFrequentItems(k int, stream ResetIterator) []string {
	candidates := make(map[string]int)
	itemsCount := 0
	for s := range stream.Iterator() {
		itemsCount++
		candidates[s]++
		if len(candidates) == k {
			keysToDelete := make([]string, 0)
			// remove one from every item
			for k, v := range candidates {
				if v == 1 {
					keysToDelete = append(keysToDelete, k)
				} else {
					candidates[k]--
				}
			}
			for _, k := range keysToDelete {
				delete(candidates, k)
			}
		}
	}

	// reset the count for the second pass
	for k := range candidates {
		candidates[k] = 0
	}

	for s := range stream.Iterator() {
		if _, ok := candidates[s]; ok {
			candidates[s]++
		}
	}

	threshold := float64(itemsCount) / float64(k)
	result := make([]string, 0)
	for k, v := range candidates {
		if float64(v) > threshold {
			result = append(result, k)
		}
	}

	return result
}
