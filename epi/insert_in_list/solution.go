package insert_in_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func InsertAfter(node *list.Node, newNode *list.Node) {
	newNode.Next = node.Next
	node.Next = newNode
}
