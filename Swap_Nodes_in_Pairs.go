package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, v := range nums {
		p.Next = &ListNode{
			Val: v,
		}
		p = p.Next
	}

	return head.Next
}

func (head *ListNode) String() string {
	p := head
	contents := ""
	for p != nil {
		contents += strconv.Itoa(p.Val) + " "
		p = p.Next
	}

	return contents
}

func swapPairs(head *ListNode) *ListNode {
	t := &ListNode{}
	t.Next = head
	if head == nil || head.Next == nil {
		return head
	}
	p0 := t.Next
	p1 := p0.Next
	ret := p1
	for p0 != nil && p1 != nil {
		// swap nodes links
		p0.Next = p1.Next
		p1.Next = p0
		t.Next = p1
		// move tail pointer
		t = p0
		// move p0,p1 to next pair
		if t.Next != nil && t.Next.Next != nil {
			p0, p1 = t.Next, t.Next.Next
		} else {
			break
		}
	}

	return ret
}

func main() {
	head := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(head)
	fmt.Println(swapPairs(head))
}
