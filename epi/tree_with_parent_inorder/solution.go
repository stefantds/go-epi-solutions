package tree_with_parent_inorder

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func InorderTraversalWithParent(t *tree.BinaryTree) []int {
	if t == nil {
		return []int{}
	}
	result := make([]int, 0)
	iter := t
	var prev *tree.BinaryTree

	for iter != nil {
		var next *tree.BinaryTree
		if iter.Parent == prev { // came down from the parent
			if iter.Left != nil {
				next = iter.Left
			} else {
				result = append(result, iter.Data)
				next = rightOrParent(iter)
			}
		} else { // came up from child
			if iter.Left == prev { // just finished a left tree
				result = append(result, iter.Data)
				next = rightOrParent(iter)
			} else { // just finished a right tree
				next = iter.Parent
			}
		}

		prev = iter
		iter = next
	}

	return result
}

// rightOrParent returns the right node if it's not nil.
// Otherwise it returns the parent.
func rightOrParent(t *tree.BinaryTree) *tree.BinaryTree {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		return t.Right
	}

	return t.Parent
}
