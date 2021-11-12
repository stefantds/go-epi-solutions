package tree_exterior

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func ExteriorBinaryTree(t *tree.BinaryTreeNode) []*tree.BinaryTreeNode {
	if t == nil {
		return []*tree.BinaryTreeNode{}
	}
	leftSide := make([]*tree.BinaryTreeNode, 0)
	leaves := make([]*tree.BinaryTreeNode, 0)
	rightSide := make([]*tree.BinaryTreeNode, 0)

	leftSide = append(leftSide, t)
	if t.Left != nil {
		collectLeftSide(t.Left, true, &leftSide, &leaves)
	}
	if t.Right != nil {
		collectRightSide(t.Right, true, &rightSide, &leaves)
	}
	return append(
		append(leftSide, leaves...),
		rightSide...,
	)
}

func collectLeftSide(
	root *tree.BinaryTreeNode,
	isLeftmost bool,
	sideCollector *[]*tree.BinaryTreeNode,
	leavesCollector *[]*tree.BinaryTreeNode,
) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*leavesCollector = append(*leavesCollector, root)
		return
	}
	if isLeftmost {
		*sideCollector = append(*sideCollector, root)
	}
	collectLeftSide(root.Left, isLeftmost && root.Left != nil, sideCollector, leavesCollector)
	collectLeftSide(root.Right, isLeftmost && root.Left == nil, sideCollector, leavesCollector)
}

func collectRightSide(
	root *tree.BinaryTreeNode,
	isRightmost bool,
	sideCollector *[]*tree.BinaryTreeNode,
	leavesCollector *[]*tree.BinaryTreeNode,
) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*leavesCollector = append(*leavesCollector, root)
		return
	}
	collectRightSide(root.Left, isRightmost && root.Right == nil, sideCollector, leavesCollector)
	collectRightSide(root.Right, isRightmost && root.Right != nil, sideCollector, leavesCollector)
	if isRightmost {
		*sideCollector = append(*sideCollector, root)
	}
}

func isLeaf(n *tree.BinaryTreeNode) bool {
	return n != nil && n.Left == nil && n.Right == nil
}
