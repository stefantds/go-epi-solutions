package lowest_common_ancestor_in_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func FindLCA(t *tree.BSTNode, s *tree.BSTNode, b *tree.BSTNode) *tree.BSTNode {
	if t == nil { // only happens if no node satisfies the condition
		return nil
	}

	var lca *tree.BSTNode
	switch {
	case s == t || b == t:
		lca = t
	case t.Data > s.Data && t.Data < b.Data:
		lca = t
	case t.Data < s.Data && t.Data > b.Data:
		lca = t
	case t.Data > s.Data: // both searched nodes are left
		lca = FindLCA(t.Left, s, b)
	case t.Data < s.Data: // both searched nodes are right
		lca = FindLCA(t.Right, s, b)
	}

	return lca
}
