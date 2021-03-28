package reverse_sublist

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func ReverseSublist(l *list.Node, start int, finish int) *list.Node {
	dummyHead := &list.Node{Next: l}
	startNode := dummyHead
	var prevStart *list.Node

	// find the start node.
	// We also need the node before start for glueing after the reverse.
	for i := 0; i < start; i++ {
		prevStart, startNode = startNode, startNode.Next
	}

	// find the finish node.
	// We also need the node after finish for glueing after the reverse.
	finishNode := dummyHead
	finishNext := finishNode.Next
	for i := 0; i < finish; i++ {
		finishNode, finishNext = finishNode.Next, finishNext.Next
	}

	// reverse order between start and finish
	reverse(startNode, finishNode)

	// rewire the sublist with the rest of the list
	if prevStart != nil {
		prevStart.Next = finishNode
	}
	startNode.Next = finishNext

	return dummyHead.Next
}

func reverse(start, finish *list.Node) {
	if start == finish {
		return
	}

	if start == nil || start.Next == nil {
		panic("unexpected nil node")
	}

	current, currentNext := start, start.Next

	for current != finish && currentNext != nil {
		currentNextNext := currentNext.Next
		currentNext.Next = current
		current, currentNext = currentNext, currentNextNext
	}
}
