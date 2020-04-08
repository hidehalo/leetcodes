package main

type Stack struct {
	Top   int
	store []byte
}

func NewStack() *Stack {
	return &Stack{
		Top:   -1,
		store: make([]byte, 0, 1000),
	}
}

func (s *Stack) Pop() byte {
	if s.Top < 0 {
		return 0
	}
	ret := s.store[s.Top]
	s.Top--

	return ret
}

func (s *Stack) Push(e byte) {
	s.Top++
	if len(s.store) <= s.Top {
		s.store = append(s.store, e)
	} else {
		s.store[s.Top] = e
	}
}

func isValid(s string) bool {
	stack := NewStack()
	pattern := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack.Push(s[i])
		} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if stack.Top < 0 {
				return false
			} else if pattern[stack.Pop()] != s[i] {
				return false
			}
		}
	}

	return stack.Top < 0
}
