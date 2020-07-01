package main

import "testing"

func Test_constrainedSubsetSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{10, 2, -10, 5, 20}, 2}, 37},
		{"2", args{[]int{-1, -2, -3}, 1}, -1},
		{"3", args{[]int{10, -2, -10, -5, 20}, 2}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constrainedSubsetSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("constrainedSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
