package tree_connect_leaves

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func CreateListOfLeaves(t *tree.BinaryTreeNode) []*tree.BinaryTreeNode {
	if t == nil {
		return []*tree.BinaryTreeNode{}
	}

	if t.Left == nil && t.Right == nil {
		return []*tree.BinaryTreeNode{t}
	}

	leftNodes := CreateListOfLeaves(t.Left)
	rightNodes := CreateListOfLeaves(t.Right)
	return append(leftNodes, rightNodes...)
}
