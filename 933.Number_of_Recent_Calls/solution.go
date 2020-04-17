package main

type RecentCounter struct {
	currentTlIdx int
	pingsTsMap map[int]int
	timeline []int
}

func Constructor() RecentCounter {
    return RecentCounter{
		currentTlIdx: 1,
		pingsTsMap: make(map[int]int),
		timeline: make([]int, 3000),
	}
}

func (this *RecentCounter) Ping(t int) int {
	ret := 0
	ts := this.timeline[this.currentTlIdx]
	startTs := t-3000
	this.pingsTsMap[t]++
	if ts < startTs {
		this.currentTlIdx = 0
		this.timeline[0] = t
		return ret
	}
	for ts <= t {
		ret += this.pingsTsMap[this.timeline[this.currentTlIdx]]
		this.currentTlIdx++
		ts = this.timeline[this.currentTlIdx]
	}
	return ret
}
