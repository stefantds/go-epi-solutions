package stack_with_max

import "math"

type Solution interface {
	Push(x int)
	Pop() int
	Max() int
	Empty() bool
}

type StackWithMax struct {
	values []int
	maxes  []int
}

func NewStackWithMax() Solution {
	return &StackWithMax{
		values: make([]int, 0),
		maxes:  make([]int, 0),
	}
}

func (q *StackWithMax) Push(x int) {
	q.values = append(q.values, x)
	if x >= q.Max() {
		q.maxes = append(q.maxes, x)
	}
}

func (q *StackWithMax) Pop() int {
	n := len(q.values)
	var v int
	q.values, v = q.values[:n-1], q.values[n-1]
	if v == q.Max() {
		q.maxes = q.maxes[:len(q.maxes)-1]
	}
	return v
}

func (q *StackWithMax) Max() int {
	if len(q.maxes) == 0 {
		return math.MinInt64
	}

	return q.maxes[len(q.maxes)-1]
}

func (q *StackWithMax) Empty() bool {
	return len(q.values) == 0
}
