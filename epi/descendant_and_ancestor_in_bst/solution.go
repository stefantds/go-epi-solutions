package descendant_and_ancestor_in_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func PairIncludesAncestorAndDescendantOfM(possibleAncOrDesc0 *tree.BSTNode, possibleAncOrDesc1 *tree.BSTNode, middle *tree.BSTNode) bool {
	return searchDirected(possibleAncOrDesc0, []*tree.BSTNode{middle, possibleAncOrDesc1}) ||
		searchDirected(possibleAncOrDesc1, []*tree.BSTNode{middle, possibleAncOrDesc0})
}

// searchDirected starts from the start node and then searches for all the nodes
// from the orderedNodes slice in order.
func searchDirected(start *tree.BSTNode, orderedNodes []*tree.BSTNode) bool {
	if len(orderedNodes) == 0 {
		// all nodes were found
		return true
	}

	if start == nil {
		return false
	}

	if orderedNodes[0] == start {
		if len(orderedNodes) == 1 {
			// we've found the last node
			return true
		}

		return searchDirected(nextPossibleNode(start, orderedNodes[1]), orderedNodes[1:])
	}

	return searchDirected(nextPossibleNode(start, orderedNodes[0]), orderedNodes)
}

// nextPossibleNode returns either the right or the left children of the current BST node,
// depending if the target (searched) node is greater or smaller than the current value.
func nextPossibleNode(current, target *tree.BSTNode) *tree.BSTNode {
	if target.Data < current.Data {
		return current.Left
	}

	return current.Right
}
