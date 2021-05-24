package enumerate_trees

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func GenerateAllBinaryTrees(numNodes int) []*tree.BinaryTreeNode {
	cache := make([][]*tree.BinaryTreeNode, numNodes+1)
	return generateAllBinaryTrees(numNodes, cache)
}

func generateAllBinaryTrees(numNodes int, cache [][]*tree.BinaryTreeNode) []*tree.BinaryTreeNode {
	if cache[numNodes] != nil {
		return cache[numNodes]
	}

	if numNodes == 0 {
		cache[0] = []*tree.BinaryTreeNode{nil} // only the empty tree has height 0
		return cache[0]
	}

	if numNodes == 1 {
		cache[1] = []*tree.BinaryTreeNode{
			{},
		}
		return cache[1]
	}

	result := make([]*tree.BinaryTreeNode, 0)

	for i := 0; i < numNodes; i++ {
		leftTrees := generateAllBinaryTrees(i, cache)
		rightTrees := generateAllBinaryTrees(numNodes-i-1, cache)

		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				result = append(result, &tree.BinaryTreeNode{
					Left:  lt,
					Right: rt,
				})
			}
		}
	}
	cache[numNodes] = result

	return result
}
