package is_tree_a_bst

import (
	"math"

	"github.com/stefantds/go-epi-judge/tree"
)

func IsBinaryTreeBST(tree *tree.BinaryTreeNode) bool {
	return areKeysInRange(tree, math.MinInt32, math.MaxInt32)
}

func areKeysInRange(tree *tree.BinaryTreeNode, min, max int) bool {
	if tree == nil {
		return true
	}

	if tree.Data > max || tree.Data < min {
		return false
	}

	return areKeysInRange(tree.Left, min, tree.Data) &&
		areKeysInRange(tree.Right, tree.Data, max)
}
