package bst_from_preorder

import (
	"math"

	"github.com/stefantds/go-epi-judge/tree"
)

func RebuildBSTFromPreorder(preorderSequence []int) *tree.BSTNode {
	tree, _ := rebuildBSTFromRootIdx(preorderSequence, 0, math.MinInt64, math.MaxInt64)
	return tree
}

// rebuildBSTFromRootIdx builds the BST rooted in rootIdx. It returns the BST and the index
// until where the values in the sequence were consumed (position at returned index was not consumed)
func rebuildBSTFromRootIdx(sequence []int, rootIdx, lowerBound, upperBound int) (*tree.BSTNode, int) {
	if rootIdx >= len(sequence) {
		return nil, rootIdx
	}

	rootValue := sequence[rootIdx]
	if rootValue < lowerBound || rootValue > upperBound {
		return nil, rootIdx
	}

	leftSubtree, rootIdx := rebuildBSTFromRootIdx(sequence, rootIdx+1, lowerBound, rootValue)
	rightSubtree, rootIdx := rebuildBSTFromRootIdx(sequence, rootIdx, rootValue, upperBound)
	return &tree.BSTNode{
		Data:  rootValue,
		Left:  leftSubtree,
		Right: rightSubtree,
	}, rootIdx
}
