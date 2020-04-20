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
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"case1",
			fields{0, 0, 0, make([]int, 3000), make([]int, 3000)},
			args{[]int{1, 100, 3001, 3002}},
			[]int{1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RecentCounter{
				head:      tt.fields.head,
				tail:      tt.fields.tail,
				lastest:   tt.fields.lastest,
				pings:     tt.fields.pings,
				timelines: tt.fields.timelines,
			}
			for i, time := range tt.args.ts {
				if got := this.Ping(time); got != tt.want[i] {
					t.Errorf("RecentCounter.Ping() = %d, want %d", got, tt.want[i])
				}
			}
		})
	}
}
