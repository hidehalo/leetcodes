package main

import "fmt"

type DoubleLinkListNode struct {
	Prev *DoubleLinkListNode
	Next *DoubleLinkListNode
	val  byte
}

func (node *DoubleLinkListNode) Val() byte {
	return node.val
}

type Stack struct {
	store *DoubleLinkListNode
	tail  *DoubleLinkListNode
	len   int
}

func (s *Stack) Push(val byte) {
	s.tail.Next = &DoubleLinkListNode{
		s.tail,
		nil,
		val,
	}
	s.tail = s.tail.Next
	s.len++
}

func (s *Stack) Pop() *DoubleLinkListNode {
	node := s.tail
	s.tail = s.tail.Prev
	s.len--

	return node
}

func (s *Stack) Len() int {
	return s.len
}

func NewStack() *Stack {
	head := &DoubleLinkListNode{}

	return &Stack{
		head,
		head,
		0,
	}
}

func isMatch(s string, p string) bool {
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
	i, j := 0, 0
	stack := NewStack()
	for i < sizeS && j < sizeP {
		if p[j] == '.' {
			// pattern "."
			if j < sizeP-1 && p[j+1] == '*' {
				return true
			}
			j++
			i++
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
					if node.Val() == s[i] {
						i++
						break
					}
				}
			} else {
				return false
			}
		}
	}
	fmt.Println(stack.Len(), i, j)
	// FIXME: it can not do backtracing match when alphabet pattern is longer
	if i < sizeS {
		return false
	}
	switch sizeP - j {
	case 0:
		return true
	case 1:
		if p[j] != '*' {
			return false
		}
		break
	case 2:
		if !(p[j] == '.' && p[j+1] == '*') {
			return false
		}
		break
	default:
		return false
	}

	return false
}

func main() {
	fmt.Println(isMatch("aaa", "a*aa.*"))
}
