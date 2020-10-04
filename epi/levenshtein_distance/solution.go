package levenshtein_distance

import (
	"github.com/stefantds/go-epi-judge/utils"
)

func LevenshteinDistance(a string, b string) int {
	cache := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		cache[i] = make([]int, len(b))
		for j := 0; j < len(b); j++ {
			cache[i][j] = -1
		}
	}
	return distanceHelper(a, len(a)-1, b, len(b)-1, cache)
}

func distanceHelper(a string, aIdx int, b string, bIdx int, cache [][]int) int {
	if aIdx < 0 {
		return bIdx + 1
	}
	if bIdx < 0 {
		return aIdx + 1
	}
	if cache[aIdx][bIdx] >= 0 {
		return cache[aIdx][bIdx]
	}
	if a[aIdx] == b[bIdx] {
		d := distanceHelper(a, aIdx-1, b, bIdx-1, cache)
		cache[aIdx][bIdx] = d
		return d
	}

	dSubstitution := distanceHelper(a, aIdx-1, b, bIdx-1, cache)
	dInsertion := distanceHelper(a, aIdx-1, b, bIdx, cache)
	dDeletion := distanceHelper(a, aIdx, b, bIdx-1, cache)
	d := 1 + utils.Min(dSubstitution, utils.Min(dInsertion, dDeletion))
	cache[aIdx][bIdx] = d
	return d
}
