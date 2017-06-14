package main

import (
	"fmt"
)

type MyQueue struct {
	len   int
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{0, make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	tmp := []int{x}
	this.queue = append(tmp, this.queue...)
	this.len = this.len + 1
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := this.queue[this.len-1]
	fmt.Println(ret)
	this.queue = this.queue[0 : this.len-1]
	this.len = this.len - 1

	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[this.len-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.len <= 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	x := 1
	y := 2
	obj := Constructor()
	obj.Push(x)
	obj.Push(y)
	ret1 := obj.Pop()
	ret2 := obj.Peek()
	fmt.Println(obj.queue, ret1, ret2)
}
