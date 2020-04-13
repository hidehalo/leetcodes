package main

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{"IDID"},
			[]int{0, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
