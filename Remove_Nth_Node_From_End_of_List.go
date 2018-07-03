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
	size := len(nums)
	head := &ListNode{}
	p := head
	for i := 0; i < size; i++ {
		p.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pa := make([]*ListNode, n)
	p := head
	var parent *ListNode
	i := 0
	for p != nil {
		parent = pa[i]
		pa[i] = p
		i = (i + 1) % n
		p = p.Next
	}
	if parent != nil && pa[i] != nil {
		parent.Next = pa[i].Next
	} else if parent == nil && pa[i] != nil {
		head = pa[i].Next
	} else {
		return nil
	}

	return head
}

func main() {
	head := Constructor([]int{1, 2})
	fmt.Println(head)
	head = removeNthFromEnd(head, 2)
	fmt.Println(head)
}
