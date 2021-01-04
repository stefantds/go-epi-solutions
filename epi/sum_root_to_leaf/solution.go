package sum_root_to_leaf

import (
	"github.com/stefantds/go-epi-judge/tree"
)

func SumRootToLeaf(t *tree.BinaryTreeNode) int {
	return sumLeafs(0, t)
}

func sumLeafs(prefix int, node *tree.BinaryTreeNode) int {
	if node == nil {
		return 0
	}

	// add the value of the current node to the value representing the
	// path above.
	sumSoFar := prefix*2 + node.Data

	if node.Right == nil && node.Left == nil {
		return sumSoFar
	}
	return sumLeafs(sumSoFar, node.Right) + sumLeafs(sumSoFar, node.Left)
}
