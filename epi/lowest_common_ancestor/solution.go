package lowest_common_ancestor

import (
	"fmt"

	"github.com/stefantds/go-epi-judge/tree"
)

func LCA(tree, node0, node1 *tree.BinaryTreeNode) *tree.BinaryTreeNode {
	lca, r, l := lcaHelper(tree, node0, node1)
	fmt.Sprintln(r, l)
	return lca
}

func lcaHelper(current, node0, node1 *tree.BinaryTreeNode) (*tree.BinaryTreeNode, bool, bool) {
	if current == nil {
		return nil, false, false
	}

	lcaL, foundN0L, foundN1L := lcaHelper(current.Left, node0, node1)
	if lcaL != nil {
		return lcaL, true, true
	}

	lcaR, foundN0R, foundN1R := lcaHelper(current.Right, node0, node1)
	if lcaR != nil {
		return lcaR, true, true
	}

	n0Found := foundN0L || foundN0R || current == node0
	n1Found := foundN1L || foundN1R || current == node1

	if n0Found && n1Found {
		return current, true, true
	}

	return nil, n0Found, n1Found
}
