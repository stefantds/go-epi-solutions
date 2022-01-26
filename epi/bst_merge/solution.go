package bst_merge

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
	"github.com/stefantds/go-epi-judge/epi/bst_to_sorted_list"
)

func MergeTwoBsts(a *tree.BSTNode, b *tree.BSTNode) *tree.BSTNode {
	a = bst_to_sorted_list.BstToDoublyLinkedList(a)
	b = bst_to_sorted_list.BstToDoublyLinkedList(b)
	m := mergeTwoSortedLists(a, b)
	result, _ := buildBSTFromDoublyLinkedList(0, countLength(m), m)
	return result
}

func countLength(l *tree.BSTNode) int {
	count := 0
	for cur := l; cur != nil; cur = cur.Right {
		count++
	}
	return count
}

func mergeTwoSortedLists(a *tree.BSTNode, b *tree.BSTNode) *tree.BSTNode {
	dummyHead := &tree.BSTNode{}
	current := dummyHead
	curA, curB := a, b

	for curA != nil && curB != nil {
		if curA.Data < curB.Data {
			current.Right = curA
			current = current.Right
			curA = curA.Right
		} else {
			current.Right = curB
			current = current.Right
			curB = curB.Right
		}
	}

	for curA != nil {
		current.Right = curA
		current = current.Right
		curA = curA.Right
	}

	for curB != nil {
		current.Right = curB
		current = current.Right
		curB = curB.Right
	}

	return dummyHead.Right
}

func buildBSTFromDoublyLinkedList(start, end int, currentHead *tree.BSTNode) (result, next *tree.BSTNode) {
	if start >= end {
		return nil, currentHead
	}

	mid := start + (end-start)/2

	// build a bst from the first half of the list
	left, head := buildBSTFromDoublyLinkedList(start, mid, currentHead)

	// set the left child
	head.Left = left

	// build a bst from the second half of the list (without mid)
	nextCurr, nextHead := buildBSTFromDoublyLinkedList(mid+1, end, head.Right)

	// set the right child
	head.Right = nextCurr

	return head, nextHead
}
