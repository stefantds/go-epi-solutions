package k_largest_values_in_bst

import (
	"github.com/stefantds/go-epi-judge/tree"
)

func FindKLargestInBst(tree *tree.BSTNode, k int) []int {
	result := make([]int, 0, k)
	findKLargestInBstHelper(tree, k, &result)

	return result
}

func findKLargestInBstHelper(tree *tree.BSTNode, k int, currentList *[]int) {
	if tree == nil || len(*currentList) == k {
		return
	}

	findKLargestInBstHelper(tree.Right, k, currentList)

	if len(*currentList) < k {
		*currentList = append(*currentList, tree.Data)
	}

	findKLargestInBstHelper(tree.Left, k, currentList)
}
