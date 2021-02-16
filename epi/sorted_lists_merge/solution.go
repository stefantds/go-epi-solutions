package sorted_lists_merge

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func MergeTwoSortedLists(l1 *list.Node, l2 *list.Node) *list.Node {
	dummyHead := &list.Node{}
	current := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Data <= l2.Data {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return dummyHead.Next
}
