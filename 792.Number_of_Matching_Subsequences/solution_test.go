package main

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s1 []byte
		s2 []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				[]byte("abcd"),
				[]byte("abcd"),
			},
			true,
		},
		{
			"case2",
			args{
				[]byte("abcd"),
				[]byte("a"),
			},
			true,
		},
		{
			"case3",
			args{
				[]byte("abcd"),
				[]byte("b"),
			},
			true,
		},
		{
			"case4",
			args{
				[]byte("abcd"),
				[]byte("c"),
			},
			true,
		},
		{
			"case5",
			args{
				[]byte("abcd"),
				[]byte("ab"),
			},
			true,
		},
		{
			"case6",
			args{
				[]byte("abcd"),
				[]byte("ac"),
			},
			true,
		},
		{
			"case7",
			args{
				[]byte("abcd"),
				[]byte("ad"),
			},
			true,
		},
		{
			"case8",
			args{
				[]byte("abcd"),
				[]byte("bc"),
			},
			true,
		},
		{
			"case9",
			args{
				[]byte("abcd"),
				[]byte("bd"),
			},
			true,
		},
		{
			"case10",
			args{
				[]byte("abcd"),
				[]byte("cd"),
			},
			true,
		},
		{
			"case11",
			args{
				[]byte("abcd"),
				[]byte("abc"),
			},
			true,
		},
		{
			"case12",
			args{
				[]byte("abcd"),
				[]byte("abd"),
			},
			true,
		},
		{
			"case13",
			args{
				[]byte("abcd"),
				[]byte("acd"),
			},
			true,
		},
		{
			"case14",
			args{
				[]byte("abcd"),
				[]byte("bcd"),
			},
			true,
		},
		{
			"case15",
			args{
				[]byte("abcd"),
				[]byte("adbc"),
			},
			false,
		},
		{
			"case16",
			args{
				[]byte("abcd"),
				[]byte("abcda"),
			},
			false,
		},
		{
			"case17",
			args{
				[]byte("abcd"),
				[]byte("efjk"),
			},
			false,
		},
		{
			"case18",
			args{
				[]byte("abcde"),
				[]byte("ace"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				"abcde",
				[]string{"a", "bb", "acd", "ace"},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMatchingSubseq(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("numMatchingSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
