package tree_inorder

import (
	"github.com/stefantds/go-epi-judge/data_structures/stack"
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

type NodeWithState struct {
	node     *tree.BinaryTreeNode
	expanded bool
}

func InorderTraversal(t *tree.BinaryTreeNode) []int {
	if t == nil {
		return []int{}
	}

	s := stack.Stack{}
	s = s.Push(NodeWithState{
		node:     t,
		expanded: false,
	})

	result := make([]int, 0)

	for !s.IsEmpty() {
		var node interface{}
		s, node = s.Pop()
		n := node.(NodeWithState)
		if n.expanded {
			// the left subtree was already processed, we can just add the
			// value to the result
			result = append(result, n.node.Data)
		} else {
			if n.node.Right != nil {
				s = s.Push(NodeWithState{
					node:     n.node.Right,
					expanded: false,
				})
			}
			s = s.Push(NodeWithState{
				node:     n.node,
				expanded: true,
			})
			if n.node.Left != nil {
				s = s.Push(NodeWithState{
					node:     n.node.Left,
					expanded: false,
				})
			}
		}
	}

	return result
}

func InorderTraversalIterV2(root *tree.BinaryTreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*tree.BinaryTreeNode, 0)
	iter := root
	for iter != nil || len(stack) > 0 {
		for iter != nil {
			stack = append(stack, iter)
			iter = iter.Left
		}

		iter, stack = stack[len(stack)-1], stack[:len(stack)-1]

		// visit node
		result = append(result, iter.Data)

		iter = iter.Right
	}

	return result
}
