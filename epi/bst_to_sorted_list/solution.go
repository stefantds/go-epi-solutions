package bst_to_sorted_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func BstToDoublyLinkedList(t *tree.BSTNode) *tree.BSTNode {
	head, _ := bstToList(t)
	return head
}

func bstToList(t *tree.BSTNode) (start, end *tree.BSTNode) {
	if t == nil {
		return t, t
	}

	leftStart, leftEnd := bstToList(t.Left)
	rightStart, rightEnd := bstToList(t.Right)

	t.Left = leftEnd
	if leftEnd != nil {
		leftEnd.Right = t
	}

	t.Right = rightStart
	if rightStart != nil {
		rightStart.Left = t
	}

	if leftStart == nil {
		leftStart = t
	}

	if rightEnd == nil {
		rightEnd = t
	}

	return leftStart, rightEnd
}
