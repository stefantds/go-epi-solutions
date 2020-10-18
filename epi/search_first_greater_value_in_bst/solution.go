package search_first_greater_value_in_bst

import (
	"github.com/stefantds/go-epi-judge/tree"
)

func FindFirstGreaterThanK(t *tree.BSTNode, k int) *tree.BSTNode {
	if t == nil {
		return nil
	}
	if t.Data > k {
		if t.Left == nil {
			return t
		} else {
			if foundLeft := FindFirstGreaterThanK(t.Left, k); foundLeft != nil {
				return foundLeft
			}
			return t
		}
	} else { // t.Data <= k
		if t.Right != nil {
			return FindFirstGreaterThanK(t.Right, k)
		}
		return nil
	}
}
