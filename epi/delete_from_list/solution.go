package delete_from_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func DeleteList(aNode *list.Node) {
	aNode.Next = aNode.Next.Next
}
