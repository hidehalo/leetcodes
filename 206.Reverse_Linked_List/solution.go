package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := make([]*ListNode, 0)
	p := head
	i := 0
	for p != nil {
		stack = append(stack, p)
		oldNext := p.Next
		if i >= 1 {
			p.Next = stack[i-1]
		}
		i++
		p = oldNext
	}
	stack[0].Next = nil
	return stack[len(stack)-1]
}
