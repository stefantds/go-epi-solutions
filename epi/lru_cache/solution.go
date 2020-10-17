package lru_cache

type KeyValue struct {
	Key   int
	Value int
}

type LRUCache struct {
	valuesMap map[int]*DoublyLinkedNode
	oldestVal *DoublyLinkedNode
	newestVal *DoublyLinkedNode
	capacity  int
}

type DoublyLinkedNode struct {
	Data       KeyValue
	Prev, Next *DoublyLinkedNode
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		valuesMap: make(map[int]*DoublyLinkedNode, 0),
		oldestVal: nil,
		newestVal: nil,
		capacity:  capacity,
	}
}

func (q *LRUCache) Lookup(key int) int {
	val, ok := q.valuesMap[key]
	if ok {
		q.moveToFront(val)
		return val.Data.Value
	}
	return -1
}

func (q *LRUCache) Insert(key, value int) {
	val, ok := q.valuesMap[key]
	if ok {
		q.moveToFront(val)
		return
	}
	newNode := DoublyLinkedNode{
		Data: KeyValue{
			Value: value,
			Key:   key,
		},
	}
	q.valuesMap[key] = &newNode
	q.insertToFront(&newNode)
	if len(q.valuesMap) > q.capacity {
		q.ejectOldest()
	}
}

func (q *LRUCache) Erase(key int) bool {
	val, ok := q.valuesMap[key]
	if !ok {
		return false
	}

	delete(q.valuesMap, key)
	q.disconnectNode(val)
	return true
}

func (q *LRUCache) moveToFront(n *DoublyLinkedNode) {
	if n == q.newestVal {
		return
	}
	q.disconnectNode(n)
	q.insertToFront(n)
}

func (q *LRUCache) insertToFront(n *DoublyLinkedNode) {
	n.Prev = nil
	n.Next = q.newestVal
	if n.Next != nil {
		n.Next.Prev = n
	}
	q.newestVal = n
	if q.oldestVal == nil {
		q.oldestVal = n
	}
}

func (q *LRUCache) ejectOldest() {
	if q.oldestVal == nil {
		return
	}
	delete(q.valuesMap, q.oldestVal.Data.Key)
	q.disconnectNode(q.oldestVal)
}

func (q *LRUCache) disconnectNode(n *DoublyLinkedNode) {
	if n == q.oldestVal {
		q.oldestVal = n.Prev
	}
	if n == q.newestVal {
		q.newestVal = n.Next
	}
	switch {
	case n.Prev == nil && n.Next == nil: // nothing to do, node is already disconnected
		return
	case n.Prev == nil: // node is first
		n.Next.Prev = nil
	case n.Next == nil: // node is last
		n.Prev.Next = nil
	default: // node has both neighbors
		n.Next.Prev = n.Prev
		n.Prev.Next = n.Next
	}
}
