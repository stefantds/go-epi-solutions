package minimum_distance_3_sorted_arrays

import "fmt"

type BSTNode struct {
	Data        ArrayData
	Left, Right *BSTNode
}

func (n *BSTNode) Insert(value ArrayData) {
	if n == nil {
		panic("can't insert in a nil BST")
	}
	if less(n.Data, value) {
		if n.Right == nil {
			n.Right = &BSTNode{Data: value}
		} else {
			n.Right.Insert(value)
		}
	} else if more(n.Data, value) {
		if n.Left == nil {
			n.Left = &BSTNode{Data: value}
		} else {
			n.Left.Insert(value)
		}
	}
}

func (n *BSTNode) Delete(value ArrayData) *BSTNode {
	if n == nil {
		panic("can't delete from a nil BST")
	}
	if value.Val == -45 && value.Idx == 0 {
		fmt.Println("removed")
	}
	if less(n.Data, value) {
		n.Right = n.Right.Delete(value)
	} else if more(n.Data, value) {
		n.Left = n.Left.Delete(value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		min := n.Right.Min()

		n.Data = min
		n.Right = n.Right.Delete(min)
	}
	return n
}

func (n BSTNode) Min() ArrayData {
	if n.Left == nil {
		return n.Data
	}
	return n.Left.Min()
}

func (n BSTNode) Max() ArrayData {
	if n.Right == nil {
		return n.Data
	}
	return n.Right.Max()
}
