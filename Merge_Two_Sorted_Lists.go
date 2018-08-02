package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l1
			l1 = p.Next.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return head.Next
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	r := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	ret := mergeTwoLists(l, r)

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
