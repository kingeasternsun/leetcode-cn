package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 1, 0, 1}}, 3},
		{"2", args{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, 5},
		{"3", args{[]int{1, 1, 1}}, 2},
		{"4", args{[]int{1, 1, 0, 0, 1, 1, 1, 0, 1}}, 4},
		{"5", args{[]int{0, 0, 0}}, 0},
		{"6", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
