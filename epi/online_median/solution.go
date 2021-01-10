package online_median

import (
	"container/heap"

	"github.com/stefantds/go-epi-judge/pq"
)

func OnlineMedian(sequence chan int) []float64 {
	topHalf := minHeap{
		IntPriorityQueue: make(pq.IntPriorityQueue, 0),
	}
	heap.Init(&topHalf)

	bottomHalf := maxHeap{
		IntPriorityQueue: make(pq.IntPriorityQueue, 0),
	}
	heap.Init(&bottomHalf)

	result := make([]float64, 0)
	var currentMedian float64

	for value := range sequence {
		heap.Push(&topHalf, value)
		v := heap.Pop(&topHalf)
		heap.Push(&bottomHalf, v)

		switch {
		case bottomHalf.Len()-topHalf.Len() > 1:
			v := heap.Pop(&bottomHalf)
			heap.Push(&topHalf, v)
			currentMedian = float64(topHalf.Peek()+bottomHalf.Peek()) / 2
		case bottomHalf.Len()-topHalf.Len() == 1:
			currentMedian = float64(bottomHalf.Peek())
		default:
			currentMedian = float64(topHalf.Peek()+bottomHalf.Peek()) / 2
		}

		result = append(result, currentMedian)
	}

	return result
}

type maxHeap struct {
	pq.IntPriorityQueue
}

func (h maxHeap) Less(i, j int) bool {
	return h.IntPriorityQueue[i] > h.IntPriorityQueue[j]
}

func (h maxHeap) Peek() int { return h.IntPriorityQueue[0] }

type minHeap struct {
	pq.IntPriorityQueue
}

func (h minHeap) Peek() int { return h.IntPriorityQueue[0] }
