package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pr, nx := head, head
	for i := 0; i < n; i++ {
		nx = nx.Next
	}
	if nx == nil {
		return head.Next
	}
	for nx.Next != nil {
		nx = nx.Next
		pr = pr.Next
	}
	pr.Next = pr.Next.Next

	return head
}

func NewList(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, n := range nums {
		p.Next = &ListNode{n, nil}
		p = p.Next
	}

	return head
}

func (head *ListNode) String() string {
	strBuilder := bytes.NewBuffer(make([]byte, 0, 1024))
	p := head
	for p.Next != nil {
		p = p.Next
		strBuilder.WriteString(strconv.FormatInt(int64(p.Val), 32))
	}

	return strBuilder.String()
}

func main() {
	head := NewList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(head)
	head = removeNthFromEnd(head, 2)
	fmt.Println(head)
	head = removeNthFromEnd(head, 5)
	fmt.Println(head)
}
