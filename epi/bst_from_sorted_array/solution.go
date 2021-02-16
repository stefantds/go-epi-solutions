package bst_from_sorted_array

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func BuildMinHeightBSTFromSortedArray(a []int) *tree.BSTNode {
	return buildMinHeightBSTFromSortedSubarray(a, 0, len(a))
}

func buildMinHeightBSTFromSortedSubarray(a []int, start int, end int) *tree.BSTNode {
	if start >= end {
		return nil
	}
	mid := start + ((end - start) / 2)
	return &tree.BSTNode{
		Data:  a[mid],
		Left:  buildMinHeightBSTFromSortedSubarray(a, start, mid),
		Right: buildMinHeightBSTFromSortedSubarray(a, mid+1, end),
	}
}
