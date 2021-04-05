package search_in_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func SearchBST(t *tree.BSTNode, key int) *tree.BSTNode {
	if t == nil {
		return nil
	}

	if t.Data == key {
		return t
	}

	if key < t.Data {
		return SearchBST(t.Left, key)
	} else {
		return SearchBST(t.Right, key)
	}
}
