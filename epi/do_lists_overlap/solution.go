package do_lists_overlap

import (
	"github.com/stefantds/go-epi-judge/epi/do_terminated_lists_overlap"
	"github.com/stefantds/go-epi-judge/epi/is_list_cyclic"
	"github.com/stefantds/go-epi-judge/list"
)

func OverlappingLists(l0 *list.Node, l1 *list.Node) *list.Node {
	root0 := is_list_cyclic.HasCycle(l0)
	root1 := is_list_cyclic.HasCycle(l1)

	if root0 == nil && root1 == nil {
		return do_terminated_lists_overlap.OverlappingNoCycleLists(l0, l1)
	}

	if root0 == nil || root1 == nil {
		// only one list has a cycle -> they can't overlap
		return nil
	}

	cursor := root0.Next
	for cursor != root0 && cursor != root1 {
		cursor = cursor.Next
	}

	if cursor == root1 {
		return root1 // the cycles are the same cycle
	}
	return nil // the cycles don't intersect
}
