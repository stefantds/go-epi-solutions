package tree_postorder

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func PostorderTraversal(t *tree.BinaryTreeNode) []int {
	if t == nil {
		return []int{}
	}
	result := make([]int, 0)
	postder(t, &result)
	return result
}

func postder(t *tree.BinaryTreeNode, result *[]int) {
	if t == nil {
		return
	}
	postder(t.Left, result)
	postder(t.Right, result)
	(*result) = append((*result), t.Data)
}

func PostorderTraversalNoRecursion(t *tree.BinaryTreeNode) []int {
	if t == nil {
		return []int{}
	}
	result := make([]int, 0)

	stack := make([]*tree.BinaryTreeNode, 0)
	visited := make(map[*tree.BinaryTreeNode]bool)
	stack = append(stack, t)

	for len(stack) > 0 {
		var current *tree.BinaryTreeNode
		stack, current = stack[:len(stack)-1], stack[len(stack)-1]
		if (current.Left == nil || visited[current.Left]) &&
			(current.Right == nil || visited[current.Right]) {
			result = append(result, current.Data)
		} else {
			stack = append(stack, current)
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
			if current.Left != nil {
				stack = append(stack, current.Left)
			}
		}
		visited[current] = true
	}

	return result
}
