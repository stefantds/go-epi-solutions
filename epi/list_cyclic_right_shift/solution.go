package list_cyclic_right_shift

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func CyclicallyRightShiftList(l *list.Node, k int) *list.Node {
	if l == nil {
		return nil
	}

	tail := l
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length += 1
	}

	k = k % length // avoid unnecessary full shifts

	tail.Next = l // make the list cyclical

	stepsToNewHead := length - k
	for stepsToNewHead > 0 {
		l = l.Next
		tail = tail.Next
		stepsToNewHead -= 1
	}

	tail.Next = nil
	return l
}
