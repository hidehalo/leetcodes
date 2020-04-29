package main

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"case1",
			args{"this apple is sweet", "this apple is sour"},
			[]string{"sweet", "sour"},
		},
		{
			"case2",
			args{"apple apple", "banana"},
			[]string{"banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
