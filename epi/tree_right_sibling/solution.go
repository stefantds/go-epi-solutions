package tree_right_sibling

func ConstructRightSibling(tree *BinaryTreeNodeWithNext) {
	for tree != nil {
		linkNextLevel(tree)
		tree = tree.Left
	}
}

func linkNextLevel(node *BinaryTreeNodeWithNext) {
	for node != nil && node.Left != nil {
		node.Left.Next = node.Right
		if node.Next != nil {
			node.Right.Next = node.Next.Left
		}

		node = node.Next
	}
}
