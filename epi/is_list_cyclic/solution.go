package is_list_cyclic

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func HasCycle(head *list.Node) *list.Node {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil {
		return nil // no cycle
	}

	// there is a cycle and both slow and fast point to a node in the cycle
	cycleLength := 1

	// go through the cycle until reach slow again while counting the length
	for fast.Next != slow {
		cycleLength++
		fast = fast.Next
	}

	// advance one cursor by the length of the cycle
	advancedCursor := head
	for cycleLength > 0 {
		advancedCursor = advancedCursor.Next
		cycleLength--
	}

	// advance both cursors until they meet
	cursor := head
	for cursor != advancedCursor {
		cursor, advancedCursor = cursor.Next, advancedCursor.Next
	}

	return cursor
}
