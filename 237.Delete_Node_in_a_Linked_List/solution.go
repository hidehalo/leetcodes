package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node != nil && node.Next != nil {
		p := node.Next
		node.Val = p.Val
		node.Next = p.Next
	}
}
