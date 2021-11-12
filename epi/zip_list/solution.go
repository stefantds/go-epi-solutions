package zip_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func ZippingLinkedList(l *list.Node) *list.Node {
	l1, l2 := splitHalf(l)
	l2 = reverse(l2)
	return merge(l1, l2)
}

func splitHalf(l *list.Node) (firstHalf, secondHalf *list.Node) {
	if l == nil || l.Next == nil {
		return l, nil
	}

	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	firstHalf = l
	secondHalf = slow.Next
	slow.Next = nil
	return firstHalf, secondHalf
}

func reverse(l *list.Node) *list.Node {
	if l == nil || l.Next == nil {
		return l
	}

	var curr, next, nnext *list.Node
	curr, next, nnext = nil, l, l.Next

	for next != nil {
		next.Next = curr // inverse the link
		curr = next
		next = nnext
		if nnext != nil {
			nnext = nnext.Next
		}
	}

	return curr
}

func merge(l1, l2 *list.Node) *list.Node {
	l1Iter, l2Iter := l1, l2
	for l2Iter != nil {
		temp := l1Iter.Next
		l1Iter.Next = l2Iter
		l2Iter = l2Iter.Next
		l1Iter.Next.Next = temp
		l1Iter = l1Iter.Next.Next
	}
	return l1
}
