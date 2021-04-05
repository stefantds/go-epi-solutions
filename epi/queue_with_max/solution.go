package queue_with_max

type Solution interface {
	Enqueue(x int)
	Dequeue() int
	Max() int
}

type QueueWithMax struct {
	values []int
	max    []int
}

func NewQueueWithMax() Solution {
	return &QueueWithMax{
		values: make([]int, 0),
		max:    make([]int, 0),
	}
}

func (q *QueueWithMax) Enqueue(x int) {
	q.values = append(q.values, x)

	// we can remove the elements in the max queue that are smaller than x
	for len(q.max) > 0 && q.max[len(q.max)-1] < x {
		q.max = q.max[:len(q.max)-1]
	}
	q.max = append(q.max, x)
}

func (q *QueueWithMax) Dequeue() int {
	if len(q.values) == 0 {
		panic("dequeue called on an empty queue")
	}
	var value int
	value, q.values = q.values[0], q.values[1:]
	if value == q.max[0] {
		q.max = q.max[1:] // remove the element from max queue
	}
	return value
}

func (q *QueueWithMax) Max() int {
	if len(q.max) == 0 {
		panic("max called on an empty queue")
	}
	return q.max[0]
}
