package main

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumZero(tt.args.n)
			uniq := make(map[int]int)
			sum := 0
			for _, v := range got {
				uniq[v]++
				if uniq[v] > 1 {
					t.Errorf("%v repeated", v)
				}
				sum += v
			}
			if reflect.DeepEqual(0, sum) != true {
				t.Errorf("sum(sumZero()) = %v, want %v", 0, sum)
			}
		})
	}
}
