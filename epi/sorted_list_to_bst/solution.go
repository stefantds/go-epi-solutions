package sorted_list_to_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func BuildBSTFromSortedList(l *list.DoublyLinkedNode, length int) *list.DoublyLinkedNode {
	head := l
	result, _ := buildBST(0, length, head)
	return result
}

func buildBST(start, end int, currentHead *list.DoublyLinkedNode) (result, next *list.DoublyLinkedNode) {
	if start >= end {
		return nil, currentHead
	}

	mid := start + (end-start)/2

	// build a bst from the first half of the list
	left, head := buildBST(start, mid, currentHead)

	// curr is the new root
	curr := head

	// set the left child
	curr.Prev = left

	// build a bst from the second half of the list (without mid)
	nextCurr, nextHead := buildBST(mid+1, end, head.Next)

	// set the right child
	curr.Next = nextCurr

	return curr, nextHead
}
