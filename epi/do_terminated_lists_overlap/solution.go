package do_terminated_lists_overlap

import (
	"github.com/stefantds/go-epi-judge/list"
	"github.com/stefantds/go-epi-judge/utils"
)

func OverlappingNoCycleLists(l0 *list.Node, l1 *list.Node) *list.Node {
	l0Length := getListSize(l0)
	l1Length := getListSize(l1)

	var longer, shorter *list.Node
	if l0Length-l1Length > 0 {
		longer, shorter = l0, l1
	} else {
		longer, shorter = l1, l0
	}

	for i := 0; i < utils.Abs(l0Length-l1Length); i++ {
		longer = longer.Next
	}

	for longer != nil && longer != shorter {
		longer, shorter = longer.Next, shorter.Next
	}

	return longer
}

func getListSize(l *list.Node) int {
	size := 0
	cursor := l
	for cursor != nil {
		size++
		cursor = cursor.Next
	}

	return size
}
