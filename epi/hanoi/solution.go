package hanoi

import "github.com/stefantds/go-epi-judge/stack"

// only 3 pegs are supported, don't change this
const NumPegs = 3

func ComputeTowerHanoi(numRings int) [][]int {
	pegs := make([]stack.Stack, NumPegs)

	for i := 0; i < NumPegs; i++ {
		pegs[i] = make(stack.Stack, 0, numRings)
	}

	for i := numRings; i >= 1; i-- {
		pegs[0] = pegs[0].Push(i)
	}

	result := make([][]int, 0)

	computeTowerHanoiSteps(numRings, pegs, 0, 1, 2, &result)
	return result
}

func computeTowerHanoiSteps(numRings int, pegs []stack.Stack, fromPeg int, toPeg int, auxPeg int, result *[][]int) {
	if numRings > 0 {
		computeTowerHanoiSteps(numRings-1, pegs, fromPeg, auxPeg, toPeg, result)

		var top interface{}
		pegs[fromPeg], top = pegs[fromPeg].Pop()
		pegs[toPeg] = pegs[toPeg].Push(top)

		*result = append(*result, []int{fromPeg, toPeg})

		computeTowerHanoiSteps(numRings-1, pegs, auxPeg, toPeg, fromPeg, result)
	}
}
