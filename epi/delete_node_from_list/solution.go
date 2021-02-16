package delete_node_from_list

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func DeletionFromList(nodeToDelete *list.Node) {
	nodeToDelete.Data = nodeToDelete.Next.Data
	nodeToDelete.Next = nodeToDelete.Next.Next // nodeToDelete is guaranteed not to be the tail node
}
