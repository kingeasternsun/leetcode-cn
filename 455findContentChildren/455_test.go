package main

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{g: []int{1, 2, 3}, s: []int{1, 1}}, 1},
		{"2", args{g: []int{1, 2}, s: []int{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
