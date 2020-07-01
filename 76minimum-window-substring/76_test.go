package main

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{s: "ADOBECODEBANC", t: "ABC"}, "BANC"},
		{"2", args{s: "ab", t: "a"}, "a"},
		{"3", args{s: "ab", t: "b"}, "b"},
		{"4", args{s: "abc", t: "ac"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
