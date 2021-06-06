package pivot_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func ListPivoting(l *list.Node, x int) *list.Node {
	lHead, eHead, gHead := new(list.Node), new(list.Node), new(list.Node)
	lIter, eIter, gIter := lHead, eHead, gHead

	for iter := l; iter != nil; iter = iter.Next {
		switch {
		case iter.Data < x:
			lIter.Next = iter
			lIter = lIter.Next
		case iter.Data == x:
			eIter.Next = iter
			eIter = eIter.Next
		case iter.Data > x:
			gIter.Next = iter
			gIter = gIter.Next
		}
	}

	gIter.Next = nil
	eIter.Next = gHead.Next
	lIter.Next = eHead.Next
	return lHead.Next
}
