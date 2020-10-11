package h_index

import (
	"sort"
)

func HIndex(citations []int) int {
	sort.Ints(citations)
	for i, c := range citations {
		if c >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}
