package main

// Given the head of a singly linked list and two integers left and right where left <= right,
// reverse the nodes of the list from position left to position right, and return the reversed list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left != right {
		var p, q, leftPrev, reversed, tailPrev, tail *ListNode
		extraHead := &ListNode{}
		p = extraHead
		p.Next = head
		index := 0
		quit := false
		// find left&right nodes
		for p != nil {
			if index == right {
				quit = true
				tail = p.Next
			}
			if leftPrev == nil {
				if index == left-1 {
					leftPrev = p
				}
			} else {
				// cache p
				q = p
				// "next p after p.Next change"
				p = p.Next
				// link old p cache -> reversed
				if reversed == nil {
					tailPrev = q
				}
				q.Next = reversed
				// "mark cached p as reversed link list(head)"
				reversed = q
				if quit {
					tailPrev.Next = tail
					if leftPrev == extraHead {
						return reversed
					}
					leftPrev.Next = reversed
					break
				}
				index++
				continue
			}
			p = p.Next
			index++
		}
	}
	return head
}
