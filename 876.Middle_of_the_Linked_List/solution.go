package main

func middleNode(head *ListNode) *ListNode {
	ps := make([]*ListNode, 0)
	p := head
	for p != nil {
		ps = append(ps, p)
		p = p.Next
	}

	return ps[(len(ps)+1)/2]
}
