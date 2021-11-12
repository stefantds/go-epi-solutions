package max_of_sliding_window

func ComputeTrafficVolumes(a []TrafficElement, w int) []TrafficElement {
	q := NewQueueWithMax()
	result := make([]TrafficElement, 0)

	for _, e := range a {
		q.Enqueue(e)
		for !q.IsEmpty() && e.Time-q.Head().Time > w {
			q.Dequeue()
		}
		result = append(result, TrafficElement{Time: e.Time, Volume: q.Max().Volume})
	}

	return result
}

type QueueWithMax struct {
	values []TrafficElement
	max    []TrafficElement
}

func NewQueueWithMax() *QueueWithMax {
	return &QueueWithMax{
		values: make([]TrafficElement, 0),
		max:    make([]TrafficElement, 0),
	}
}

func (q *QueueWithMax) Enqueue(x TrafficElement) {
	q.values = append(q.values, x)

	// we can remove the elements in the max queue that are smaller than x
	for len(q.max) > 0 && q.max[len(q.max)-1].Volume < x.Volume {
		q.max = q.max[:len(q.max)-1]
	}
	q.max = append(q.max, x)
}

func (q *QueueWithMax) Dequeue() TrafficElement {
	if len(q.values) == 0 {
		panic("dequeue called on an empty queue")
	}
	var value TrafficElement
	value, q.values = q.values[0], q.values[1:]
	if value.Volume == q.max[0].Volume {
		q.max = q.max[1:] // remove the element from max queue
	}
	return value
}

func (q *QueueWithMax) Max() TrafficElement {
	if len(q.max) == 0 {
		panic("max called on an empty queue")
	}
	return q.max[0]
}

func (q *QueueWithMax) Head() TrafficElement {
	if len(q.values) == 0 {
		panic("head called on an empty queue")
	}
	return q.values[0]
}

func (q *QueueWithMax) IsEmpty() bool {
	return len(q.values) == 0
}
