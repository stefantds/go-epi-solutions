package tree_from_preorder_with_null

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func ReconstructPreorder(preorder []IntOrNull) *tree.BinaryTreeNode {
	if len(preorder) == 0 {
		return nil
	}
	result, _ := rebuildPreorder(preorder, 0)
	return result
}

func rebuildPreorder(preorder []IntOrNull, start int) (result *tree.BinaryTreeNode, end int) {
	v := preorder[start]
	if !v.Valid {
		return nil, start
	}

	leftTree, ltEnd := rebuildPreorder(preorder, start+1)
	rightTree, rtEnd := rebuildPreorder(preorder, ltEnd+1)
	result = &tree.BinaryTreeNode{
		Data:  v.Value,
		Right: rightTree,
		Left:  leftTree,
	}

	return result, rtEnd
}
