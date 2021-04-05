package k_largest_in_heap

import "container/heap"

func KLargestInBinaryHeap(a []int, k int) []int {
	if len(a) < k {
		panic("not enough elements")
	}
	if len(a) == 0 {
		return []int{}
	}
	candidates := make(MaxHeap, 0)
	heap.Init(&candidates)
	result := make([]int, 0)

	heap.Push(&candidates, HeapItem{Value: a[0], Index: 0})

	for len(result) < k {
		// the current element is the max, add it to the result
		currentItem := heap.Pop(&candidates).(HeapItem)
		result = append(result, currentItem.Value)

		// add the children of the current item to the candidates
		if leftIdx := 2*currentItem.Index + 1; leftIdx < len(a) {
			heap.Push(&candidates, HeapItem{Value: a[leftIdx], Index: leftIdx})
		}
		if rightIdx := 2*currentItem.Index + 2; rightIdx < len(a) {
			heap.Push(&candidates, HeapItem{Value: a[rightIdx], Index: rightIdx})
		}
	}

	return result
}

type HeapItem struct {
	Index int
	Value int
}

type MaxHeap []HeapItem

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Value > h[j].Value // we need to inverse the direction as we want a max heap
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(HeapItem)) }

func (h *MaxHeap) Pop() interface{} {
	var item interface{}
	n := len(*h)
	*h, item = (*h)[0:n-1], (*h)[n-1]
	return item
}
