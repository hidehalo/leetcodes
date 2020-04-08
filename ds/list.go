package ds

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	h := &ListNode{}
	p := h
	for i := 0; i < len(nums); i++ {
		p.Next = &ListNode{nums[i], nil}
		p = p.Next
	}

	return h
}

type DoubleLinkListNode struct {
	Prev *DoubleLinkListNode
	Next *DoubleLinkListNode
	Val  byte
}
