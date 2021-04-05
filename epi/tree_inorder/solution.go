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
