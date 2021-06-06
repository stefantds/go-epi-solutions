package successor_in_tree

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func FindSuccessor(node *tree.BinaryTree) *tree.BinaryTree {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return firstInorder(node.Right)
	}

	return firstParentOfLeft(node)
}

func firstInorder(node *tree.BinaryTree) *tree.BinaryTree {
	if node.Left == nil {
		return node
	}
	return firstInorder(node.Left)
}

func firstParentOfLeft(node *tree.BinaryTree) *tree.BinaryTree {
	if node.Parent == nil {
		return nil
	}
	if node.Parent.Left == node {
		return node.Parent
	}
	return firstParentOfLeft(node.Parent)
}
