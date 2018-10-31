package main

import (
	"errors"
	"fmt"
)

type DoubleLinkListNode struct {
	Val  int
	Prev *DoubleLinkListNode
	Next *DoubleLinkListNode
}

type Stack struct {
	head *DoubleLinkListNode
	tail *DoubleLinkListNode
}

func NewStack() *Stack {
	head := &DoubleLinkListNode{}

	return &Stack{head, head}
}

func (stack *Stack) IsEmpty() bool {
	return stack.head.Next == nil
}

func (stack *Stack) Push(val int) {
	node := &DoubleLinkListNode{val, stack.tail, nil}
	stack.tail.Next = node
	stack.tail = node
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	node := stack.tail
	stack.tail = node.Prev
	node.Prev = nil
	stack.tail.Next = nil

	return node.Val, nil
}

func (stack *Stack) ToArray() []int {
	p := stack.head
	arr := make([]int, 0, 100)
	for p.Next != nil {
		p = p.Next
		arr = append(arr, p.Val)
	}

	return arr
}

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0, 1000)
	size := len(candidates)
	if size <= 0 {
		return ret
	}
	stack := NewStack()
	recur(&candidates, target, 0, stack, &ret)

	return ret
}

func recur(candidates *[]int, target int, level int, stack *Stack, ret *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		if !stack.IsEmpty() {
			*ret = append(*ret, stack.ToArray())
		}
	} else {
		for i := level; i < len(*candidates); i++ {
			n := (*candidates)[i]
			if n > target {
				continue
			}
			stack.Push(n)
			recur(candidates, target-n, i, stack, ret)
			stack.Pop()
		}
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
