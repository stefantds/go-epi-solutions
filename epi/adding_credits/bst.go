package adding_credits

// BSTMap implements a BST-based dictionary that can hold keys of
// type int and values of type set of strings (map[string]bool)
type BSTMap struct {
	Key         int
	Value       map[string]bool
	Left, Right *BSTMap
}

func (n *BSTMap) Insert(key int, value map[string]bool) *BSTMap {
	if n == nil {
		return &BSTMap{
			Key:   key,
			Value: value,
		}
	}
	if n.Key < key {
		if n.Right == nil {
			n.Right = &BSTMap{Value: value}
		} else {
			n = n.Right.Insert(key, value)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &BSTMap{Value: value}
		} else {
			n = n.Left.Insert(key, value)
		}
	}
	return n
}

func (n *BSTMap) Find(key int) *BSTMap {
	if n == nil {
		return nil
	}
	if n.Key == key {
		return n
	}
	if n.Key > key {
		return n.Left.Find(key)
	}
	return n.Right.Find(key)
}

func (n *BSTMap) Delete(key int) *BSTMap {
	if n == nil {
		return n
	}
	if n.Key < key {
		n.Right = n.Right.Delete(key)
	} else if n.Key > key {
		n.Left = n.Left.Delete(key)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		min := n.Right.Min()

		n.Key = min.Key
		n.Right = n.Right.Delete(min.Key)
	}
	return n
}

func (n *BSTMap) Min() *BSTMap {
	if n.Left == nil {
		return n
	}
	return n.Left.Min()
}

func (n *BSTMap) Max() *BSTMap {
	if n.Right == nil {
		return n
	}
	return n.Right.Max()
}
