package range_lookup_in_bst

import (
	"github.com/stefantds/go-epi-judge/data_structures/tree"
)

func RangeLookupInBst(t *tree.BSTNode, interval Interval) []int {
	result := make([]int, 0)
	lookup(t, interval, &result)
	return result
}

func lookup(t *tree.BSTNode, interval Interval, result *[]int) {
	if t == nil {
		return
	}
	switch {
	case t.Data < interval.Left:
		lookup(t.Right, interval, result)
	case t.Data > interval.Right:
		lookup(t.Left, interval, result)
	default:
		lookup(t.Left, interval, result)
		*result = append(*result, t.Data)
		lookup(t.Right, interval, result)
	}
}
