package main

import "testing"

func Test_minTaps(t *testing.T) {
	type args struct {
		n      int
		ranges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{5, []int{3, 4, 1, 1, 0, 0}}, 1},
		{"2", args{3, []int{0, 0, 0, 0}}, -1},
		{"3", args{7, []int{1, 2, 1, 0, 2, 1, 0, 1}}, 3},
		{"4", args{8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}}, 2},
		{"5", args{8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}}, 1},
		{"6", args{8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}}, 1},
		{"7", args{1, []int{0, 0}}, -1},
		{"8", args{0, []int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTaps(tt.args.n, tt.args.ranges); got != tt.want {
				t.Errorf("minTaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
