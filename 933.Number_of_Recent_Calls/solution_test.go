package main

import "testing"

func TestRecentCounter_Ping(t *testing.T) {
	type fields struct {
		head      int
		tail      int
		lastest   int
		pings     []int
		timelines []int
	}
	type args struct {
		ts []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{1, 100, 3001, 3002}},
			[]int{1, 2, 3, 3},
		},
		{
			"case2",
			args{[]int{642, 1849, 4921, 5936, 5957}},
			[]int{1, 2, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for i, time := range tt.args.ts {
				if got := this.Ping(time); got != tt.want[i] {
					t.Errorf("RecentCounter.Ping() = %d, want %d", got, tt.want[i])
				}
			}
		})
	}
}
