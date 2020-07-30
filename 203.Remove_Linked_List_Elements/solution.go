package main

func removeElements(head *ListNode, val int) *ListNode {
	var parent *ListNode
	current := head
	for current != nil {
		if current.Val == val {
			if parent != nil {
				parent.Next = current.Next
				current.Next = nil
				current = parent.Next
			} else {
				head = current.Next
				current.Next = nil
                current = head
			}
		} else {
			parent = current
			current = current.Next
		}
	}

	return head
}