package lowest_common_ancestor_close_ancestor

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func LCA(node0 *tree.BinaryTree, node1 *tree.BinaryTree) *tree.BinaryTree {
	seenNodes := make(map[*tree.BinaryTree]bool)
	for node0 != nil || node1 != nil {
		if node0 != nil {
			if ok := seenNodes[node0]; ok {
				return node0
			}
			seenNodes[node0] = true
			node0 = node0.Parent
		}

		if node1 != nil {
			if ok := seenNodes[node1]; ok {
				return node1
			}
			seenNodes[node1] = true
			node1 = node1.Parent
		}
	}

	return nil // if we get here no LCA exists
}
