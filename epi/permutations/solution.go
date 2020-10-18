package permutations

func Permutations(a []int) [][]int {
	result := make([][]int, 0)
	permutationsHelper([]int{}, a, &result)
	return result
}

func permutationsHelper(partialPerm []int, remainingArray []int, result *[][]int) {
	if len(remainingArray) == 0 {
		newPerm := make([]int, len(partialPerm))
		copy(newPerm, partialPerm)
		*result = append(*result, newPerm)
	} else {
		for i := range remainingArray {
			// bring item on first place
			remainingArray[0], remainingArray[i] = remainingArray[i], remainingArray[0]
			permutationsHelper(
				append(partialPerm, remainingArray[0]),
				remainingArray[1:],
				result,
			)
			// put the element back
			remainingArray[0], remainingArray[i] = remainingArray[i], remainingArray[0]
		}
	}
}
