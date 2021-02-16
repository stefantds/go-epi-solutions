package remove_duplicates_from_sorted_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func RemoveDuplicates(l *list.Node) *list.Node {
	for iter := l; iter != nil; {
		if iter.Next != nil && iter.Data == iter.Next.Data {
			iter.Next = iter.Next.Next
		} else {
			iter = iter.Next
		}
	}

	return l
}
