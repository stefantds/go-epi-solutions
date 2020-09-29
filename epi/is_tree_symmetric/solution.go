package is_tree_symmetric

import (
	"github.com/stefantds/go-epi-judge/tree"
)

func IsSymmetric(tree *tree.BinaryTreeNode) bool {
	if tree == nil {
		return true
	}

	return checkSymmetric(tree.Left, tree.Right)
}

func checkSymmetric(subtree0 *tree.BinaryTreeNode, subtree1 *tree.BinaryTreeNode) bool {
	if subtree0 == nil && subtree1 == nil {
		return true
	}

	if subtree0 != nil && subtree1 != nil {
		return subtree0.Data.(int) == subtree1.Data.(int) &&
			checkSymmetric(subtree0.Left, subtree1.Right) &&
			checkSymmetric(subtree0.Right, subtree1.Left)
	}

	return false
}
