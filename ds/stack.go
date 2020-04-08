package ds

type Stack struct {
	store []byte
	top   int
}

func (s *Stack) Push(elm byte) {
	s.store = append(s.store, elm)
	s.top++
}

func (s *Stack) Pop() byte {
	ret := s.store[s.top]
	s.top--

	return ret
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Len() int {
	return s.top + 1
}

func NewStack() *Stack {
	return &Stack{
		top: -1,
	}
}
