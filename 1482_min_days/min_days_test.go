package mindays

import (
	"testing"
)

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 10, 3, 10, 2}, 3, 1}, 3},
		{"2", args{[]int{1, 10, 3, 10, 2}, 3, 2}, -1},
		{"3", args{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3}, 12},
		{"4", args{[]int{1000000000, 1000000000}, 1, 1}, 1000000000},
		{"5", args{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{9}, 2}, 3},
		{"2", args{[]int{2, 4, 8, 2}, 4}, 2},
		{"3", args{[]int{7, 17}, 2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.args.nums, tt.args.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
