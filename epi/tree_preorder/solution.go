package tree_preorder

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func PreorderTraversal(t *tree.BinaryTreeNode) []int {
	if t == nil {
		return []int{}
	}
	result := make([]int, 0)
	preorder(t, &result)
	return result
}

func preorder(t *tree.BinaryTreeNode, result *[]int) {
	if t == nil {
		return
	}
	(*result) = append((*result), t.Data)
	preorder(t.Left, result)
	preorder(t.Right, result)
}

func PreorderTraversalNoRecursion(t *tree.BinaryTreeNode) []int {
	if t == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := make([]*tree.BinaryTreeNode, 0)
	stack = append(stack, t)
	for len(stack) > 0 {
		var top *tree.BinaryTreeNode
		stack, top = stack[:len(stack)-1], stack[len(stack)-1]
		result = append(result, top.Data)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return result
}
