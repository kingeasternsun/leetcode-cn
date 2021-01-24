package leetcode

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 3, 5, 4, 7}}, 3},
		{"2", args{[]int{2, 2, 2, 2, 2}}, 1},
		{"3", args{[]int{2}}, 1},
		{"4", args{[]int{}}, 0},
		{"5", args{[]int{1, 3, 5, 7}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
