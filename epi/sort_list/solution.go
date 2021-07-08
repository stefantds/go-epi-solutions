package sort_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func StableSortList(l *list.Node) *list.Node {
	if l == nil || l.Next == nil {
		return l
	}
	var preSlow *list.Node
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow, fast = slow.Next, fast.Next.Next
	}

	// split the two lists
	if preSlow != nil {
		preSlow.Next = nil
	}

	part1 := StableSortList(l)
	part2 := StableSortList(slow)

	return mergeSortedLists(part1, part2)
}

func mergeSortedLists(l1, l2 *list.Node) *list.Node {
	dummyHead := &list.Node{}
	sortedCursor := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			sortedCursor.Next = l1
			l1 = l1.Next
		} else {
			sortedCursor.Next = l2
			l2 = l2.Next
		}
		sortedCursor = sortedCursor.Next
	}

	for l1 != nil {
		sortedCursor.Next = l1
		l1 = l1.Next
		sortedCursor = sortedCursor.Next
	}

	for l2 != nil {
		sortedCursor.Next = l2
		l2 = l2.Next
		sortedCursor = sortedCursor.Next
	}

	sortedCursor.Next = nil

	return dummyHead.Next
}
