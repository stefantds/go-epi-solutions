package queue_from_stacks

type Solution interface {
	Enqueue(x int)
	Dequeue() int
}

type QueueFromStacks struct {
	inStack, outStack []int
}

func NewQueueFromStacks() Solution {
	return &QueueFromStacks{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *QueueFromStacks) Enqueue(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *QueueFromStacks) Dequeue() int {
	if len(q.outStack) == 0 {
		if len(q.inStack) == 0 {
			panic("trying to dequeue from an empty queue")
		}
		q.transferStackItems()
	}

	var value int
	q.outStack, value = q.outStack[:len(q.outStack)-1], q.outStack[len(q.outStack)-1]
	return value
}

// moves all the elements from the inStack to the outStack
func (q *QueueFromStacks) transferStackItems() {
	for len(q.inStack) > 0 {
		var value int
		q.inStack, value = q.inStack[:len(q.inStack)-1], q.inStack[len(q.inStack)-1]
		q.outStack = append(q.outStack, value)
	}
}
