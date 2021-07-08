package tree_from_preorder_inorder

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func BinaryTreeFromPreorderInorder(preorder []int, inorder []int) *tree.BinaryTreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	var rootPos, v int
	for rootPos, v = range inorder {
		if v == rootVal {
			break
		}
	}

	return &tree.BinaryTreeNode{
		Data:  preorder[0],
		Left:  BinaryTreeFromPreorderInorder(preorder[1:rootPos+1], inorder[:rootPos]),
		Right: BinaryTreeFromPreorderInorder(preorder[rootPos+1:], inorder[rootPos+1:]),
	}
}
