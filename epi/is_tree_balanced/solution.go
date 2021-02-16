package is_tree_balanced

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func IsBalanced(t *tree.BinaryTreeNode) bool {
	if t == nil {
		return true
	}

	result, _ := areChildrenBalanced(t)

	return result
}

// areChildrenBalanced returns true if the heights of the children have a difference of at most 1.
// Also returns the height of the tree rooted in tree.
func areChildrenBalanced(t *tree.BinaryTreeNode) (bool, int) {
	if t == nil {
		return true, 0
	}
	leftBalance, leftHeight := areChildrenBalanced(t.Left)
	rightBalance, rightHeight := areChildrenBalanced(t.Right)
	return leftBalance && rightBalance && utils.Abs(leftHeight-rightHeight) <= 1, 1 + utils.Max(leftHeight, rightHeight)
}
