package leetcode

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{0, 1, 3}}, 2},
		{"2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9}}, 8},
		{"3", args{[]int{0, 1, 3, 4, 5, 6, 7, 8, 9}}, 2},
		{"4", args{[]int{0}}, 1},
		{"5", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
