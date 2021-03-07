package tree_right_sibling

func ConstructRightSibling(tree *BinaryTreeNodeWithNext) {
	for tree != nil {
		constructNextLevel(tree)
		tree = tree.Left
	}
}

func constructNextLevel(node *BinaryTreeNodeWithNext) {
	for node != nil && node.Left != nil {
		node.Left.Next = node.Right
		if node.Next != nil {
			node.Right.Next = node.Next.Left
		}

		node = node.Next
	}
}
