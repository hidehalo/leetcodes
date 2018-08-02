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

func NewList(nums []int) *ListNode {
	root := ListNode{}
	p := &root
	for _, v := range nums {
		next := ListNode{Val: v}
		p.Next = &next
		p = &next
	}

	return root.Next
}

func (root *ListNode) String() string {
	contents := ""
	p := root
	for p != nil {
		if p.Next == nil {
			contents += fmt.Sprintf("%d", p.Val)
		} else {
			contents += fmt.Sprintf("%d->", p.Val)
		}
		p = p.Next
	}

	return contents
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == l2 {
		return l1
	}
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

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k < 1 {
		return nil
	}
	newLists := make([]*ListNode, 0, k)
	l, r := 0, k-1
	for l <= r {
		newLists = append(newLists, mergeTwoLists(lists[l], lists[r]))
		l++
		r--
	}
	if len(newLists) > 1 {
		return mergeKLists(newLists)
	}

	return newLists[0]
	// var l1, l2 *ListNode
	// for i := 0; i < k; i++ {
	// 	l1 = lists[i]
	// 	l2 = mergeTwoLists(l1, l2)
	// }

	// return l2
}

func main() {
	lists := []*ListNode{
		NewList([]int{2, 5, 8, 12, 15, 18, 22, 100}),
		NewList([]int{1, 2, 3, 4, 5}),
		NewList([]int{6, 7, 8, 9, 10, 11, 13, 21}),
	}
	fmt.Println(mergeKLists(lists))
}
