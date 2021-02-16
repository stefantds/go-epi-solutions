package lowest_common_ancestor_with_parent

import (
	"errors"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func LCA(node0 *tree.BinaryTree, node1 *tree.BinaryTree) *tree.BinaryTree {
	depthN0 := getDepth(node0)
	depthN1 := getDepth(node1)

	if depthN0 > depthN1 {
		node0 = moveUp(node0, depthN0-depthN1)
	} else if depthN1 > depthN0 {
		node1 = moveUp(node1, depthN1-depthN0)
	}

	for node0.Parent != nil && node1.Parent != nil && node0 != node1 {
		node0, node1 = node0.Parent, node1.Parent
	}

	if *node0 != *node1 {
		panic(errors.New("the nodes don't have a common ancestor."))
	}

	return node0
}

func moveUp(node *tree.BinaryTree, nbLevels int) *tree.BinaryTree {
	for i := 0; i < nbLevels; i++ {
		if node.Parent == nil {
			panic(errors.New("can't move up for a root node"))
		}
		node = node.Parent
	}

	return node
}

func getDepth(node *tree.BinaryTree) int {
	depth := 0
	for node.Parent != nil {
		depth++
		node = node.Parent
	}

	return depth
}
