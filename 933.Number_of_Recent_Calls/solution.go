package main

type RecentCounter struct {
	head      int
	tail      int
	lastest   int
	pings     []int
	timelines []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		head:      0,
		tail:      0,
		lastest:   0,
		pings:     make([]int, 3000),
		timelines: make([]int, 3000),
	}
}

func (this *RecentCounter) Ping(t int) int {
	startTs := t - 3000
	if this.lastest == 0 || this.lastest != t {
		this.lastest = t
		this.timelines[this.tail] = t
		this.pings[this.tail]++
		this.tail = (this.tail + 1) % 3000
	} else {
		this.pings[this.tail-1]++
	}
	for i := 0; i < 3000; i++ {
		if this.timelines[i] >= startTs {
			this.head = i
			break
		}
		this.pings[i] = 0
	}
	ret := 0
	if this.head > this.tail {
		for i := this.tail; i >= 0; i-- {
			ret += this.pings[i]
		}
		for i := this.head; i < 3000; i++ {
			ret += this.pings[i]
		}
	} else {
		for i := this.head; i < this.tail; i++ {
			ret += this.pings[i]
		}
	}

	return ret
}
