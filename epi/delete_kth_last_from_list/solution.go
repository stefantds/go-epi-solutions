package delete_kth_last_from_list

import (
	"fmt"

	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func RemoveKthLast(l *list.Node, k int) *list.Node {
	dummyHead := &list.Node{Next: l}
	firstIter, secondIter := dummyHead, dummyHead
	for i := 0; i < k; i++ {
		if firstIter.Next == nil {
			panic(fmt.Errorf("the list is too short, expected at least %v elements", k))
		}
		firstIter = firstIter.Next
	}

	for firstIter.Next != nil {
		firstIter = firstIter.Next
		secondIter = secondIter.Next
	}

	// secondIter.Next is the node to be deleted
	secondIter.Next = secondIter.Next.Next

	return dummyHead.Next
}
