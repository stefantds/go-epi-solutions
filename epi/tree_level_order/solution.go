package tree_level_order

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func BinaryTreeDepthOrder(t *tree.BinaryTreeNode) [][]int {
	result := make([][]int, 0)
	if t == nil {
		return result
	}

	currentLevelNodes := make([]*tree.BinaryTreeNode, 0)
	currentLevelNodes = append(currentLevelNodes, t)

	nextLevelNodes := make([]*tree.BinaryTreeNode, 0)

	for len(currentLevelNodes) != 0 {
		currentLevelValues := make([]int, 0)
		for len(currentLevelNodes) != 0 {
			var e *tree.BinaryTreeNode
			e, currentLevelNodes = currentLevelNodes[0], currentLevelNodes[1:]
			currentLevelValues = append(currentLevelValues, e.Data)

			if e.Left != nil {
				nextLevelNodes = append(nextLevelNodes, e.Left)
			}
			if e.Right != nil {
				nextLevelNodes = append(nextLevelNodes, e.Right)
			}
		}

		result = append(result, currentLevelValues)
		currentLevelNodes = nextLevelNodes
		nextLevelNodes = make([]*tree.BinaryTreeNode, 0)
	}

	return result
}
