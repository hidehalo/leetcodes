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
		i := p.Val[2]
		for i > 0 {
			bytes = append(bytes, byte(p.Val[0]))
			i--
		}
		p = p.Next
	}

	return string(bytes)
}

func insert(head *doubleLinkListNode, val, rank, count int) {
	node := &doubleLinkListNode{nil, head.Tail, nil, 0, [3]int{val, rank, count}}
	if head.Tail != nil {
		head.Tail.Next = node
		node.Prev = head.Tail
	} else {
		head.Next = node
		node.Prev = head
	}
	head.Tail = node
	head.Len++
}

func delete(head, node *doubleLinkListNode) {
	node.Prev.Next = node.Next
	if head.Tail == node {
		head.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
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
	stop := 0
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
	p := head.Next
	order := 0 //0: ASC|1: DESC
	// O(n)
	for i := 0; i < len(ui8s); i++ {
		// fmt.Printf("%d:%c rank:%d,count:%d\n", i, p.Val[0], p.Val[1], p.Val[2])
		if head.Len == 1 {
			ret = append(ret, byte(head.Next.Val[0]))
			continue
		}
		if p.Val[2] > 0 {
			ret = append(ret, byte(p.Val[0]))
			// fmt.Printf("Append %c\n", p.Val[0])
			q := p
			if stop == 0 &&
				p.Val[2] > 1 &&
				((i > 0 && p == head.Next && order == 1) || (p == head.Tail && order == 0)) {
				// fmt.Println("!!! >>> Stop one step of p")
				stop = 1
			} else {
				if p == head.Next {
					order = 0
				} else if p == head.Tail {
					order = 1
				}
				if order == 0 {
					p = p.Next
				} else {
					p = p.Prev
				}
				stop = 0
			}
			q.Val[2]--
			if q.Val[2] <= 0 {
				// fmt.Printf("Delete %c\n", q.Val[0])
				delete(head, q)
			}
		}
	}

	return string(ret)
}
