package even_odd_list_merge

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func EvenOddMerge(l *list.Node) *list.Node {
	dummyHeadEven := &list.Node{}
	tailEven := dummyHeadEven
	dummyHeadOdd := &list.Node{}
	tailOdd := dummyHeadOdd

	for current, idx := l, 0; current != nil; current, idx = current.Next, idx+1 {
		if idx%2 == 0 {
			tailEven.Next = current
			tailEven = tailEven.Next
		} else {
			tailOdd.Next = current
			tailOdd = tailOdd.Next
		}
	}

	tailEven.Next = dummyHeadOdd.Next
	tailOdd.Next = nil

	return dummyHeadEven.Next
}
