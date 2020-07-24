package main

type MinStack struct {
	min   int
	top   int
	store []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   0,
		top:   -1,
		store: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if this.top < 0 || this.min > x {
		this.min = x
	}
	this.top++
	if this.top < len(this.store) {
		this.store[this.top] = x
	} else {
		this.store = append(this.store, x)
	}
}

func (this *MinStack) Pop() {
	if this.top >= 0 {
		val := this.store[this.top]
		if val == this.min {
			newMin := this.store[0]
			for i := 0; i < this.top; i++ {
				if this.store[i] < newMin {
					newMin = this.store[i]
				}
			}
			this.min = newMin
		}
		this.top--
	}
}

func (this *MinStack) Top() int {
	if this.top >= 0 {
		return this.store[this.top]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
