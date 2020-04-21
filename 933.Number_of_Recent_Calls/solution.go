package main

type RecentCounter struct {
	timelines []int
	index     int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.timelines = append(this.timelines, t)
	for this.index < len(this.timelines) && this.timelines[this.index] < t-3000 {
		this.index++
	}

	if this.index > 100 {
		this.timelines = this.timelines[this.index:]
		this.index = 0
	}

	return len(this.timelines) - this.index
}
