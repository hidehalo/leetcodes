package main

import "../ds"

type Stack struct {
	store *ds.DoubleLinkListNode
	tail  *ds.DoubleLinkListNode
	len   int
}

func (s *Stack) Push(val byte) {
	s.tail.Next = &ds.DoubleLinkListNode{
		s.tail,
		nil,
		val,
	}
	s.tail = s.tail.Next
	s.len++
}

func (s *Stack) Pop() *ds.DoubleLinkListNode {
	node := s.tail
	s.tail = s.tail.Prev
	s.len--

	return node
}

func (s *Stack) Len() int {
	return s.len
}

func NewStack() *Stack {
	head := &ds.DoubleLinkListNode{}

	return &Stack{
		head,
		head,
		0,
	}
}

func isMatch(s string, p string) bool {
	// base validations
	if p == ".*" {
		return true
	}
	if p == "*" {
		return false
	}
	sizeS, sizeP := len(s), len(p)
	if sizeS <= 0 {
		return false
	}
	if sizeP <= 0 {
		return false
	}
	// matching algorithm
	i, j := 0, 0
	stack := NewStack()
	for i < sizeS && j < sizeP {
		if p[j] == '.' {
			// pattern "."
			if j == sizeP-2 && p[j+1] == '*' {
				return true
			} else if j < sizeP-1 && p[j+1] == '*' {
				// .*xxxxx
				// xxxx.*xxxxx
				stack.Push('.')
				j += 2
			}
			if stack.Len() > 0 {
				node := stack.Pop()
				for i < sizeS && (s[i] == node.Val || node.Val == '.') {
					if i < sizeS-1 && node.Val == '.' && s[i] != s[i+1] {
						i++
						break
					}
					i++
				}
			} else {
				j++
				i++
			}
			continue
		} else if j < sizeP-1 && p[j+1] == '*' {
			// pattern "char*"
			// use stack impl lazy match
			stack.Push(p[j])
			j += 2
			continue
		} else {
			// pattern ascii char
			if p[j] == s[i] {
				i++
				j++
			} else if p[j] != s[i] && stack.Len() > 0 {
				for stack.Len() > 0 {
					node := stack.Pop()
					for i < sizeS && (s[i] == node.Val || node.Val == '.') {
						if i < sizeS-1 && node.Val == '.' && s[i] != s[i+1] {
							i++
							break
						}
						i++
					}
				}
			} else {
				return false
			}
		}
	}
	// check tail chars
	if i == sizeS && j != sizeP {
		switch sizeP - j {
		case 1:
			if p[j] != '*' {
				return false
			} else {
				return true
			}
		case 2:
			if !(p[j] != '*' && p[j+1] == '*') {
				return false
			} else {
				return true
			}
		default:
			return false
		}
	}
	// fix up pointer j
	if i < sizeS && j == sizeP {
		for i < sizeS && stack.Len() > 0 {
			node := stack.Pop()
			for i < sizeS && (s[i] == node.Val || node.Val == '.') {
				i++
			}
		}
	}
	// final validations
	if i < sizeS {
		return false
	}
	if sizeP-j == 0 {
		return true
	}

	return false
}
