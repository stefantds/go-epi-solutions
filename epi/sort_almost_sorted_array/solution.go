package sort_almost_sorted_array

import (
	"container/heap"

	"github.com/stefantds/go-epi-judge/pq"
)

func SortApproximatelySortedData(sequence <-chan int, k int) []int {
	pq := make(pq.IntPriorityQueue, 0)
	heap.Init(&pq)

	result := make([]int, 0)

	for i := 0; i <= k; i++ {
		if v, ok := <-sequence; ok {
			heap.Push(&pq, v)
		}
	}

	for pq.Len() > 0 {
		result = append(result, heap.Pop(&pq).(int))
		if v, ok := <-sequence; ok {
			heap.Push(&pq, v)
		}
	}

	return result
}
