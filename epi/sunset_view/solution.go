package sunset_view

import "github.com/stefantds/go-epi-judge/data_structures/stack"

type building struct {
	index  int
	height int
}

func ExamineBuildingsWithSunset(sequence chan int) []int {
	s := make(stack.Stack, 0)
	i := 0
	for h := range sequence {
		for !s.IsEmpty() && s.Peek().(building).height <= h {
			s, _ = s.Pop()
		}

		s = s.Push(building{
			index:  i,
			height: h,
		})
		i++
	}

	return reverseBuidlingIndices(s)
}

func reverseBuidlingIndices(s stack.Stack) []int {
	result := make([]int, 0)
	for !s.IsEmpty() {
		var e interface{}
		s, e = s.Pop()
		result = append(result, e.(building).index)
	}

	return result
}
