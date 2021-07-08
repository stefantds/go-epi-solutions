package k_closest_stars

import "container/heap"

func FindClosestKStars(stars <-chan Star, k int) []Star {
	maxHeap := make(StarMaxHeap, 0)
	heap.Init(&maxHeap)

	for s := range stars {
		heap.Push(&maxHeap, s)
		if maxHeap.Len() > k {
			heap.Pop(&maxHeap)
		}
	}

	return ([]Star)(maxHeap)
}

type StarMaxHeap []Star

func (h StarMaxHeap) Len() int { return len(h) }

func (h StarMaxHeap) Less(i, j int) bool { return h[i].Distance() > h[j].Distance() }

func (h StarMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StarMaxHeap) Push(x interface{}) { *h = append(*h, x.(Star)) }

func (h *StarMaxHeap) Pop() interface{} {
	var item interface{}
	n := len(*h)
	*h, item = (*h)[0:n-1], (*h)[n-1]
	return item
}
