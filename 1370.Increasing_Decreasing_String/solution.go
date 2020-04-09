package main

import (
	"sort"
)

type doubleLinkListNode struct {
	Tail *doubleLinkListNode
	Prev *doubleLinkListNode
	Next *doubleLinkListNode
	Len  int
	Val  [3]int
}

func (head *doubleLinkListNode) String() string {
	bytes := make([]byte, 0)
	p := head.Next
	for p != nil {
		bytes = append(bytes, byte(p.Val[0]))
		p = p.Next
	}

	return string(bytes)
}

func insert(head *doubleLinkListNode, val, rank, count int) {
	node := &doubleLinkListNode{nil, head.Tail, nil, 0, [3]int{val, rank, count}}
	if head.Tail != nil {
		head.Tail.Next = node
	} else {
		head.Next = node
	}
	head.Tail = node
	head.Len++
}

func delete(head, node *doubleLinkListNode) {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if head.Tail == node {
		head.Tail = node.Prev
	}
	node = nil
	head.Len--
}

func sortString(s string) string {
	// convert to int slice O(n)
	ui8s := make([]int, 0)
	for j := 0; j < len(s); j++ {
		ui8s = append(ui8s, int(s[j]))
	}
	// sort order ASC O(nlgn)
	sort.Ints(ui8s)
	counts := make(map[int]int)
	// note: double link list is better than hash table
	ranks := make(map[int]int)
	current := -1
	rank := 0 // init rank
	// build ranks&counts hash table O(n)
	for i := 0; i < len(ui8s); i++ {
		if current != ui8s[i] {
			current = ui8s[i]
			rank++
			ranks[rank] = current
		}
		counts[current]++
	}
	// note: trick for unique character string
	if len(ranks) <= 1 {
		return s
	}
	// build double link list O(n)
	head := &doubleLinkListNode{nil, nil, nil, 0, [3]int{}}
	for i := 1; i <= len(ranks); i++ {
		insert(head, ranks[i], i, counts[ranks[i]])
	}
	ret := make([]byte, 0)
	q, p := head, head.Next
	order := 0 //0: ASC|1: DESC
	// O(n)
	for i := 0; i < len(ui8s); i++ {
		if head.Len == 1 {
			ret = append(ret, byte(head.Next.Val[0]))
			continue
		} else if p == head.Next {
			order = 0
		} else if p == head.Tail {
			order = 1
		}

		if p.Val[2] > 0 {
			ret = append(ret, byte(p.Val[0]))
			if order == 0 {
				p = p.Next
				q = q.Next
			} else {
				p = p.Prev
				q = q.Prev
			}
			if p.Val[2] == 1 {
				delete(head, q)
			}
		}
	}

	return string(ret)
}
