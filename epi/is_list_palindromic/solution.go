package is_list_palindromic

import (
	"github.com/stefantds/go-epi-judge/data_structures/list"
)

func IsLinkedListAPalindrome(l *list.Node) bool {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := reverse(slow)

	for i, j := l, secondHalf; i != nil && j != nil; i, j = i.Next, j.Next {
		if i.Data != j.Data {
			return false
		}
	}

	return true
}

func reverse(l *list.Node) *list.Node {
	if l == nil || l.Next == nil {
		return l
	}
	i, j, k := l, l.Next, l.Next.Next
	i.Next = nil
	for j != nil {
		j.Next = i // reverse the link between i and j
		i = j      // advance i
		j = k      // advance j
		if k != nil {
			k = k.Next
		}
	}

	return i
}
