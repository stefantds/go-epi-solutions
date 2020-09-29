package sorted_arrays_merge

import "container/heap"

func MergeSortedArrays(sortedArrays [][]int) []int {
	nbArrays := len(sortedArrays)
	currentIndices := make([]int, nbArrays)
	pq := make(PriorityQueue, nbArrays)
	result := make([]int, 0)

	// put the first items from each array into the PQ
	for i := 0; i < nbArrays; i++ {
		pq[i] = &Item{
			Value:  sortedArrays[i][0],
			Source: i,
		}
		currentIndices[i] += 1
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.Value)
		if currentIndices[item.Source] < len(sortedArrays[item.Source]) {
			newItem := Item{
				Value:  sortedArrays[item.Source][currentIndices[item.Source]],
				Source: item.Source,
			}
			heap.Push(&pq, &newItem)
			currentIndices[item.Source] += 1
		}
	}

	return result
}
