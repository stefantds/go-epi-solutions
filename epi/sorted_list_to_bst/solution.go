package sorted_list_to_bst

import (
	"github.com/stefantds/go-epi-judge/list"
)

func BuildBSTFromSortedList(l *list.DoublyLinkedNode, length int) *list.DoublyLinkedNode {
	head := l
	result, _ := buildSortedListHelper(0, length, head)
	return result
}

func buildSortedListHelper(start, end int, currentHead *list.DoublyLinkedNode) (result, next *list.DoublyLinkedNode) {
	if start >= end {
		return nil, currentHead
	}

	mid := start + (end-start)/2
	left, head := buildSortedListHelper(start, mid, currentHead)

	curr := head
	head.Prev = left
	head = head.Next

	nextCurr, nextHead := buildSortedListHelper(mid+1, end, head)
	curr.Next = nextCurr

	return curr, nextHead
}
