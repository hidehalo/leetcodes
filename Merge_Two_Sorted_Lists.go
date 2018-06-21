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
	pl := l1
	pr := l2
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := head
	for pl != nil && pr != nil {
		if pl.Val > pr.Val {
			p.Next = &ListNode{
				Val:  pr.Val,
				Next: nil,
			}
			p = p.Next
			pr = pr.Next
		} else {
			p.Next = &ListNode{
				Val:  pl.Val,
				Next: nil,
			}
			p = p.Next
			pl = pl.Next
		}
	}
	for pl != nil {
		p.Next = &ListNode{
			Val:  pl.Val,
			Next: nil,
		}
		p = p.Next
		pl = pl.Next
	}
	for pr != nil {
		p.Next = &ListNode{
			Val:  pr.Val,
			Next: nil,
		}
		p = p.Next
		pr = pr.Next
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
