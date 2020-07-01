package main

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{A: []int{1, 2, 3, 2, 1}, B: []int{3, 2, 1, 4, 7}}, 3},
		{"2", args{A: []int{0, 1, 1, 1, 1}, B: []int{1, 0, 1, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
