package search_in_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func SearchList(l *list.Node, key int) *list.Node {
	cursor := l
	for cursor.Data != key {
		cursor = cursor.Next
		if cursor == nil {
			return nil
		}
	}

	return cursor
}
