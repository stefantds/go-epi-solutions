package circular_queue

import "errors"

type Queue struct {
	values     []int
	itemsCount int
	tail       int
}

func NewQueue(capacity int) Queue {
	return Queue{
		values:     make([]int, capacity),
		itemsCount: 0,
		tail:       0,
	}
}

func (q *Queue) Enqueue(x int) {
	if q.itemsCount == len(q.values) {
		q.expand()
	}
	nextPos := (q.tail + q.itemsCount) % len(q.values)
	q.values[nextPos] = x
	q.itemsCount++
}

func (q *Queue) Dequeue() int {
	if q.itemsCount == 0 {
		panic(errors.New("dequeue from empty queue"))
	}
	val := q.values[q.tail]
	q.tail = (q.tail + 1) % len(q.values)
	q.itemsCount--
	return val
}

func (q *Queue) Size() int {
	return q.itemsCount
}

// expand doubles the underlying slice nd repositions the existing elements
// to start from position 0
func (q *Queue) expand() {
	newSlice := make([]int, 2*len(q.values), 2*cap(q.values))
	copy(newSlice, append(q.values[q.tail:], q.values[:q.tail]...))
	q.tail = 0
	q.values = newSlice
}
