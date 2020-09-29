package lru_cache

import "github.com/stefantds/go-epi-judge/list"

type KeyValue struct {
	Key   int
	Value int
}

type LRUCache struct {
	valuesMap map[int]*list.DoublyLinkedNode
	oldestVal *list.DoublyLinkedNode
	newestVal *list.DoublyLinkedNode
	capacity  int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		valuesMap: make(map[int]*list.DoublyLinkedNode, 0),
		oldestVal: nil,
		newestVal: nil,
		capacity:  capacity,
	}
}

func (q *LRUCache) Lookup(key int) int {
	val, ok := q.valuesMap[key]
	if ok {
		q.moveToFront(val)
		return val.Data.(KeyValue).Value
	}
	return -1
}

func (q *LRUCache) Insert(key, value int) {
	val, ok := q.valuesMap[key]
	if ok {
		q.moveToFront(val)
		return
	}
	newNode := list.DoublyLinkedNode{
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

func (q *LRUCache) moveToFront(n *list.DoublyLinkedNode) {
	if n == q.newestVal {
		return
	}
	q.disconnectNode(n)
	q.insertToFront(n)
}

func (q *LRUCache) insertToFront(n *list.DoublyLinkedNode) {
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
	delete(q.valuesMap, q.oldestVal.Data.(KeyValue).Key)
	q.disconnectNode(q.oldestVal)
}

func (q *LRUCache) disconnectNode(n *list.DoublyLinkedNode) {
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
