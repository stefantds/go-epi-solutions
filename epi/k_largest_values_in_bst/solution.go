package k_largest_values_in_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func FindKLargestInBst(t *tree.BSTNode, k int) []int {
	result := make([]int, 0, k)
	findKLargestInBstHelper(t, k, &result)

	return result
}

func findKLargestInBstHelper(t *tree.BSTNode, k int, currentList *[]int) {
	if t == nil || len(*currentList) == k {
		return
	}

	findKLargestInBstHelper(t.Right, k, currentList)

	if len(*currentList) < k {
		*currentList = append(*currentList, t.Data)
	}

	findKLargestInBstHelper(t.Left, k, currentList)
}
