package int_as_list_add

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func AddTwoNumbers(l1 *list.Node, l2 *list.Node) *list.Node {
	carry := 0
	sDummyHead := new(list.Node)
	s := sDummyHead
	i1, i2 := l1, l2

	for i1 != nil || i2 != nil || carry != 0 {
		sum := carry
		if i1 != nil {
			sum = sum + i1.Data
			i1 = i1.Next
		}
		if i2 != nil {
			sum = sum + i2.Data
			i2 = i2.Next
		}
		s.Next = &list.Node{
			Data: sum % 10,
		}
		s = s.Next
		carry = sum / 10
	}

	return sDummyHead.Next
}
