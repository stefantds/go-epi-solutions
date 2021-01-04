package path_sum

import (
	"github.com/stefantds/go-epi-judge/tree"
)

func HasPathSum(t *tree.BinaryTreeNode, remainingWeight int) bool {
	if t == nil {
		return false
	}

	if t.Left == nil && t.Right == nil && t.Data == remainingWeight {
		return true
	}

	return HasPathSum(t.Left, remainingWeight-t.Data) || HasPathSum(t.Right, remainingWeight-t.Data)
}
