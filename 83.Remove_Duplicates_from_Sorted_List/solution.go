package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p, q *ListNode
	p = head
	q = p.Next
	for p != nil && q != nil {
		if p.Val == q.Val {
			q = q.Next
			p.Next = q
		} else {
			p = q
			q = q.Next
		}
	}

	return head
}
