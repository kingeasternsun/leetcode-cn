package leetcode

import "testing"

func Test_maxResult(t *testing.T) {
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
		{"1", args{[]int{1, -1, -2, 4, -7, 3}, 2}, 7},
		{"2", args{[]int{10, -5, -2, 4, 0, 3}, 3}, 17},
		{"3", args{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
