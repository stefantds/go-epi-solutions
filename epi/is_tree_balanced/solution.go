package is_tree_balanced

import (
	"github.com/stefantds/go-epi-judge/tree"
	"github.com/stefantds/go-epi-judge/utils"
)

func IsBalanced(tree *tree.BinaryTreeNode) bool {
	if tree == nil {
		return true
	}

	result, _ := areChildrenBalanced(tree)

	return result
}

// areChildrenBalanced returns true if the heights of the children have a difference of at most 1.
// Also returns the height of the tree rooted in tree.
func areChildrenBalanced(tree *tree.BinaryTreeNode) (bool, int) {
	if tree == nil {
		return true, 0
	}
	leftBalance, leftHeight := areChildrenBalanced(tree.Left)
	rightBalance, rightHeight := areChildrenBalanced(tree.Right)
	return leftBalance && rightBalance && utils.Abs(leftHeight-rightHeight) <= 1, 1 + utils.Max(leftHeight, rightHeight)
}
