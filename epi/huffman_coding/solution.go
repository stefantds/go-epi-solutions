package huffman_coding

import (
	"container/heap"
)

type hufTree struct {
	freq        float64
	left, right *hufTree
}

// minHeapHufTree is a min heap that contains *HufTree elements
// the comparison is done on the total frequency of a tree.
type minHeapHufTree []*hufTree

func (h minHeapHufTree) Len() int { return len(h) }

func (h minHeapHufTree) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h minHeapHufTree) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeapHufTree) Push(x interface{}) { *h = append(*h, x.(*hufTree)) }

func (h *minHeapHufTree) Pop() interface{} {
	var item interface{}
	n := len(*h)
	*h, item = (*h)[0:n-1], (*h)[n-1]
	return item
}

func HuffmanEncoding(symbols []CharWithFrequency) float64 {
	// build the huffman tree
	huffmanTree := buildHuffmanTree(symbols)

	// assign codes to all nodes in the tree
	// and return the aggregated code length
	return assignCodes(huffmanTree, 0)
}

func buildHuffmanTree(symbols []CharWithFrequency) *hufTree {
	if len(symbols) == 0 {
		return nil
	}

	h := make(minHeapHufTree, 0)
	heap.Init(&h)
	for _, s := range symbols {
		heap.Push(&h, &hufTree{
			freq: s.Freq,
		})
	}

	for len(h) > 1 {
		left := heap.Pop(&h).(*hufTree)
		right := heap.Pop(&h).(*hufTree)
		heap.Push(&h, &hufTree{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		})
	}
	return heap.Pop(&h).(*hufTree)
}

func assignCodes(t *hufTree, codeLength int) float64 {
	if t == nil {
		return 0
	}

	// if it's a leaf return the weighted code length
	if t.left == nil && t.right == nil {
		return t.freq * float64(codeLength) / 100
	}

	return assignCodes(t.left, codeLength+1) +
		assignCodes(t.right, codeLength+1)
}
