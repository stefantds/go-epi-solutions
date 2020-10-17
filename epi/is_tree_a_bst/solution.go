package is_tree_a_bst

import (
	"math"

	"github.com/stefantds/go-epi-judge/tree"
)

func IsBinaryTreeBST(t *tree.BinaryTreeNode) bool {
	return areKeysInRange(t, math.MinInt32, math.MaxInt32)
}

func areKeysInRange(t *tree.BinaryTreeNode, min, max int) bool {
	if t == nil {
		return true
	}

	if t.Data > max || t.Data < min {
		return false
	}

	return areKeysInRange(t.Left, min, t.Data) &&
		areKeysInRange(t.Right, t.Data, max)
}
