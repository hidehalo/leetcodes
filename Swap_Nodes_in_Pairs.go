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
	p0 := t.Next
	p1 := p0.Next
	for p0 != nil && p1 != nil {
		//swap nodes
		// fmt.Println(p0.Val, p1.Val)
		p1, p0 = p0, p1
		// fmt.Println(p0.Val, p1.Val)
		//set new link from old tail node to new first node
		t.Next = p1
		//set new link from new tail node to next node
		p0.Next = p1.Next
		//set new link between nodes
		p1.Next = p0
		//set new tail node
		t = p0
		// next pair
		p0, p1 = t.Next, t.Next.Next
		// fmt.Println(t.Val, p0.Val, p1.Val)
		// break
	}
	//fix up for odd list
	// if p0 != nil && p1 == nil {
	// 	t.Next = p0
	// }

	return head
}

func main() {
	head := Constructor([]int{1, 2, 3, 4})
	fmt.Println(head)
	swapPairs(head)
	// fmt.Println(swapPairs(head))
}
