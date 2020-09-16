package main

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{4},
			[]int{2, 2},
		},
		{
			"case2",
			args{125},
			[]int{25, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
