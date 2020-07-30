package main

type MyStack struct {
	head  int
	tail  int
	store []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		head:  0,
		tail:  -1,
		store: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.tail++
	if cap(this.store) < len(this.store) {
		this.store = append(this.store, x)
	} else {
		this.store[this.tail] = x
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() == true {
		return -1
	}
	val := this.store[this.tail]
	this.tail--
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() == true {
		return -1
	}
	return this.tail
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.tail < this.head
}
