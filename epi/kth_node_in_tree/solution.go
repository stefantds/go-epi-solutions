package kth_node_in_tree

func FindKthNodeBinaryTree(tree *BinaryTreeNode, k int) *BinaryTreeNode {
	if k > tree.Size {
		panic("invalid input: k > tree.Size")
	}
	current := tree
	for current != nil {
		leftSize := 0
		if current.Left != nil {
			leftSize = current.Left.Size
		}
		if leftSize == k-1 {
			return current
		} else if leftSize < k-1 {
			k -= (leftSize + 1)
			current = current.Right
		} else /* the k-th node is in the left side */ {
			current = current.Left
		}
	}

	return nil
}
