package copy_posting_list

func CopyPostingsListExtraSpace(l *PostingListNode) *PostingListNode {
	// nodeCopies maps a node in the original list to its copy
	nodeCopies := make(map[*PostingListNode]*PostingListNode)

	for c := l; c != nil; c = c.Next {
		nodeCopies[c] = &PostingListNode{Order: c.Order}
	}
	for c := l; c != nil; c = c.Next {
		nodeCopies[c].Next = nodeCopies[c.Next]
		nodeCopies[c].Jump = nodeCopies[c.Jump]
	}

	return nodeCopies[l]
}

func CopyPostingsList(l *PostingListNode) *PostingListNode {
	if l == nil {
		return nil
	}
	// make a copy for every node and link it as the Next pointer
	for c := l; c != nil; c = c.Next.Next {
		copyC := &PostingListNode{
			Order: c.Order,
			Next:  c.Next,
		}
		c.Next = copyC
	}

	copyL := l.Next

	// set the jump values
	for c := l; c != nil; c = c.Next.Next {
		copyC := c.Next
		if c.Jump != nil {
			copyC.Jump = c.Jump.Next // point to the copy of the jump
		}
	}

	// restore the original next links
	for c := l; c != nil; c = c.Next {
		copyC := c.Next
		c.Next = copyC.Next
		if copyC.Next != nil {
			copyC.Next = copyC.Next.Next // point to the copy of the original next
		}
	}

	return copyL
}
