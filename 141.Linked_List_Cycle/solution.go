package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	pos := -1
	i := 0
	parents := make(map[*ListNode]int)
	parents[head] = 1
	p := head.Next
	for p != nil {
		if parents[p] == 1 {
			pos = i
			break
		}
		parents[p] = 1
		p = p.Next
		i++
	}
	return pos >= 0
}
