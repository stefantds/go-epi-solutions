package minimum_distance_3_sorted_arrays

import (
	"math"

	"github.com/stefantds/go-epi-judge/utils"
)

func FindMinDistanceSortedArrays(sortedArrays [][]int) int {
	heads := make(map[int]int)   // holds the current head index for each array
	var currentElements *BSTNode // holds the current set of elements

	// initialize the first set of elements with
	// elements at position 0 in each array
	for i, arr := range sortedArrays {
		if len(arr) < 1 {
			panic("got an empty slice")
		}
		heads[i] = 0
		if currentElements == nil {
			currentElements = &BSTNode{
				Data: ArrayData{
					Idx: i,
					Val: arr[0],
				},
			}
		} else {
			currentElements.Insert(
				ArrayData{
					Idx: i,
					Val: arr[0],
				},
			)
		}
	}

	minDistance := math.MaxInt32
	for {
		minElement := currentElements.Min()
		maxElement := currentElements.Max()
		idxMin := minElement.Idx

		minDistance = utils.Min(
			minDistance,
			maxElement.Val-minElement.Val,
		)
		heads[idxMin] = heads[idxMin] + 1

		if heads[idxMin] >= len(sortedArrays[idxMin]) {
			return minDistance
		}
		currentElements = currentElements.Delete(minElement)
		currentElements.Insert(ArrayData{
			Idx: idxMin,
			Val: sortedArrays[idxMin][heads[idxMin]],
		})
	}
}
