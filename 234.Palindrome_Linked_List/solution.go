package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	len := 0
	p := head
	for p != nil {
		p = p.Next
		len++
	}
	if len == 0 || len == 1 {
		return true
	}

	// The original list split into two equal-length list
	var q, nextHead *ListNode
	p = head
	mid := len / 2
	for mid > 0 {
		mid--
		q = p
		p = p.Next
	}
	// Ensure that the two lists has same length
	if len%2 != 0 {
		nextHead = p.Next
	} else {
		nextHead = p
	}
	q.Next = nil

	// Reverse new list
	var next, tail *ListNode
	p = nextHead
	for p != nil {
		next = p.Next
		p.Next = tail
		tail = p
		p = next
	}
	nextHead = tail

	// Compare
	p = head
	q = nextHead
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}

	return q == nil && p == nil
}
